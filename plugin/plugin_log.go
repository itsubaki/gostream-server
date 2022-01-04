package plugin

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	libgostream "github.com/itsubaki/gostream"
	"github.com/itsubaki/gostream-server/handler"
	"github.com/itsubaki/gostream/stream"
)

var _ Plugin = (*LogEventPlugin)(nil)

type LogEvent struct {
	ID      string
	Time    time.Time `json:"Time"`
	Level   int       `json:"Level"`
	Message string    `json:"Message"`
}

type RequestID struct {
	ID string
}

type LogEventPlugin struct{}

func NewLogEvent(body io.ReadCloser) (LogEvent, error) {
	b, err := io.ReadAll(body)
	if err != nil {
		return LogEvent{}, err
	}
	defer body.Close()

	var event LogEvent
	if uerr := json.Unmarshal(b, &event); uerr != nil {
		return LogEvent{}, uerr
	}

	event.ID = uuid.Must(uuid.NewUUID()).String()

	return event, nil
}

func (p *LogEventPlugin) Setup(h *handler.Handler, path, query string) error {
	s, err := libgostream.New().
		Add(LogEvent{}).
		Query(query)
	if err != nil {
		return fmt.Errorf("new gostream: %v", err)
	}
	h.SetStream(path, s)

	h.GET(path, func(c *gin.Context) {
		s, err := h.Stream(c.Request.RequestURI)
		if err != nil {
			c.JSON(400, err)
			return
		}

		select {
		case events := <-s.Output():
			c.JSON(200, events[len(events)-1])
		default:
			c.JSON(200, stream.Event{Time: time.Now()})
		}
	})

	h.POST(path, func(c *gin.Context) {
		event, err := NewLogEvent(c.Request.Body)
		if err != nil {
			c.JSON(400, err)
			return
		}

		s, err := h.Stream(c.Request.RequestURI)
		if err != nil {
			c.JSON(400, err)
			return
		}

		s.Input() <- event
		c.JSON(200, RequestID{event.ID})
	})

	return nil
}

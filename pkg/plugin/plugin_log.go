package plugin

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/itsubaki/gostream-server/pkg/gostream"
	"github.com/itsubaki/gostream/pkg/event"
	"github.com/itsubaki/gostream/pkg/parser"
)

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
	b, err := ioutil.ReadAll(body)
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

func (h *LogEventPlugin) Setup(g *gostream.GoStream, path, query string) error {
	p := parser.New()
	p.Register("LogEvent", LogEvent{})

	s, err := p.Parse(query)
	if err != nil {
		return fmt.Errorf("parse %s: %v", query, err)
	}
	g.SetWindow(path, s.New())

	g.GET(path, func(c *gin.Context) {
		w, err := g.Window(c.Request.RequestURI)
		if err != nil {
			c.JSON(400, err)
			return
		}

		select {
		case events := <-w.Output():
			c.JSON(200, event.Newest(events))
		default:
			c.JSON(200, event.Event{Time: time.Now()})
		}
	})

	g.POST(path, func(c *gin.Context) {
		event, err := NewLogEvent(c.Request.Body)
		if err != nil {
			c.JSON(400, err)
			return
		}

		w, err := g.Window(c.Request.RequestURI)
		if err != nil {
			c.JSON(400, err)
			return
		}

		w.Input() <- event
		c.JSON(200, RequestID{event.ID})
	})

	return nil
}

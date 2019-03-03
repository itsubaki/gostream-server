package plugin

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/itsubaki/gocep/pkg/event"
	"github.com/itsubaki/gocep/pkg/parser"
	"github.com/itsubaki/gostream/pkg/config"
	"github.com/itsubaki/gostream/pkg/gostream"
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

func (h *LogEventPlugin) Setup(g *gostream.GoStream, r *config.Router) error {
	p := parser.New()
	p.Register("LogEvent", LogEvent{})

	s, err := p.Parse(r.Query)
	if err != nil {
		return fmt.Errorf("parse %s: %v", r.Query, err)
	}
	g.SetWindow(r.Path, s.New())

	g.GET(r.Path, func(c *gin.Context) {
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

	g.POST(r.Path, func(c *gin.Context) {
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

package main

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
)

type LogEvent struct {
	ID      string
	Time    time.Time `json:"time"`
	Level   int       `json:"level"`
	Message string    `json:"message"`
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

	uuid, err := uuid.NewUUID()
	if err != nil {
		return LogEvent{}, err
	}

	event.ID = uuid.String()

	return event, nil
}

func (h *LogEventPlugin) Setup(gost *GoStream, r *Router) error {
	p := parser.New()
	p.Register("LogEvent", LogEvent{})

	s, err := p.Parse(r.Query)
	if err != nil {
		return fmt.Errorf("parse %s: %v", r.Query, err)
	}
	gost.SetWindow(r.Path, s.New(1024))

	gost.GET(r.Path, func(c *gin.Context) {
		w, err := gost.Window(c.Request.RequestURI)
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

	gost.POST(r.Path, func(c *gin.Context) {
		event, err := NewLogEvent(c.Request.Body)
		if err != nil {
			c.JSON(400, err)
			return
		}

		w, err := gost.Window(c.Request.RequestURI)
		if err != nil {
			c.JSON(400, err)
			return
		}

		w.Input() <- event
		c.JSON(200, RequestID{event.ID})
	})

	return nil
}

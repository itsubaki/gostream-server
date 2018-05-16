package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsubaki/gocep"
	uuid "github.com/satori/go.uuid"
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

func NewLogEvent(body io.ReadCloser) (LogEvent, error) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return LogEvent{}, err
	}

	var event LogEvent
	if err := json.Unmarshal(b, &event); err != nil {
		return LogEvent{}, err
	}
	event.ID = uuid.NewV4().String()

	return event, nil
}

func SetupRouter(gost *GoStream, r *Router) error {
	p := gocep.NewParser()
	p.Register("LogEvent", LogEvent{})

	s, err := p.Parse(r.Query)
	if err != nil {
		return fmt.Errorf("parse %s: %v", r.Query, err)
	}
	gost.Register(r.Path, s.New(1024))

	gost.GET(r.Path, func(c *gin.Context) {
		w, err := gost.Window(c.Request.RequestURI)
		if err != nil {
			c.JSON(400, err)
			return
		}

		events := <-w.Output()
		c.JSON(200, gocep.Newest(events))
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

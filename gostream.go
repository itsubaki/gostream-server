package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsubaki/gocep"
	uuid "github.com/satori/go.uuid"
)

type GoStream struct {
	engine *gin.Engine
	window map[string]gocep.Window
	port   string
}

func NewGoStream(config *Config) *GoStream {
	return &GoStream{gin.New(), make(map[string]gocep.Window), config.Port}
}

func (gost *GoStream) Register(name string, w gocep.Window) {
	gost.window[name] = w
}

func (gost *GoStream) Window(name string) (gocep.Window, error) {
	v, ok := gost.window[name]
	if ok {
		return v, nil
	}

	return nil, fmt.Errorf("%s not found.", name)
}

func (gost *GoStream) Run() error {
	gost.ShutdownHook()

	err := gost.Setup()
	if err != nil {
		return err
	}

	return gost.engine.Run(gost.port)
}

func (gost *GoStream) ShutdownHook() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		gost.Close()
		os.Exit(0)
	}()
}

func (gost *GoStream) Close() {
	for name := range gost.window {
		gost.window[name].Close()
	}
}

func (gost *GoStream) Setup() error {

	// select count(*) from LogEvent(10sec) where Level > 2
	w := gocep.NewTimeWindow(10 * time.Second)
	w.SetSelector(gocep.EqualsType{Accept: LogEvent{}})
	w.SetSelector(gocep.LargerThanInt{Name: "Level", Value: 2})
	w.SetFunction(gocep.Count{As: "count(*)"})
	gost.Register("/", w)

	gost.engine.POST("/", func(c *gin.Context) {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, err)
			return
		}

		var event LogEvent
		if unerr := json.Unmarshal(b, &event); unerr != nil {
			c.JSON(400, unerr.Error())
			return
		}
		event.ID = uuid.NewV4().String()

		uri := c.Request.RequestURI
		w, err := gost.Window(uri)
		if err != nil {
			c.JSON(400, err)
			return
		}

		w.Input() <- event
		c.JSON(200, RequestID{event.ID})
	})

	gost.engine.GET("/", func(c *gin.Context) {
		uri := c.Request.RequestURI
		w, err := gost.Window(uri)
		if err != nil {
			c.JSON(400, err)
			return
		}

		events := <-w.Output()
		c.JSON(200, gocep.Newest(events))
	})

	return nil
}

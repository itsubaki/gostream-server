package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/itsubaki/gocep"
)

type GoStream struct {
	engine *gin.Engine
	window WindowSet
	port   string
}

type WindowSet struct {
	set map[string]gocep.Window
}

func (ws *WindowSet) Put(name string, w gocep.Window) {
	ws.set[name] = w
}

func (ws *WindowSet) Get(name string) (gocep.Window, error) {
	v, ok := ws.set[name]
	if ok {
		return v, nil
	}

	return nil, fmt.Errorf("%s not found.", name)
}

func NewGoStream(config *Config) *GoStream {
	wset := WindowSet{make(map[string]gocep.Window)}
	return &GoStream{gin.New(), wset, config.Port}
}

func (gost *GoStream) Run() {
	gost.ShutdownHook()
	gost.engine.Run(gost.port)
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
	for name := range gost.window.set {
		gost.window.set[name].Close()
	}
}

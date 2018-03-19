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
	window map[string]gocep.Window
	port   string
}

func NewGoStream(config *Config) *GoStream {
	return &GoStream{gin.New(), make(map[string]gocep.Window), config.Port}
}

func (gost *GoStream) Register(path string, w gocep.Window) {
	gost.window[path] = w
}

func (gost *GoStream) Window(path string) (gocep.Window, error) {
	v, ok := gost.window[path]
	if ok {
		return v, nil
	}

	return nil, fmt.Errorf("%s not found.", path)
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

func (gost *GoStream) Run() error {
	gost.ShutdownHook()
	return gost.engine.Run(gost.port)
}

func (gost *GoStream) POST(path string, handlers ...gin.HandlerFunc) {
	gost.engine.POST(path, handlers...)
}

func (gost *GoStream) GET(path string, handlers ...gin.HandlerFunc) {
	gost.engine.GET(path, handlers...)
}

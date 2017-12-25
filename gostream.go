package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"

	"github.com/gin-gonic/gin"
	"github.com/itsubaki/gocep"
)

type GoStream struct {
	ctx    context.Context
	cancel func()
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
	ctx, cancel := context.WithCancel(context.Background())
	wset := WindowSet{make(map[string]gocep.Window)}
	return &GoStream{ctx, cancel, gin.New(), wset, config.Port}
}

func (gost *GoStream) Run() {
	gost.shutdownHook()
	gost.engine.Run(gost.port)
}

func (gost *GoStream) shutdownHook() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		gost.Close()
		os.Exit(0)
	}()
}

func (gost *GoStream) Close() {
	gost.cancel()
}

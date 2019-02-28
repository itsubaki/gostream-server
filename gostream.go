package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/itsubaki/gocep/pkg/window"
)

type GoStream struct {
	engine *gin.Engine
	window map[string]window.Window
	plugin map[string]GoStreamPlugin
	config *Config
}

func NewGoStream(config *Config) *GoStream {
	gost := &GoStream{
		gin.New(),
		make(map[string]window.Window),
		make(map[string]GoStreamPlugin),
		config,
	}
	return gost
}

func (gost *GoStream) SetPlugin(name string, plugin GoStreamPlugin) {
	gost.plugin[name] = plugin
}

func (gost *GoStream) SetWindow(path string, w window.Window) {
	gost.window[path] = w
}

func (gost *GoStream) Window(path string) (window.Window, error) {
	v, ok := gost.window[path]
	if ok {
		return v, nil
	}

	return nil, fmt.Errorf("window not found=%s", path)
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
	for _, r := range gost.config.Router {
		p, ok := gost.plugin[r.Plugin]
		if !ok {
			fmt.Printf("plugin not found=%s\n", r.Plugin)
			continue
		}

		if err := p.Setup(gost, &r); err != nil {
			fmt.Printf("setup plugin %v: %v\n", r, err)
		}
	}

	gost.ShutdownHook()
	return gost.engine.Run(gost.config.Port)
}

func (gost *GoStream) POST(path string, handlers ...gin.HandlerFunc) {
	gost.engine.POST(path, handlers...)
}

func (gost *GoStream) GET(path string, handlers ...gin.HandlerFunc) {
	gost.engine.GET(path, handlers...)
}

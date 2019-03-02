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
	config *Config
	engine *gin.Engine
	window map[string]window.Window
	plugin map[string]GoStreamPlugin
}

func NewGoStream(config *Config) *GoStream {
	return &GoStream{
		config: config,
		engine: gin.New(),
		window: make(map[string]window.Window),
		plugin: make(map[string]GoStreamPlugin),
	}
}

func (g *GoStream) SetPlugin(name string, plugin GoStreamPlugin) {
	g.plugin[name] = plugin
}

func (g *GoStream) SetWindow(path string, w window.Window) {
	g.window[path] = w
}

func (g *GoStream) Window(path string) (window.Window, error) {
	v, ok := g.window[path]
	if ok {
		return v, nil
	}

	return nil, fmt.Errorf("window not found=%s", path)
}

func (g *GoStream) ShutdownHook() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		g.Close()
		os.Exit(0)
	}()
}

func (g *GoStream) Close() {
	for name := range g.window {
		g.window[name].Close()
	}
}

func (g *GoStream) Run() error {
	for _, r := range g.config.Router {
		p, ok := g.plugin[r.Plugin]
		if !ok {
			fmt.Printf("plugin not found=%s\n", r.Plugin)
			continue
		}

		if err := p.Setup(g, &r); err != nil {
			fmt.Printf("setup plugin %v: %v\n", r, err)
		}
	}

	g.ShutdownHook()
	return g.engine.Run(g.config.Port)
}

func (g *GoStream) POST(path string, handlers ...gin.HandlerFunc) {
	g.engine.POST(path, handlers...)
}

func (g *GoStream) GET(path string, handlers ...gin.HandlerFunc) {
	g.engine.GET(path, handlers...)
}

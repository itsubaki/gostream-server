package gostream

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/itsubaki/gostream/pkg/window"
)

type GoStream struct {
	engine *gin.Engine
	window map[string]window.Window
}

func New() *GoStream {
	return &GoStream{
		engine: gin.New(),
		window: make(map[string]window.Window),
	}
}

func (g *GoStream) Handler() *gin.Engine {
	return g.engine
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

func (g *GoStream) Close() {
	for name := range g.window {
		g.window[name].Close()
	}
}

func (g *GoStream) POST(path string, handlers ...gin.HandlerFunc) {
	g.engine.POST(path, handlers...)
}

func (g *GoStream) GET(path string, handlers ...gin.HandlerFunc) {
	g.engine.GET(path, handlers...)
}

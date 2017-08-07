package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	cep "github.com/itsubaki/gocep"
)

type GoStream struct {
	ctx     context.Context
	config  *Config
	handler []Handler
	cancel  func()
}

func NewGoStream(config *Config) *GoStream {
	ctx, cancel := context.WithCancel(context.Background())
	handler := []Handler{}

	window := cep.NewTimeWindow(3*time.Second, 1024)
	window.SetFunction(cep.Count{As: "cnt"})
	stream := cep.NewStream(1024)
	stream.SetWindow(window)

	h := &RequestHandler{
		ctx,
		"",
		stream,
		NewOutput(config),
	}
	handler = append(handler, h)

	gost := &GoStream{
		ctx,
		config,
		handler,
		cancel,
	}

	return gost
}

func (s *GoStream) Run() {
	router := gin.Default()
	for _, h := range s.handler {
		router.POST(h.URI(), h.POST)
		go h.Listen()
	}
	router.Run(s.config.Port)
}

func (s *GoStream) ShutdownHook() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c // blocking
		s.Close()
		os.Exit(0)
	}()
}

func (s *GoStream) Close() {
	s.cancel()
	for _, h := range s.handler {
		h.Output().Close()
	}
}

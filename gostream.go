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
	config  Config
	handler []Handler
	ctx     context.Context
	cancel  func()
}

func NewGoStream(config Config) *GoStream {
	ctx, cancel := context.WithCancel(context.Background())
	handler := []Handler{}

	stream := cep.NewStream(1024)
	window := cep.NewTimeWindow(3*time.Second, 1024)
	window.SetFunction(cep.Count{As: "cnt"})
	stream.SetWindow(window)
	handler = append(handler, &RequestHandler{"", stream, ctx, []cep.Event{}})

	gost := &GoStream{
		config,
		handler,
		ctx,
		cancel,
	}

	return gost
}

func (s *GoStream) Close() {
	s.cancel()
}

func (s *GoStream) Run() {
	router := gin.Default()
	for _, h := range s.handler {
		router.POST(h.URI(), h.POST)
		router.GET(h.URI(), h.GET)
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

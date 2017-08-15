package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"

	"github.com/gin-gonic/gin"
	cep "github.com/itsubaki/gocep"
	"github.com/itsubaki/gostream/config"
	hdl "github.com/itsubaki/gostream/handler"
	"github.com/itsubaki/gostream/output"
)

type GoStream struct {
	ctx     context.Context
	config  *config.Config
	handler []hdl.Handler
	cancel  func()
}

func NewGoStream(config *config.Config) *GoStream {
	ctx, cancel := context.WithCancel(context.Background())

	window := cep.NewTimeWindow(3*time.Second, 1024)
	window.SetFunction(cep.Count{As: "cnt"})
	stream := cep.NewStream(1024)
	stream.SetWindow(window)

	h := hdl.NewDefaultHandler(
		ctx,
		"",
		stream,
		output.New(config),
	)

	handler := []hdl.Handler{}
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
	router := gin.New()
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

func (s *GoStream) Close() {
	s.cancel()
	for _, h := range s.handler {
		h.Output().Close()
	}
}

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"

	"github.com/gin-gonic/gin"
	"github.com/itsubaki/gostream/config"
	hdl "github.com/itsubaki/gostream/handler"
	"github.com/itsubaki/gostream/output"
)

type GoStream struct {
	ctx     context.Context
	cancel  func()
	port    string
	handler []hdl.Handler
}

func NewGoStream(config *config.Config) *GoStream {
	ctx, cancel := context.WithCancel(context.Background())
	handler := []hdl.Handler{}

	out := output.New(config.OutConfig)
	if h, err := hdl.NewDefaultHandler(ctx, "", out); err != nil {
		log.Println(err)
	} else {
		handler = append(handler, h)
	}

	gost := &GoStream{ctx, cancel, config.Port, handler}

	return gost
}

func (s *GoStream) Run() {
	router := gin.New()

	for _, h := range s.handler {
		router.POST(h.URI(), h.POST)
		router.GET(h.URI(), h.GET)
		go h.Listen()
	}

	router.Run(s.port)
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

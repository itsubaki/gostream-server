package main

import (
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

func NewGoStream(config *config.Config) (*GoStream, error) {
	ctx, cancel := context.WithCancel(context.Background())
	gost := &GoStream{ctx, cancel, config.Port, []hdl.Handler{}}

	out := output.New(config.OutConfig)
	h, err := hdl.NewDefaultHandler("", out)
	if err != nil {
		return nil, err
	}
	gost.AddHandler(h)

	return gost, nil
}

func (s *GoStream) AddHandler(h hdl.Handler) {
	s.handler = append(s.handler, h)
}

func (s *GoStream) Run() {
	engine := gin.New()

	for _, h := range s.handler {
		h.Add(engine)
		go h.Listen()
	}

	engine.Run(s.port)
}

func (s *GoStream) ShutdownHook() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		s.Close()
		os.Exit(0)
	}()
}

func (s *GoStream) Close() {
	s.cancel()
	for _, h := range s.handler {
		h.Close()
	}
}

package lib

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	cep "github.com/itsubaki/gocep"
)

type GoStream struct {
	config Config
	stream *cep.Stream
	Canceller
}

func NewGoStream(config Config) *GoStream {
	stream := NewBuilder().Build()

	gost := &GoStream{
		config,
		stream,
		NewCanceller(),
	}

	go gost.listen()
	return gost
}

func (s *GoStream) Close() {
	s.Cancel()
}

func (s *GoStream) Update(e []cep.Event) {
	log.Println(e)
}

func (s *GoStream) Rcv(c *gin.Context) {
	m := make(map[string]interface{})
	for k, v := range c.Request.Header {
		m[k] = v[0]
	}
	s.stream.Input() <- cep.MapEvent{Record: m}
}

func (s *GoStream) Run() {
	router := gin.Default()
	router.POST("/", s.Rcv)
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

func (s *GoStream) listen() {
	for {
		select {
		case <-s.Ctx.Done():
			return
		case e := <-s.stream.Output():
			s.Update(e)
		}
	}
}

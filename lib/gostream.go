package lib

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	cep "github.com/itsubaki/gocep"
)

type GoStream struct {
	config  Config
	handler [](*RequestHandler)
	Canceller
}

func NewGoStream(config Config) *GoStream {
	canceller := NewCanceller()
	handler := [](*RequestHandler){}

	for i := 0; i < 3; i++ {
		stream := cep.NewStream(1024)
		window := cep.NewTimeWindow(3*time.Second, 1024)
		window.Function(cep.Count{As: "cnt"})
		stream.Window(window)
		handler = append(handler, &RequestHandler{"/foobar/" + strconv.Itoa(i), stream, canceller.Ctx})
	}

	gost := &GoStream{
		config,
		handler,
		canceller,
	}

	return gost
}

func (s *GoStream) Close() {
	s.Cancel()
}

func (s *GoStream) Update(e []cep.Event) {
	log.Println(e)
}

func (s *GoStream) Run() {
	router := gin.Default()
	for _, h := range s.handler {
		router.POST(h.uri, h.Handle)
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

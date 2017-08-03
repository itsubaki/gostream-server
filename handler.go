package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	cep "github.com/itsubaki/gocep"
)

type Handler interface {
	URI() string
	POST(c *gin.Context)
	GET(c *gin.Context)
	Listen()
	Update(e []cep.Event)
}

type RequestHandler struct {
	uri    string
	stream *cep.Stream
	ctx    context.Context
}

func (h *RequestHandler) URI() string {
	return h.uri
}

func (h *RequestHandler) GET(c *gin.Context) {
	c.JSON(http.StatusOK, h.stream.Window())
}

func (h *RequestHandler) POST(c *gin.Context) {
	m := make(map[string]interface{})
	for k, v := range c.Request.Header {
		m[k] = v[0]
	}
	h.stream.Input() <- cep.MapEvent{Record: m}
}

func (h *RequestHandler) Listen() {
	for {
		select {
		case <-h.ctx.Done():
			break
		case e := <-h.stream.Output():
			h.Update(e)
		}
	}
}

func (h *RequestHandler) Update(event []cep.Event) {
	for _, e := range event {
		log.Println(e)
	}
}

package lib

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	cep "github.com/itsubaki/gocep"
)

type RequestHandler struct {
	uri    string
	stream *cep.Stream
	ctx    context.Context
}

func (h *RequestHandler) Handle(c *gin.Context) {
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

func (h *RequestHandler) Update(e []cep.Event) {
	log.Println(e)
}
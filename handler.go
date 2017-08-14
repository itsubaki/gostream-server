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
	Output() Output
	Listen()
	POST(c *gin.Context)
	GET(c *gin.Context)
}

type RequestHandler struct {
	ctx    context.Context
	uri    string
	stream *cep.Stream
	output Output
}

func (h *RequestHandler) URI() string {
	return h.uri
}

func (h *RequestHandler) Output() Output {
	return h.output
}

func (h *RequestHandler) Listen() {
	for {
		select {
		case <-h.ctx.Done():
			break
		case e := <-h.stream.Output():
			h.Output().Update(e)
		}
	}
}

func (h *RequestHandler) POST(c *gin.Context) {
	m := make(map[string]interface{})
	for k, v := range c.Request.Header {
		m[k] = v[0]
	}
	h.stream.Input() <- cep.MapEvent{Record: m}
}

func (h *RequestHandler) GET(c *gin.Context) {
	json, err := Json(h.stream.Window()[0].Event(), true)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, json)
}

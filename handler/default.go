package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	cep "github.com/itsubaki/gocep"
	"github.com/itsubaki/gostream/config"
	"github.com/itsubaki/gostream/output"
	"golang.org/x/net/context"
)

type DefaultHandler struct {
	ctx    context.Context
	uri    string
	stream *cep.Stream
	out    output.Output
}

func NewDefaultHandler(ctx context.Context, uri string, st *cep.Stream, out output.Output) *DefaultHandler {
	return &DefaultHandler{ctx, uri, st, out}
}

func (h *DefaultHandler) URI() string {
	return h.uri
}

func (h *DefaultHandler) Output() output.Output {
	return h.out
}

func (h *DefaultHandler) Listen() {
	for {
		select {
		case <-h.ctx.Done():
			break
		case e := <-h.stream.Output():
			h.Output().Update(e)
		}
	}
}

func (h *DefaultHandler) POST(c *gin.Context) {
	m := make(map[string]interface{})
	for k, v := range c.Request.Header {
		m[k] = v[0]
	}
	h.stream.Input() <- cep.MapEvent{Record: m}
}

func (h *DefaultHandler) GET(c *gin.Context) {
	json, err := config.Json(h.stream.Window()[0].Event(), true)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, json)
}

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
	cancel func()
	uri    string
	out    output.Output
	stream *cep.Stream
}

func NewDefaultHandler(uri string, out output.Output) (*DefaultHandler, error) {
	ctx, cancel := context.WithCancel(context.Background())

	q := "select * from MapEvent.length(3)"
	stmt, err := cep.NewParser(q).Parse()
	if err != nil {
		return nil, err
	}
	stream := stmt.NewStream(1024)

	return &DefaultHandler{ctx, cancel, uri, out, stream}, nil
}

func (h *DefaultHandler) URI() string {
	return h.uri
}

func (h *DefaultHandler) Output() output.Output {
	return h.out
}

func (h *DefaultHandler) Close() {
	h.cancel()
	h.out.Close()
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

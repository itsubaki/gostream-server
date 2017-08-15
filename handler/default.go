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
	Context context.Context
	Uri     string
	Stream  *cep.Stream
	Out     output.Output
}

func (h *DefaultHandler) URI() string {
	return h.Uri
}

func (h *DefaultHandler) Output() output.Output {
	return h.Out
}

func (h *DefaultHandler) Listen() {
	for {
		select {
		case <-h.Context.Done():
			break
		case e := <-h.Stream.Output():
			h.Output().Update(e)
		}
	}
}

func (h *DefaultHandler) POST(c *gin.Context) {
	m := make(map[string]interface{})
	for k, v := range c.Request.Header {
		m[k] = v[0]
	}
	h.Stream.Input() <- cep.MapEvent{Record: m}
}

func (h *DefaultHandler) GET(c *gin.Context) {
	json, err := config.Json(h.Stream.Window()[0].Event(), true)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, json)
}

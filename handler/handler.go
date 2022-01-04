package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/itsubaki/gostream/stream"
)

type Handler struct {
	raw    *gin.Engine
	stream map[string]*stream.Stream
}

func New() *Handler {
	return &Handler{
		raw:    gin.New(),
		stream: make(map[string]*stream.Stream),
	}
}

func (h *Handler) Raw() *gin.Engine {
	return h.raw
}

func (h *Handler) SetStream(path string, s *stream.Stream) {
	h.stream[path] = s
}

func (h *Handler) Stream(path string) (*stream.Stream, error) {
	v, ok := h.stream[path]
	if ok {
		return v, nil
	}

	return nil, fmt.Errorf("stream not found=%s", path)
}

func (h *Handler) Close() {
	for n := range h.stream {
		h.stream[n].Close()
	}
}

func (h *Handler) POST(path string, handlers ...gin.HandlerFunc) {
	h.raw.POST(path, handlers...)
}

func (h *Handler) GET(path string, handlers ...gin.HandlerFunc) {
	h.raw.GET(path, handlers...)
}

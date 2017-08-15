package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/itsubaki/gostream/output"
)

type Handler interface {
	URI() string
	Output() output.Output
	Listen()
	POST(c *gin.Context)
	GET(c *gin.Context)
}

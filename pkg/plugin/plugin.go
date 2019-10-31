package plugin

import (
	"github.com/itsubaki/gostream-api/pkg/config"
	"github.com/itsubaki/gostream-api/pkg/gostream"
)

type Plugin interface {
	Setup(g *gostream.GoStream, r *config.Router) error
}

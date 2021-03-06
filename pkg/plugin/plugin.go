package plugin

import (
	"github.com/itsubaki/gostream-server/pkg/config"
	"github.com/itsubaki/gostream-server/pkg/gostream"
)

type Plugin interface {
	Setup(g *gostream.GoStream, r *config.Router) error
}

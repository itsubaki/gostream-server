package plugin

import (
	"github.com/itsubaki/gostream/pkg/config"
	"github.com/itsubaki/gostream/pkg/gostream"
)

type Plugin interface {
	Setup(g *gostream.GoStream, r *config.Router) error
}

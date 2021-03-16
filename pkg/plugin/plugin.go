package plugin

import (
	"github.com/itsubaki/gostream-server/pkg/gostream"
)

type Plugin interface {
	Setup(g *gostream.GoStream, path, query string) error
}

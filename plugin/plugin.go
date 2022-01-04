package plugin

import "github.com/itsubaki/gostream-server/handler"

type Plugin interface {
	Setup(h *handler.Handler, path, query string) error
}

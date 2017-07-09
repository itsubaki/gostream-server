package lib

import (
	"time"

	cep "github.com/itsubaki/gocep"
)

type Builder struct {
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build() *cep.Stream {
	stream := cep.NewStream(1024)
	window := cep.NewTimeWindow(3*time.Second, 1024)
	window.Function(cep.Count{As: "cnt"})
	stream.Window(window)

	return stream
}

package output

import (
	"fmt"
	"log"

	cep "github.com/itsubaki/gocep"
	"github.com/itsubaki/gostream/config"
)

type OutputStdOut struct {
	pretty bool
}

func (o *OutputStdOut) Update(event []cep.Event) {
	json, err := config.Json(event, o.pretty)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(json)
}

func (o *OutputStdOut) Close() {
	//noop
}

package main

import (
	"fmt"
	"os"

	"github.com/itsubaki/gostream-server/pkg/config"
	"github.com/itsubaki/gostream-server/pkg/gostream"
	"github.com/itsubaki/gostream-server/pkg/plugin"
)

func main() {
	c, err := config.New()
	if err != nil {
		fmt.Printf("new config: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("config=%v\n", c)

	g := gostream.New()
	p := map[string]plugin.Plugin{
		"LogEventPlugin": &plugin.LogEventPlugin{},
	}

	for _, r := range c.Router {
		pp, ok := p[r.Plugin]
		if !ok {
			fmt.Printf("invalid plugin=%v", r.Plugin)
			os.Exit(1)
		}

		if err := pp.Setup(g, r.Path, r.Query); err != nil {
			fmt.Printf("setup failed. path=%v query=%v: %v", r.Path, r.Query, err)
			os.Exit(1)
		}
	}

	if err := g.Run(c.Port); err != nil {
		fmt.Printf("run: %v\n", err)
		os.Exit(1)
	}
}

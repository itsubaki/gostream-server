package main

import (
	"fmt"
	"os"

	"github.com/itsubaki/gostream-api/pkg/config"
	"github.com/itsubaki/gostream-api/pkg/gostream"
	"github.com/itsubaki/gostream-api/pkg/plugin"
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
		p[r.Plugin].Setup(g, &r)
	}

	if err := g.Run(c.Port); err != nil {
		fmt.Printf("run: %v\n", err)
		os.Exit(1)
	}
}

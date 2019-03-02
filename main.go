package main

import (
	"fmt"
	"os"
)

func main() {
	c, err := NewConfig()
	if err != nil {
		fmt.Printf("new config: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("config=%v\n", c)

	g := NewGoStream(c)
	g.SetPlugin("LogEventPlugin", &LogEventPlugin{})

	if err := g.Run(); err != nil {
		fmt.Printf("run: %v\n", err)
		os.Exit(1)
	}
}

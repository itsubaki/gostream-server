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
	fmt.Printf("config=%s\n", c.String())

	gost := NewGoStream(c)
	gost.plugin["LogEventPlugin"] = &LogEventPlugin{}

	if err := gost.Run(); err != nil {
		fmt.Printf("run: %v\n", err)
		os.Exit(1)
	}
}

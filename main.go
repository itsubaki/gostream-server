package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	c, err := NewConfig()
	if err != nil {
		fmt.Printf("new config: %v", err)
		os.Exit(1)
	}
	log.Println("config: " + c.String())

	gost := NewGoStream(c)

	for _, r := range c.Router {
		if err := SetupRouter(gost, &r); err != nil {
			fmt.Printf("setup handler %v: %v", r, err)
			os.Exit(1)
		}
	}

	if err := gost.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

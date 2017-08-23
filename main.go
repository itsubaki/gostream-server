package main

import (
	"log"
	"os"

	"github.com/itsubaki/gostream/config"
)

func main() {
	c := config.New()
	log.Println("config: " + c.String())

	gost, err := NewGoStream(c)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	gost.ShutdownHook()
	gost.Run()
}

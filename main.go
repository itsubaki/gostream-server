package main

import (
	"log"

	"github.com/itsubaki/gostream/config"
)

func main() {
	c := config.New()
	log.Println("config: " + c.String())

	gost := NewGoStream(c)
	gost.ShutdownHook()
	gost.Run()
}

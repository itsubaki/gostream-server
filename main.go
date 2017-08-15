package main

import (
	"log"

	"github.com/itsubaki/gostream/config"
)

func main() {
	c := config.NewConfig()
	log.Println("config: " + c.String())

	gost := NewGoStream(c)
	gost.ShutdownHook()
	gost.Run()
}

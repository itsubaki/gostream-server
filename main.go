package main

import "log"

func main() {
	c := NewConfig()
	log.Println("config: " + c.String())

	gost := NewGoStream(c)
	gost.ShutdownHook()
	gost.Run()
}

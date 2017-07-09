package main

import "github.com/itsubaki/gostream/lib"

func main() {
	c := lib.NewConfig()
	gost := lib.NewGoStream(c)
	gost.ShutdownHook()
	gost.Run()
}

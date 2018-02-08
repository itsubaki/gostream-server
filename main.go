package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	c := NewConfig()
	log.Println("config: " + c.String())

	gost := NewGoStream(c)
	err := gost.Run()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

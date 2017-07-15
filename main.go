package main

func main() {
	c := NewConfig()
	gost := NewGoStream(c)
	gost.ShutdownHook()
	gost.Run()
}

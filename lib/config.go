package lib

import "os"

type Config struct {
	Port string
}

func NewConfig() Config {
	port := Get("GOSTREAM_PORT", "1234")
	return Config{
		Port: ":" + port,
	}
}

func Get(env, init string) string {
	val := os.Getenv(env)
	if len(val) == 0 {
		return init
	}
	return val
}

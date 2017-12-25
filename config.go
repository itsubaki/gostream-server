package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{
		Port: GetString("GOSTREAM_LISTEN_PORT", ":1234"),
	}
}

func (c *Config) String() string {
	b, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func GetString(env, init string) string {
	val := os.Getenv(env)
	if len(val) == 0 {
		return init
	}
	return val
}

func GetBool(env string, init bool) bool {
	val := os.Getenv(env)
	if len(val) == 0 {
		return init
	}
	ret, err := strconv.ParseBool(val)
	if err != nil {
		log.Println(err)
		return init
	}
	return ret
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Port   string   `yaml:"port"`
	Router []Router `yaml:"router"`
}

type Router struct {
	Path  string `yaml:"path"`
	Query string `yaml:"query"`
}

func NewConfig() (*Config, error) {
	p := GetString("GOSTREAM_CONFIG", "gostream.yml")
	buf, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("read %s: %v", p, err)
	}

	var config Config
	if err := yaml.Unmarshal(buf, &config); err != nil {
		return nil, fmt.Errorf("unmarshal %s: %v", p, err)
	}

	return &config, nil
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

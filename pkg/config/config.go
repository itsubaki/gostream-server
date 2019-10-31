package config

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Port   string   `yaml:"port"`
	Router []Router `yaml:"router"`
}

type Router struct {
	Plugin string `yaml:"plugin"`
	Path   string `yaml:"path"`
	Query  string `yaml:"query"`
}

func New() (*Config, error) {
	path := os.Getenv("GOSTREAM_CONFIG")
	if len(path) < 1 {
		path = "config.yml"
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %v", path, err)
	}

	var config Config
	if err := yaml.Unmarshal(buf, &config); err != nil {
		return nil, fmt.Errorf("unmarshal %s: %v", path, err)
	}

	return &config, nil
}

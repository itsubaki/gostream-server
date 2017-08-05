package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port      string
	Output    string
	ProjectID string
	Topic     string
	Logger    string
}

func NewConfig() Config {
	port := Get("GOSTREAM_PORT", "1234")
	output := Get("GOSTREAM_OUTPUT", "stdout")
	projectID := Get("GOSTREAM_PROJECT_ID", "")
	topic := Get("GOSTREAM_PUBSUB_TOPIC", "")
	logger := Get("GOSTREAM_LOGGING_LOGGER", "gostream")

	return Config{
		Port:      ":" + port,
		Output:    output,
		ProjectID: projectID,
		Topic:     topic,
		Logger:    logger,
	}
}

func (c *Config) String() string {
	b, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func Get(env, init string) string {
	val := os.Getenv(env)
	if len(val) == 0 {
		return init
	}
	return val
}

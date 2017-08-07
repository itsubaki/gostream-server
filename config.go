package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port      string
	Output    string
	ProjectID string
	Topic     string
	Logger    string
	Pretty    bool
}

func NewConfig() *Config {
	port := Get("GOSTREAM_PORT", "1234")
	output := Get("GOSTREAM_OUTPUT", "stdout")
	pretty := Get("GOSTREAM_OUTPUT_PRETTY", "false")
	projectID := Get("GOSTREAM_PROJECT_ID", "")
	topic := Get("GOSTREAM_PUBSUB_TOPIC", "")
	logger := Get("GOSTREAM_LOGGING_LOGGER", "gostream")

	tof, err := strconv.ParseBool(pretty)
	if err != nil {
		tof = false
		log.Println(err)
	}

	return &Config{
		Port:      ":" + port,
		Output:    output,
		ProjectID: projectID,
		Topic:     topic,
		Logger:    logger,
		Pretty:    tof,
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

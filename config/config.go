package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port      string
	OutConfig *OutputConfig
}

type OutputConfig struct {
	Output    string
	ProjectID string
	Topic     string
	Logger    string
	Pretty    bool
}

func New() *Config {
	port := Get("GOSTREAM_PORT", "1234")
	output := Get("GOSTREAM_OUTPUT", "stdout")
	projectID := Get("GOSTREAM_PROJECT_ID", "")
	topic := Get("GOSTREAM_PUBSUB_TOPIC", "")
	logger := Get("GOSTREAM_LOGGING_LOGGER", "gostream")
	pretty := GetBool("GOSTREAM_OUTPUT_PRETTY", false)

	return &Config{
		Port: ":" + port,
		OutConfig: &OutputConfig{
			Output:    output,
			ProjectID: projectID,
			Topic:     topic,
			Logger:    logger,
			Pretty:    pretty,
		},
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

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
	return &Config{
		Port: Get("GOSTREAM_PORT", ":1234"),
		OutConfig: &OutputConfig{
			Output:    Get("GOSTREAM_OUTPUT", "stdout"),
			ProjectID: Get("GOSTREAM_PROJECT_ID", ""),
			Topic:     Get("GOSTREAM_PUBSUB_TOPIC", ""),
			Logger:    Get("GOSTREAM_LOGGING_LOGGER", "gostream"),
			Pretty:    GetBool("GOSTREAM_OUTPUT_PRETTY", false),
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

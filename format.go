package main

import (
	"bytes"
	"encoding/json"

	yaml "gopkg.in/yaml.v2"
)

func Json(event interface{}, pretty bool) (string, error) {
	b, err := json.Marshal(event)
	if err != nil {
		return "", err
	}

	if !pretty {
		return string(b), nil
	}

	var buf bytes.Buffer
	err = json.Indent(&buf, b, "", " ")
	if err != nil {
		return "", err
	}

	return string(buf.Bytes()) + "\n", nil
}

func Yaml(data interface{}) (string, error) {
	b, err := yaml.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

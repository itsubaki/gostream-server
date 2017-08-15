package config

import (
	"bytes"
	"encoding/json"

	cep "github.com/itsubaki/gocep"
)

func Json(event []cep.Event, pretty bool) (string, error) {
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

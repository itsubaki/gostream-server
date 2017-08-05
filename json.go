package main

import (
	"bytes"
	"encoding/json"

	cep "github.com/itsubaki/gocep"
)

func Json(event []cep.Event) (string, error) {
	b, err := json.Marshal(event)
	if err != nil {
		return "", err
	}
	var pretty bytes.Buffer
	err = json.Indent(&pretty, b, "", " ")
	if err != nil {
		return "", err
	}

	return string(pretty.Bytes()) + "\n", nil
}

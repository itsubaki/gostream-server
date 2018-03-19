package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"time"

	uuid "github.com/satori/go.uuid"
)

type LogEvent struct {
	ID      string
	Time    time.Time `json:"time"`
	Level   int       `json:"level"`
	Message string    `json:"message"`
}

type RequestID struct {
	ID string
}

func NewLogEvent(body io.ReadCloser) (LogEvent, error) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return LogEvent{}, err
	}

	var event LogEvent
	if err := json.Unmarshal(b, &event); err != nil {
		return LogEvent{}, err
	}
	event.ID = uuid.NewV4().String()

	return event, nil
}

package main

import "time"

type LogEvent struct {
	ID      string
	Time    time.Time `json:"time"`
	Level   int       `json:"level"`
	Message string    `json:"message"`
}

type RequestID struct {
	ID string
}

package main

import (
	"fmt"
	"log"

	"cloud.google.com/go/logging"
	"cloud.google.com/go/pubsub"
	cep "github.com/itsubaki/gocep"
	"golang.org/x/net/context"
)

type Output interface {
	Update(event []cep.Event)
	Close()
}

func NewOutput(config *Config) Output {
	ctx := context.Background()

	if config.Output == "pubsub" {
		pubsub, err := pubsub.NewClient(ctx, config.ProjectID)
		if err != nil {
			log.Println(err)
			return &OutputStdOut{config.Pretty}
		}

		topic := pubsub.Topic(config.Topic)
		return &OutputPubSub{ctx, pubsub, topic, config.Pretty}
	}

	if config.Output == "logging" {
		logging, err := logging.NewClient(ctx, config.ProjectID)
		if err != nil {
			log.Println(err)
			return &OutputStdOut{config.Pretty}
		}
		logger := logging.Logger(config.Logger)
		return &OutputLogging{logging, logger, config.Pretty}
	}

	return &OutputStdOut{config.Pretty}
}

type OutputLogging struct {
	client *logging.Client
	logger *logging.Logger
	pretty bool
}

func (o *OutputLogging) Update(event []cep.Event) {
	json, err := Json(event, o.pretty)
	if err != nil {
		log.Println(err)
		return
	}

	entry := logging.Entry{
		Severity: logging.Info,
		Payload:  json,
	}

	o.logger.Log(entry)
}

func (o *OutputLogging) Close() {
	o.logger.Flush()
	if err := o.client.Close(); err != nil {
		log.Println(err)
	}
}

type OutputPubSub struct {
	ctx    context.Context
	client *pubsub.Client
	topic  *pubsub.Topic
	pretty bool
}

func (o *OutputPubSub) Update(event []cep.Event) {
	json, err := Json(event, o.pretty)
	if err != nil {
		log.Println(err)
		return
	}

	res := o.topic.Publish(o.ctx, &pubsub.Message{
		Data: []byte(json),
	})

	msgID, err := res.Get(o.ctx)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(msgID)
}

func (o *OutputPubSub) Close() {
	if err := o.client.Close(); err != nil {
		log.Println(err)
	}
}

type OutputStdOut struct {
	pretty bool
}

func (o *OutputStdOut) Update(event []cep.Event) {
	json, err := Json(event, o.pretty)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(json)
}

func (o *OutputStdOut) Close() {
	//noop
}

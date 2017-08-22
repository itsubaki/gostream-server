package output

import (
	"log"

	"golang.org/x/net/context"

	"cloud.google.com/go/logging"
	"cloud.google.com/go/pubsub"
	cep "github.com/itsubaki/gocep"
	"github.com/itsubaki/gostream/config"
)

type Output interface {
	Update(event []cep.Event)
	Close()
}

func New(config *config.OutputConfig) Output {
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

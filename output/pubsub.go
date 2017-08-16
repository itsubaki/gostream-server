package output

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	"cloud.google.com/go/pubsub"
	cep "github.com/itsubaki/gocep"
	"github.com/itsubaki/gostream/config"
)

type OutputPubSub struct {
	ctx    context.Context
	client *pubsub.Client
	topic  *pubsub.Topic
	pretty bool
}

func (o *OutputPubSub) Update(event []cep.Event) {
	json, err := config.Json(event, o.pretty)
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
	err := o.client.Close()
	if err != nil {
		log.Println(err)
	}
}

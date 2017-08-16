package output

import (
	"log"

	"cloud.google.com/go/logging"
	cep "github.com/itsubaki/gocep"
	"github.com/itsubaki/gostream/config"
)

type OutputLogging struct {
	client *logging.Client
	logger *logging.Logger
	pretty bool
}

func (o *OutputLogging) Update(event []cep.Event) {
	json, err := config.Json(event, o.pretty)
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
	err := o.client.Close()
	if err != nil {
		log.Println(err)
	}
}

package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	log "github.com/tommzn/go-log"
	core "github.com/tommzn/hdb-datasource-core"
)

func NewProcessor(persistence Persistence, logger log.Logger) *EventProcessor {
	return &EventProcessor{
		persistence: persistence,
		logger:      logger,
	}
}

// Handle processes given SQS events.
func (processor *EventProcessor) Handle(ctx context.Context, sqsEvent events.SQSEvent) error {

	var err error
	for _, message := range sqsEvent.Records {
		processError := processor.processMessage(message)
		if processError != nil {
			processor.logger.Errorf("Unable to process event %s, reason: %s", message.MessageId, processError)
			if err == nil {
				err = processError
			}
		}
	}
	return err
}

func (processor *EventProcessor) processMessage(message events.SQSMessage) error {

	if attribute, ok := message.MessageAttributes[core.ORIGIN_QUEUE]; ok {
		queue := attribute.StringValue
		return processor.persistence.archiveMessage(message.MessageId, message.Body, *queue)
	} else {
		return fmt.Errorf("Attribute not found: %s", core.ORIGIN_QUEUE)
	}
}

package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	log "github.com/tommzn/go-log"
)

func New(logger log.Logger) *EventProcessor {
	return &EventProcessor{logger: logger}
}

// Handle processes given SQS events.
func (processor *EventProcessor) Handle(ctx context.Context, sqsEvent events.SQSEvent) error {
	return nil
}

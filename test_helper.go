package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-lambda-go/events"
	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
	core "github.com/tommzn/hdb-datasource-core"
)

func processorForTest() core.SqsEventProcessor {
	return NewProcessor(localPersistenceForTest(), loggerForTest())
}

func localPersistenceForTest() Persistence {
	return &localPersistence{
		storage: make(map[string]string),
	}
}

type localPersistence struct {
	storage map[string]string
}

func (persistence *localPersistence) archiveMessage(id, body, queue string) error {
	persistence.storage[id] = fmt.Sprintf("%s:%s", queue, body)
	return nil
}

func sqsEventForTest() events.SQSEvent {

	content, _ := ioutil.ReadFile("./fixtures/sqs_event.json")
	event := events.SQSEvent{}
	json.Unmarshal(content, &event)
	return event
}

// loggerForTest creates a new stdout logger for testing.
func loggerForTest() log.Logger {
	return log.NewLogger(log.Debug, nil, nil)
}

// loadConfigForTest loads test config.
func loadConfigForTest() config.Config {

	configFile, ok := os.LookupEnv("CONFIG_FILE")
	if !ok {
		configFile = "testconfig.yml"
	}
	configLoader := config.NewFileConfigSource(&configFile)
	config, _ := configLoader.Load()
	return config
}

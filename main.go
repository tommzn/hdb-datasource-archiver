package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
	secrets "github.com/tommzn/go-secrets"

	core "github.com/tommzn/hdb-datasource-core"
)

func main() {

	minion, err := bootstrap(nil)
	if err != nil {
		panic(err)
	}
	lambda.Start(minion.Handle)
}

// bootstrap loads config and creates a processor for SQS events.
func bootstrap(conf config.Config) (core.SqsEventProcessor, error) {

	var err error
	if conf == nil {
		conf, err = loadConfig()
		if err != nil {
			return nil, err
		}
	}

	persistence, err := NewS3Uploader(conf)
	if err != nil {
		return nil, err
	}

	secretsManager := newSecretsManager()
	logger := newLogger(conf, secretsManager)
	return NewProcessor(persistence, logger), nil
}

// loadConfig from config file.
func loadConfig() (config.Config, error) {

	configSource, err := config.NewS3ConfigSourceFromEnv()
	if err != nil {
		return nil, err
	}

	conf, err := configSource.Load()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// newSecretsManager retruns a new secrets manager from passed config.
func newSecretsManager() secrets.SecretsManager {
	return secrets.NewSecretsManager()
}

// newLogger creates a new logger from  passed config.
func newLogger(conf config.Config, secretsMenager secrets.SecretsManager) log.Logger {
	return log.NewLoggerFromConfig(conf, secretsMenager)
}

package main

import log "github.com/tommzn/go-log"

type EventProcessor struct {
	persistence Persistence
	logger      log.Logger
}

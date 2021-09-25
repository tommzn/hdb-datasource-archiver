package main

type Persistence interface {
	archiveMessage(id, body, queue string) error
}

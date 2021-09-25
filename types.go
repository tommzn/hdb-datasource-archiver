package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	log "github.com/tommzn/go-log"
)

type EventProcessor struct {
	persistence Persistence
	logger      log.Logger
}

type S3Uploader struct {
	bucket     *string
	path       *string
	awsConfig  *aws.Config
	s3Client   *s3.S3
	awsSession *session.Session
	s3Uploader *s3manager.Uploader
}

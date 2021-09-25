package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	config "github.com/tommzn/go-config"
)

// NewS3Uploader create a new S3 uploader.
func NewS3Uploader(conf config.Config) (Persistence, error) {

	bucket := conf.Get("aws.s3.bucket", nil)
	if bucket == nil {
		return nil, errors.New("No S3 bucket defined.")
	}
	path := conf.Get("aws.s3.path", nil)
	return &S3Uploader{
		awsConfig: &aws.Config{Region: conf.Get("aws.s3.region", config.AsStringPtr("eu-central-1"))},
		bucket:    bucket,
		path:      path,
	}, nil
}

func (uploader *S3Uploader) archiveMessage(id, body, queue string) error {

	_, err := uploader.getS3Uploader().Upload(
		&s3manager.UploadInput{
			Bucket: uploader.bucket,
			Key:    uploader.getObjectKey(id, queue),
			Body:   strings.NewReader(body),
		})
	return err
}

func (uploader *S3Uploader) s3Session() *session.Session {
	if uploader.awsSession == nil {
		uploader.awsSession = session.Must(session.NewSession(uploader.awsConfig))
	}
	return uploader.awsSession
}

func (uploader *S3Uploader) getS3Uploader() *s3manager.Uploader {
	if uploader.s3Uploader == nil {
		uploader.s3Uploader = s3manager.NewUploader(uploader.s3Session())
	}
	return uploader.s3Uploader
}

func (uploader *S3Uploader) getObjectKey(id, queue string) *string {
	objectKey := fmt.Sprintf("%s/%s", queue, id)
	if uploader.path != nil {
		objectKey = fmt.Sprintf("%s/%s", *uploader.path, objectKey)
	}
	return &objectKey
}

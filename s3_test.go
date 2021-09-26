package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
	config "github.com/tommzn/go-config"
)

type S3PersistenceTestSuite struct {
	suite.Suite
	conf config.Config
}

func TestS3PersistenceTestSuite(t *testing.T) {
	suite.Run(t, new(S3PersistenceTestSuite))
}

func (suite *S3PersistenceTestSuite) SetupTest() {
	suite.conf = loadConfigFromFile("fixtures/s3.testconfig.yml")
}

func (suite *S3PersistenceTestSuite) TestCreateS3Upload() {

	conf1 := loadConfigFromFile("fixtures/testconfig.yml")
	uploader1, err1 := NewS3Uploader(conf1)
	suite.Nil(err1)
	suite.NotNil(uploader1)

	conf2 := loadConfigFromFile("fixtures/incomplete.testconfig.yml")
	uploader2, err2 := NewS3Uploader(conf2)
	suite.NotNil(err2)
	suite.Nil(uploader2)
}

func (suite *S3PersistenceTestSuite) TestUpload() {

	skipCI(suite.T())

	uploader, _ := NewS3Uploader(suite.conf)
	suite.Nil(uploader.archiveMessage("Test-Message-1", "xXx", "test-queue"))
}

package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	config "github.com/tommzn/go-config"
)

type ProcessorTestSuite struct {
	suite.Suite
	conf config.Config
}

func TestProcessorTestSuite(t *testing.T) {
	suite.Run(t, new(ProcessorTestSuite))
}

func (suite *ProcessorTestSuite) SetupTest() {
	suite.conf = loadConfigForTest()
}

func (suite *ProcessorTestSuite) TestProcessEvents() {

	provessor := processorForTest()
	event := sqsEventForTest()

	err := provessor.Handle(context.Background(), event)
	suite.Nil(err)
}

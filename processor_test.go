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

func (suite *ProcessorTestSuite) TestBootstrapProcessor() {

	processor, err := bootstrap(suite.conf)
	suite.NotNil(processor)
	suite.Nil(err)
}

func (suite *ProcessorTestSuite) TestProcessEvents() {

	processor := processorForTest()
	event := sqsEventForTest()

	err := processor.Handle(context.Background(), event)
	suite.Nil(err)
}

func (suite *ProcessorTestSuite) TestProcessEventsWithoutExpectedAttribute() {

	processor := processorForTest()
	event := sqsEventWithoutAttributeForTest()

	err := processor.Handle(context.Background(), event)
	suite.NotNil(err)
}

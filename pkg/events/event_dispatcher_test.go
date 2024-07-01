package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TestEvent struct{
	Name string
	Payload interface{}
}

type TestEventHandler struct{}

func (e* TestEvent) GetPayload() interface{}{
	return e.Payload
}

func (e* TestEvent) GetDateTime() time.Time{
	return time.Now()
}

func (h *TestEventHandler) Handle(event EventInterface){

}

type EventDispatcherTestSuite struct{
	suite.Suite
	event TestEvent
	event2 TestEvent
	handler TestEventHandler
	handler2 TestEventHandler
	EventDispatcher EventDispatcher
}

func TestSuite(t *testing.T){
	suite.Run(t, new(EventDispatcherTestSuite))
}
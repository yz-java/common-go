package rabbitmq

import (
	"testing"
	"common-go/log"
)

func Test_RabbitMQ(t *testing.T) {
	Connect("amqp://admin:admin@192.168.1.160:5672")
	CreateExchange("test_exchage")
	go Receive("test",nil)
	err := Push("test_exchage", "test", "haha")
	log.Logger.Error(err)
	sys := make(chan bool)
	<-sys
}

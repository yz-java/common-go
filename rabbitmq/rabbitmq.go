package rabbitmq

import (
	"github.com/streadway/amqp"
	"common-go/log"
	"bytes"
)

var (
	conn *amqp.Connection
	channel *amqp.Channel
	queueMap map[string]*amqp.Queue = make(map[string]*amqp.Queue)
)


func Connect(mqurl string) {
	var err error
	conn, err = amqp.Dial(mqurl)
	if err != nil {
		panic(err)
	}
	channel, err = conn.Channel()
	if err != nil {
		panic(err)
	}

}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}

func createQueue(queueName string)(amqp.Queue,error) {
	return channel.QueueDeclare(
		queueName,    // name  有名字！
		true,         // durable  持久性的,如果事前已经声明了该队列，不能重复声明
		false,        // delete when unused
		false,        // exclusive 如果是真，连接一断开，队列删除
		false,        // no-wait
		nil,          // arguments
	)
}

func CreateExchange(exchangeName string) {
	err := channel.ExchangeDeclare(exchangeName, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
}

func Push(exchange,queueName,msgContent string) error  {
	if channel == nil {
		panic("rabbitmq channel is nil")
	}
	//if queueMap[queueName] == nil {
	//	queue,err:=createQueue(queueName)
	//	if err != nil {
	//		log.Logger.Error(err)
	//		return err
	//	}
	//	queueMap[queueName]=&queue
	//}
	return channel.Publish(exchange, queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})
}


func Receive(queueName string,f func(string)) {
	if channel == nil {
		panic("rabbitmq channel is nil")
	}
	msgs, err := channel.Consume(queueName,"",true, false, false, false, nil)
	if err != nil {
		log.Logger.Error("consummer register fail ---- ",err)
		return
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := BytesToString(&(d.Body))
			log.Logger.Info("receve msg is :%s \n", *s)
			f(*s)
		}
	}()
	<-forever
}
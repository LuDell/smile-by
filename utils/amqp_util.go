package utils

import (
	"fmt"
	"github.com/streadway/amqp"
)

var MQConn *amqp.Connection

func init()  {
	conn, err := amqp.Dial("amqp://admin:1q2w3e4r@144.34.237.147:5672/")
	if err != nil {
		fmt.Errorf("mq connection fail %s",err)
	}

	MQConn = conn
}

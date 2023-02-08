package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("GO RabbotMQ Tutorial")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer conn.Close()
	fmt.Println("successfully Connected to RabbitMq")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)

	}
	q, err := ch.QueueDeclare("TestChannel", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = ch.Publish("", "TestChannel", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hi publishing to the queue"),
	})
	fmt.Println(q)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("successfully published to the mq")
}

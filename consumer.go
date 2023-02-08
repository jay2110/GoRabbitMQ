package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()
	msg, err := ch.Consume("TestChannel", "", true, false, false, false, nil)

	forever := make(chan bool)
	go func() {
		for d := range msg {
			fmt.Printf("Msg Received %s\n", d.Body)
		}
	}()
	fmt.Println("successfully read msg from RabbitMq instance")
	fmt.Print("[*]- waiting for the message")
	<-forever
}

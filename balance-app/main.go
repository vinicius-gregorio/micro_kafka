package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	fmt.Println("Hello, World!")

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(fmt.Errorf("failed to create consumer: %v", err))
	}
	defer c.Close()

	for {
		err := subscribeToBalancesTopic(c)
		if err != nil {
			fmt.Printf("Failed to subscribe to balances topic: %v\n", err)
			time.Sleep(5 * time.Second) // Retry after 5 seconds
		} else {
			break // Exit the loop if successfully subscribed
		}
	}

	run := true
	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !isTimeoutError(err) {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func subscribeToBalancesTopic(c *kafka.Consumer) error {
	topics := []string{"balances"}
	err := c.SubscribeTopics(topics, nil)
	return err
}

func isTimeoutError(err error) bool {
	kafkaErr, ok := err.(kafka.Error)
	return ok && kafkaErr.IsTimeout()
}

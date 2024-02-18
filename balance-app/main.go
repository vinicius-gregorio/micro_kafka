package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello, World!")

	// configMap := ckafka.ConfigMap{
	// 	"bootstrap.servers": "kafka:29092",
	// 	"group.id":          "wallet",
	// }
	// kafkaConsumer := kafka.NewConsumer(&configMap, []string{"balance"})
	// kafkaConsumer.ConfigMap = &configMap
}

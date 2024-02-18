package handler

import (
	"fmt"
	"sync"

	"github.com/vinicius-gregorio/fc_walletcore/pkg/events"
	"github.com/vinicius-gregorio/fc_walletcore/pkg/kafka"
)

type TransactionCreatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewTransactionCreatedKafkaHandler(kafka *kafka.Producer) *TransactionCreatedKafkaHandler {
	return &TransactionCreatedKafkaHandler{Kafka: kafka}
}

func (h *TransactionCreatedKafkaHandler) Handle(msg events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	h.Kafka.Publish(msg, nil, "transactions")
	fmt.Println("TransactionCreatedKafkaHandler: ", msg)
}

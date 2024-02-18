package handler

import (
	"fmt"
	"sync"

	"github.com/vinicius-gregorio/fc_walletcore/pkg/events"
	"github.com/vinicius-gregorio/fc_walletcore/pkg/kafka"
)

type UpdateBalanceKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewUpdateBalanceKafkaHandler(kafka *kafka.Producer) *UpdateBalanceKafkaHandler {
	return &UpdateBalanceKafkaHandler{Kafka: kafka}
}

func (h *UpdateBalanceKafkaHandler) Handle(msg events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	h.Kafka.Publish(msg, nil, "balances")
	fmt.Println("UpdateBalanceKafkaHandler: ", msg)
}

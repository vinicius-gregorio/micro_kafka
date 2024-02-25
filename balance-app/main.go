package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vinicius-gregorio/balance-api/database"
	"github.com/vinicius-gregorio/balance-api/entity"
	"github.com/vinicius-gregorio/balance-api/web_client"
)

func main() {
	fmt.Println("Balances APP started")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", "root", "mysql", "3306", "wallet_balance"))

	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer db.Close()

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

	// Start the web server
	balanceDB := database.NewBalanceDB(db)
	if balanceDB == nil {
		log.Fatal("Failed to initialize balance database")
	}
	server := web_client.NewServer(balanceDB)
	go server.Start()

	run := true
	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			err := updateBalance(string(msg.Value), db)
			if err != nil {
				fmt.Printf("Error updating balance: %v\n", err)
			}
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

type MessagePayload struct {
	Name    string      `json:"Name"`
	Payload PayloadData `json:"Payload"`
}
type PayloadData struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

func updateBalance(message string, db *sql.DB) error {
	// Parse the message payload
	var msgPayload MessagePayload
	if err := json.Unmarshal([]byte(message), &msgPayload); err != nil {
		return err
	}

	// Check if the message is of interest
	if msgPayload.Name != "balance.updated" {
		fmt.Printf("Ignoring message %s\n", msgPayload.Name)
		return nil
	}

	// Initialize BalanceDB
	balanceDB := database.NewBalanceDB(db)
	if balanceDB == nil {
		return errors.New("failed to initialize balance database")
	}

	fmt.Printf("Updating balances for accounts %s and %s\n", msgPayload.Payload.AccountIDFrom, msgPayload.Payload.AccountIDTo)

	// Update balance for AccountIDFrom
	if err := updateSingleBalance(msgPayload.Payload.AccountIDFrom, msgPayload.Payload.BalanceAccountIDFrom, balanceDB); err != nil {
		return err
	}

	// Update balance for AccountIDTo
	if err := updateSingleBalance(msgPayload.Payload.AccountIDTo, msgPayload.Payload.BalanceAccountIDTo, balanceDB); err != nil {
		return err
	}

	return nil
}

func updateSingleBalance(accountID string, newAmount float64, balanceDB *database.BalanceDB) error {
	// Check if balanceDB is nil
	if balanceDB == nil {
		return errors.New("balanceDB is nil")
	}

	// Check if balance for the account exists
	balance, err := balanceDB.FindByID(accountID)
	if err != nil {
		return err // Return error from FindByID
	}

	// If balance doesn't exist, create a new entry
	if balance == nil {
		newBalance := entity.NewBalance(&entity.Balance{AccountID: accountID, Amount: newAmount})
		return balanceDB.Save(newBalance)
	}

	// If balance exists, update the amount
	balance.Amount = newAmount
	return balanceDB.UpdateBalance(balance)
}

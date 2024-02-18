package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vinicius-gregorio/fc_walletcore/internal/database"
	"github.com/vinicius-gregorio/fc_walletcore/internal/event"
	"github.com/vinicius-gregorio/fc_walletcore/internal/event/handler"
	"github.com/vinicius-gregorio/fc_walletcore/internal/usecase/create_account"
	"github.com/vinicius-gregorio/fc_walletcore/internal/usecase/create_client"
	"github.com/vinicius-gregorio/fc_walletcore/internal/usecase/create_transaction"
	"github.com/vinicius-gregorio/fc_walletcore/internal/web"
	"github.com/vinicius-gregorio/fc_walletcore/internal/web/webserver"
	"github.com/vinicius-gregorio/fc_walletcore/pkg/events"
	"github.com/vinicius-gregorio/fc_walletcore/pkg/kafka"
	"github.com/vinicius-gregorio/fc_walletcore/pkg/uow"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", "root", "mysql", "3306", "wallet"))

	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	evtDis := events.NewEventDispatcher()
	evtDis.Register("transaction.created", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	evtDis.Register("balance.updated", handler.NewUpdateBalanceKafkaHandler(kafkaProducer))

	trCreatedEvt := event.NewTransactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)
	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUsecase := create_client.NewCreateClientUsecase(clientDb)
	createAccUsecase := create_account.NewCreateAccountUsecase(accountDb, clientDb)
	createTransactionUsecase := create_transaction.NewCreateTransactionUsecase(uow, evtDis, trCreatedEvt, balanceUpdatedEvent)

	clientHandler := web.NewWebClientHandler(*createClientUsecase)
	accountHandler := web.NewWebAccountHandler(*createAccUsecase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUsecase)

	webserver := webserver.NewWebServer(":8080")

	webserver.AddHandler("/client", clientHandler.CreateClient)
	webserver.AddHandler("/account", accountHandler.CreateAccount)
	webserver.AddHandler("/transaction", transactionHandler.CreateTransaction)

	fmt.Println("Server running on port 8080")
	webserver.Start()

}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/krotovkk/bwgtransaction/config"
	"github.com/krotovkk/bwgtransaction/internal/core/services/clientservice"
	"github.com/krotovkk/bwgtransaction/internal/handlers/clienthandler"
	"github.com/krotovkk/bwgtransaction/internal/repository/clientrepository"
	"github.com/krotovkk/bwgtransaction/internal/repository/historyrepository"
)

func main() {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s connect_timeout=%d sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName, config.ConnectTimeout)

	conn, err := pgx.Connect(context.Background(), psqlConn)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	clientRepo := clientrepository.NewClientRepository(conn)

	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(config.Brokers, cfg)

	historyRepo := historyrepository.NewHistoryRepository(producer)

	clientService := clientservice.NewClientService(&clientservice.Options{
		ClientRepo:  clientRepo,
		HistoryRepo: historyRepo,
	})
	handler := clienthandler.NewHttpHandler(clientService)

	router := gin.New()
	router.GET("/clients/:id", handler.GetBalance)
	router.POST("/client/:id/:diff", handler.ChangeBalance)
	router.POST("/client/create", handler.Create)

	router.Run(":8000")
}

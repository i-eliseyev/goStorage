package main

import (
	"goStorage/internal/server"
	"goStorage/internal/transaction_logger"
	"log"
)

func main() {
	if err := transaction_logger.InitializeTransactionLog(); err != nil {
		log.Fatal(err)
	}
	log.Fatal(server.Start())
}

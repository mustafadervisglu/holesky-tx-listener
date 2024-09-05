package main

import (
	"fmt"
	"holesxy-tx-listener/internal/config"
	"holesxy-tx-listener/internal/db"
	"holesxy-tx-listener/internal/ethereum"
	model "holesxy-tx-listener/internal/models"
	"log"
)

func main() {
	ethUrl := config.LoadConfig().Ethereum
	dbDns := config.LoadConfig().Database

	newClient, err := ethereum.NewClient(ethUrl)

	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	database, err := db.Connect(dbDns)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	if err := database.AutoMigrate(&model.Block{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	fmt.Println("Connected to db:", database)
	fmt.Println("Connected to Ethereum client:", &newClient)
	listener := ethereum.NewListener(newClient, database)
	listener.Start()
}

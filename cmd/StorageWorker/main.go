package main

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/controller"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	controller.StartInsertListener(IoC.InsertKafkaController)
}

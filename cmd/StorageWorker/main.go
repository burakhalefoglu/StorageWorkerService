package main

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/IoC/golobby"
	"StorageWorkerService/internal/controller"
	"StorageWorkerService/internal/controller/kafka"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	IoC.InjectContainers(golobby.InjectionConstructor())

	var waitGroup sync.WaitGroup
	IController.StartInsertListener(&waitGroup, KafkaController.InsertControllerConstructor())
	waitGroup.Wait()
}

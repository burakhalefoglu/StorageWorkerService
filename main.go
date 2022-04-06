package main

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/IoC/golobby"
	IController "StorageWorkerService/internal/controller"
	KafkaController "StorageWorkerService/internal/controller/kafka"
	"StorageWorkerService/pkg/helper"
	"log"
	"runtime"
	"sync"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/joho/godotenv"
)

func main() {
	defer helper.DeleteHealthFile()
	logger.Log.App = "StorageWorkerService"
	runtime.MemProfileRate = 0
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

package main

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/controller"
)

func main() {



	controller.StartInsertListener(IoC.InsertKafkaController)
}

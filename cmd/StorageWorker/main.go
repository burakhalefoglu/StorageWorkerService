package main

import (
	InsertDataController "StorageWorkerService/internal/controller"
	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
)

func main() {

	mongodb.ConnectMongoDb()
	InsertDataController.StartListening()
}

package AdvEventDataManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
)

type AdvEventDataModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	LevelName string
	LevelIndex int32
	AdvType string
	InMinutes float32
	TrigerdTime time.Time
}


func InsertAdvEventDataModel(reader *kafka.Reader, m kafka.Message) {

		advEventDataModel := AdvEventDataModel{}
		jsonParser.DecodeJson(m.Value, &advEventDataModel)
		log.Println(advEventDataModel)
	
	_, err:= mongodb.AddCollection(m.Topic, advEventDataModel)
		if(err == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}
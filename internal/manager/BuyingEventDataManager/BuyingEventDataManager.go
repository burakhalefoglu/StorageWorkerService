package BuyingEventDataManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
)

type BuyingEventDataModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	LevelName string
	LevelIndex int32
	ProductType string
	InWhatMinutes float32
	TrigerdTime time.Time
}


func InsertBuyingEventDataModel(reader *kafka.Reader, m kafka.Message) {

	buyingEventDataModel := BuyingEventDataModel{}
		jsonParser.DecodeJson(m.Value, &buyingEventDataModel)
		log.Println(buyingEventDataModel)
	
	_, err:= mongodb.AddCollection(m.Topic, buyingEventDataModel)
		if(err == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}
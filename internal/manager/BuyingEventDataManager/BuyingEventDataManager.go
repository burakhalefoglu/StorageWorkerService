package BuyingEventDataManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
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

type ClientDataModel struct{

	ClientId string
	ProjectId string
	IsPaidClient int
	CreatedAt time.Time
	PaidTime time.Time
}

func InsertBuyingEventDataModel(reader *kafka.Reader, m kafka.Message) {

	buyingEventDataModel := BuyingEventDataModel{}
		jsonParser.DecodeJson(m.Value, &buyingEventDataModel)
		log.Println(buyingEventDataModel)
	
	_, err:= mongodb.AddCollection(m.Topic, buyingEventDataModel)

		_, updateErr:= mongodb.UpdateCollection(m.Topic,
			 bson.M{"ProjectId": buyingEventDataModel.ProjectId},
			 bson.D{
				{"$set", bson.D{{"IsPaidClient", 1}}},
			})

		if(err == nil && updateErr == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
	//TODO: burada kullanıcı paid olarak işaretlenecek.

}
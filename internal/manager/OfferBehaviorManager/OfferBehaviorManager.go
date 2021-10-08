package OfferBehaviorManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
)

type OfferBehaviorModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	OfferId int
	IsBuyOffer int
	DateTime time.Time
}


func InsertOfferBehaviorModel(reader *kafka.Reader, m kafka.Message) {

	offerBehaviorModel := OfferBehaviorModel{}
		jsonParser.DecodeJson(m.Value, &offerBehaviorModel)
		log.Println(offerBehaviorModel)
	
	_, err:= mongodb.AddCollection(m.Topic, offerBehaviorModel)
		if(err == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}
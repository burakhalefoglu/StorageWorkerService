package LocationDataManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
)

type LocationDataModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	Continent string
	Country string
	City string
	Query string
	Region string
	Org string
	CreatedAt time.Time
}


func InsertLocationDataModel(reader *kafka.Reader, m kafka.Message) {

	locationDataModel := LocationDataModel{}
		jsonParser.DecodeJson(m.Value, &locationDataModel)
		log.Println(locationDataModel)
	
	_, err:= mongodb.AddCollection(m.Topic, locationDataModel)
		if(err == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}
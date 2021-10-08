package GameSessionEveryLoginDataManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
)

type GameSessionEveryLoginDataModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	SessionTimeMinute float32
	SessionStartTime time.Time
	SessionFinishTime time.Time
}


func InsertGameSessionEveryLoginDataModel(reader *kafka.Reader, m kafka.Message) {

	gameSessionEveryLoginDataModel := GameSessionEveryLoginDataModel{}
		jsonParser.DecodeJson(m.Value, &gameSessionEveryLoginDataModel)
		log.Println(gameSessionEveryLoginDataModel)
	
	_, err:= mongodb.AddCollection(m.Topic, gameSessionEveryLoginDataModel)
		if(err == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}
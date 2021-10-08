package LevelBaseSessionDataManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
)

type LevelBaseSessionDataModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	LevelName string
	LevelIndex int
	SessionTimeMinute float32
	SessionStartTime time.Time
	SessionFinishTime time.Time
}


func InsertLevelBaseSessionDataModel(reader *kafka.Reader, m kafka.Message) {

	levelBaseSessionDataModel := LevelBaseSessionDataModel{}
		jsonParser.DecodeJson(m.Value, &levelBaseSessionDataModel)
		log.Println(levelBaseSessionDataModel)
	
	_, err:= mongodb.AddCollection(m.Topic, levelBaseSessionDataModel)
		if(err == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}
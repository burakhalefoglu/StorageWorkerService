package ScreenSwipeDataManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
)

type ScreenSwipeDataModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	SwipeDirection int
	StartLocX float32
	StartLocY float32
	FinishLocX float32
	FinishLocY float32
	LevelName string
	LevelIndex int
	CreatedAt time.Time
}


func InsertScreenSwipeDataModel(reader *kafka.Reader, m kafka.Message) {

	screenSwipeDataModel := ScreenSwipeDataModel{}
		jsonParser.DecodeJson(m.Value, &screenSwipeDataModel)
		log.Println(screenSwipeDataModel)
	
	_, err:= mongodb.AddCollection(m.Topic, screenSwipeDataModel)
		if(err == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}
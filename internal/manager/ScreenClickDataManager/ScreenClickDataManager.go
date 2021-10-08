package ScreenClickDataManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
)

type ScreenClickDataModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	StartLocX float32
	StartLocY float32
	FinishLocX float32
	FinishLocY float32
	LevelName string
	LevelIndex int
	TabCount int
	FingerID int
	CreatedAt time.Time
}


func InsertScreenClickDataModel(reader *kafka.Reader, m kafka.Message) {

	screenClickDataModel := ScreenClickDataModel{}
		jsonParser.DecodeJson(m.Value, &screenClickDataModel)
		log.Println(screenClickDataModel)
	
	_, err:= mongodb.AddCollection(m.Topic, screenClickDataModel)
		if(err == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}
package EnemyBaseWithLevelFailDataManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
)

type EnemyBaseWithLevelFailDataModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	LevelName string
	LevelIndex int
	DiyingTimeAfterLevelStarting int
	FailLocationX float32
	FailLocationY float32
	FailLocationZ float32
	DateTime time.Time
}


func InsertEnemyBaseWithLevelFailDataModel(reader *kafka.Reader, m kafka.Message) {

	enemyBaseWithLevelFailDataModel := EnemyBaseWithLevelFailDataModel{}
		jsonParser.DecodeJson(m.Value, &enemyBaseWithLevelFailDataModel)
		log.Println(enemyBaseWithLevelFailDataModel)
	
	_, err:= mongodb.AddCollection(m.Topic, enemyBaseWithLevelFailDataModel)
		if(err == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}
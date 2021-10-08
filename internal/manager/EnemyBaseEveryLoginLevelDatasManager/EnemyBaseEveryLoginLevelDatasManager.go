package EnemyBaseEveryLoginLevelDatasManager

import (
	"context"
	"log"
	"time"

	mongodb "StorageWorkerService/internal/dataAccess/mongodb"
	jsonParser "StorageWorkerService/internal/helper/jsonParser"

	"github.com/segmentio/kafka-go"
)

type EnemyBaseEveryLoginLevelDatasModel struct{

	ClientId string
	ProjectId string
	CustomerId string
	LevelName string
	LevelIndex int
	PlayingTime int
	AverageScores int
	IsDead int
	TotalPowerUsage int
	DateTime time.Time
}


func InsertEnemyBaseEveryLoginLevelDatasModel(reader *kafka.Reader, m kafka.Message) {

	enemyBaseEveryLoginLevelDatasModel := EnemyBaseEveryLoginLevelDatasModel{}
		jsonParser.DecodeJson(m.Value, &enemyBaseEveryLoginLevelDatasModel)
		log.Println(enemyBaseEveryLoginLevelDatasModel)
	
	_, err:= mongodb.AddCollection(m.Topic, enemyBaseEveryLoginLevelDatasModel)
		if(err == nil) {
				if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}
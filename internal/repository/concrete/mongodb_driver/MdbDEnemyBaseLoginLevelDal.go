package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbDEnemyBaseLoginLevelDal struct {
	Client *mongo.Client
}

func MDbDEnemyBaseLoginLevelDalConstruct() *mDbDEnemyBaseLoginLevelDal {
	return &mDbDEnemyBaseLoginLevelDal{Client: mongodb.GetMongodbClient()}
}

func (m *mDbDEnemyBaseLoginLevelDal) Add(data *model.EnemyBaseLoginLevelModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection("enemyBaseLoginLevels")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"DateTime", data.DateTime},
		{"LevelName", data.LevelName},
		{"LevelIndex", data.LevelIndex},
		{"PlayingTime", data.PlayingTime},
		{"AverageScores", data.AverageScores},
		{"IsDead", data.IsDead},
		{"TotalPowerUsage", data.TotalPowerUsage},
	})
	if err != nil {
		return err
	}
	return nil
}

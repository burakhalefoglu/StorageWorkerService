package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbDEnemyBaseLevelFailDal struct {
	Client *mongo.Client
}

func MDbDEnemyBaseLevelFailDalConstructor() *mDbDEnemyBaseLevelFailDal {
	return &mDbDEnemyBaseLevelFailDal{Client: mongodb.GetMongodbClient()}
}

func (m *mDbDEnemyBaseLevelFailDal) Add(data *model.EnemyBaseLevelFailModel) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("Client").Collection("EnemyBaseLevelFail")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"DateTime", data.DateTime},
		{"LevelName", data.LevelName},
		{"LevelIndex", data.LevelIndex},
		{"DyingTimeAfterLevelStarting", data.DiyingTimeAfterLevelStarting},
		{"FailLocationX", data.FailLocationX},
		{"FailLocationY", data.FailLocationY},
		{"FailLocationZ", data.FailLocationZ},
	})

	if err != nil {
		return err
	}
	return nil
}

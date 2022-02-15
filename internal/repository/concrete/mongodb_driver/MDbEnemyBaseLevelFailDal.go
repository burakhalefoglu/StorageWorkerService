package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbEnemyBaseLevelFailDal struct {
	Client *mongo.Client
	Table  string
}

func NewMDbEnemyBaseLevelFailDal(Table string) *mDbEnemyBaseLevelFailDal {
	return &mDbEnemyBaseLevelFailDal{Client: mongodb.GetMongodbClient(),
		Table: Table}
}

func (m *mDbEnemyBaseLevelFailDal) Add(data *model.EnemyBaseLevelFailModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection(m.Table)
	var _, err = collection.InsertOne(ctx, bson.D{
		{"Id", data.Id},
		{"Status", data.Status},
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"DateTime", data.DateTime},
		{"LevelName", data.LevelName},
		{"LevelIndex", data.LevelIndex},
		{"DyingTimeAfterLevelStarting", data.FailTimeAfterLevelStarting},
		{"FailLocationX", data.FailLocationX},
		{"FailLocationY", data.FailLocationY},
		{"FailLocationZ", data.FailLocationZ},
	})

	if err != nil {
		return err
	}
	return nil
}

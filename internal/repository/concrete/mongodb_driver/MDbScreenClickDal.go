package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbScreenClickDal struct {
	Client *mongo.Client
	Table  string
}

func NewMDbScreenClickDal(Table string) *mDbScreenClickDal {
	return &mDbScreenClickDal{Client: mongodb.GetMongodbClient(),
		Table: Table}
}

func (m *mDbScreenClickDal) Add(data *model.ScreenClickModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection(m.Table)
	var _, err = collection.InsertOne(ctx, bson.D{
		{"Id", data.Id},
		{"Status", data.Status},
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"CreatedAt", data.CreatedAt},
		{"FingerID", data.FingerId},
		{"TabCount", data.TabCount},
		{"StartLocX", data.StartLocX},
		{"StartLocY", data.StartLocY},
		{"FinishLocX", data.FinishLocX},
		{"FinishLocY", data.FinishLocY},
		{"LevelName", data.LevelName},
		{"LevelIndex", data.LevelIndex},
	})
	if err != nil {
		return err
	}
	return nil
}

package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbDScreenClickDal struct {
	Client *mongo.Client
}

func MDbDScreenClickDalConstructor() *mDbDScreenClickDal {
	return &mDbDScreenClickDal{Client: mongodb.GetMongodbClient()}
}

func (m *mDbDScreenClickDal) Add(data *model.ScreenClickModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection("screenClicks")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"CreatedAt", data.CreatedAt},
		{"FingerID", data.FingerID},
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

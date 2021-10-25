package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MDbDScreenClickDal struct {
	Client *mongo.Client
}

func (m *MDbDScreenClickDal) Add(data *model.ScreenClickModel) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("Client").Collection("ScreenClick")
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

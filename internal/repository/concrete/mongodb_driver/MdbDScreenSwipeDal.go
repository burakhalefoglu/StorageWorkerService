package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MDbDScreenSwipeDal struct {
	Client *mongo.Client
}

func (m *MDbDScreenSwipeDal) Add(data *model.ScreenSwipeModel) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("Client").Collection("ScreenSwipe")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"CreatedAt", data.CreatedAt},
		{"SwipeDirection", data.SwipeDirection},
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
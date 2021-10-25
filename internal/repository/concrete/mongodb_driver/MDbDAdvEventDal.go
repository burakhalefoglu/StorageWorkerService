package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MDbDAdvEventDal struct {
	Client *mongo.Client
}

func (m *MDbDAdvEventDal) Add(data *model.AdvEventDataModel) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("Client").Collection("AdvEvent")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"AdvType", data.AdvType},
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"InMinutes", data.InMinutes},
		{"LevelIndex", data.LevelIndex},
		{"LevelName", data.LevelName},
		{"TriggeredTime", data.TriggeredTime},
		})
	if err != nil {
		return err
	}
	return nil
}
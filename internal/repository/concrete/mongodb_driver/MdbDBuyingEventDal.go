package mongodb_driver

import (
"StorageWorkerService/internal/model"
"context"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/mongo"
"time"
)

type MDbDBuyingEventDal struct {
	Client *mongo.Client
}

func (m *MDbDBuyingEventDal) Add(data *model.BuyingEventModel) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("Client").Collection("BuyingEvent")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"AdvType", data.ProductType},
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"LevelName", data.LevelName},
		{"InWhatMinutes", data.InWhatMinutes},
		{"TriggeredTime", data.TriggeredTime},
	})
	if err != nil {
		return err
	}
	return nil
}
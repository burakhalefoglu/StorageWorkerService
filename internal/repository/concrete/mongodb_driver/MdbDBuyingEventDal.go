package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbDBuyingEventDal struct {
	Client *mongo.Client
}

func MDbDBuyingEventDalConstructor() *mDbDBuyingEventDal {
	return &mDbDBuyingEventDal{Client: mongodb.GetMongodbClient()}
}

func (m *mDbDBuyingEventDal) Add(data *model.BuyingEventModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection("buyingEvents")
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

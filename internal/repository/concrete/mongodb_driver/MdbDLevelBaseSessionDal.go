package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MDbDLevelBaseSessionDal struct {
	Client *mongo.Client
}

func (m *MDbDLevelBaseSessionDal) Add(data *model.LevelBaseSessionModel) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("Client").Collection("LevelBaseSession")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"LevelName", data.LevelName},
		{"LevelIndex", data.LevelIndex},
		{"SessionTimeMinute", data.SessionTimeMinute},
		{"SessionStartTime", data.SessionStartTime},
		{"SessionFinishTime", data.SessionFinishTime},

	})
	if err != nil {
		return err
	}
	return nil
}
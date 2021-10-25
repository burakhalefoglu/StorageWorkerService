package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MDbDGameSessionDal struct {
	Client *mongo.Client
}

func (m *MDbDGameSessionDal) Add(data *model.GameSessionModel) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("Client").Collection("GameSession")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"SessionTimeMinute", data.SessionTimeMinute},
		{"SessionStartTime", data.SessionStartTime},
		{"SessionFinishTime", data.SessionFinishTime},



	})
	if err != nil {
		return err
	}
	return nil
}

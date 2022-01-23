package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbDGameSessionDal struct {
	Client *mongo.Client
}

func MDbDGameSessionDalConstructor() *mDbDGameSessionDal {
	return &mDbDGameSessionDal{Client: mongodb.GetMongodbClient()}
}

func (m *mDbDGameSessionDal) Add(data *model.GameSessionModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection("gameSessions")
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

package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbDLevelBaseSessionDal struct {
	Client *mongo.Client
}

func MDbDLevelBaseSessionDalConstructor() *mDbDLevelBaseSessionDal {
	return &mDbDLevelBaseSessionDal{Client: mongodb.GetMongodbClient()}
}

func (m *mDbDLevelBaseSessionDal) Add(data *model.LevelBaseSessionModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection("levelBaseSessions")
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

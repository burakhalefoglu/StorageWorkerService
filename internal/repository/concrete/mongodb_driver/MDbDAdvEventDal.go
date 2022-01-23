package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"StorageWorkerService/pkg/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbDAdvEventDal struct {
	Client *mongo.Client
	Logger *logger.ILog
}

func MDbDAdvEventDalConstructor() *mDbDAdvEventDal {
	return &mDbDAdvEventDal{Client: mongodb.GetMongodbClient()}
}

func (m *mDbDAdvEventDal) Add(data *model.AdvEventDataModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection("advEvents")
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

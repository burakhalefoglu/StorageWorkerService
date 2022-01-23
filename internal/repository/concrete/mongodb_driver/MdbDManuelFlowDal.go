package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbDManuelFlowDal struct {
	Client *mongo.Client
}

func MDbDManuelFlowDalConstructor() *mDbDManuelFlowDal {
	return &mDbDManuelFlowDal{Client: mongodb.GetMongodbClient()}
}

func (m *mDbDManuelFlowDal) Add(data *model.ManuelFlowModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection("manuelFlows")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"DifficultyLevel", data.DifficultyLevel},
		{"DateTime", data.DateTime},
	})
	if err != nil {
		return err
	}
	return nil
}

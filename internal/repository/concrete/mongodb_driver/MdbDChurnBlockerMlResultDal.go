package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MDbDChurnBlockerMlResultDal struct {
	Client *mongo.Client
}

func (m *MDbDChurnBlockerMlResultDal) Add(data *model.ChurnBlockerMlResultModel) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("Client").Collection("ChurnBlockerMlResult")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"ModelType", data.ModelType},
		{"ModelResult", data.ModelResult},
		{"DateTime", data.DateTime},

	})
	if err != nil {
		return err
	}
	return nil
}
package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MDbDChurnPredictionMlResultDal struct {
	Client *mongo.Client
}

func (m *MDbDChurnPredictionMlResultDal) Add(data *model.ChurnPredictionMlResultModel) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("Client").Collection("ChurnPrediction")
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

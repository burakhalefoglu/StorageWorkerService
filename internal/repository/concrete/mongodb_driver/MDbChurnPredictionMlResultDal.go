package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbChurnPredictionMlResultDal struct {
	Client *mongo.Client
	Table  string
}

func NewMDbChurnPredictionMlResultDal(Table string) *mDbChurnPredictionMlResultDal {
	return &mDbChurnPredictionMlResultDal{Client: mongodb.GetMongodbClient(),
		Table: Table}
}

func (m *mDbChurnPredictionMlResultDal) Add(data *model.ChurnPredictionMlResultModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection(m.Table)
	var _, err = collection.InsertOne(ctx, bson.D{
		{"Id", data.Id},
		{"Status", data.Status},
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

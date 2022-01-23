package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbDChurnBlockerMlResultDal struct {
	Client *mongo.Client
}

func MDbDChurnBlockerMlResultDalConstructor() *mDbDChurnBlockerMlResultDal {
	return &mDbDChurnBlockerMlResultDal{Client: mongodb.GetMongodbClient()}
}

func (m *mDbDChurnBlockerMlResultDal) Add(data *model.ChurnBlockerMlResultModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection("churnBlockerMlResults")
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

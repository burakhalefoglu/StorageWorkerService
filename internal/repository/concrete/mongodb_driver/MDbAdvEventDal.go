package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbAdvEventDal struct {
	Client *mongo.Client
	Table  string
}

func NewMDbAdvEventDal(Table string) *mDbAdvEventDal {
	return &mDbAdvEventDal{Client: mongodb.GetMongodbClient(),
		Table: Table}
}

func (m *mDbAdvEventDal) Add(data *model.AdvEventDataModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection(m.Table)
	var _, err = collection.InsertOne(ctx, bson.D{
		{"Id", data.Id},
		{"Status", data.Status},
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

package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbBuyingEventDal struct {
	Client *mongo.Client
	Table  string
}

func NewMDbBuyingEventDal(Table string) *mDbBuyingEventDal {
	return &mDbBuyingEventDal{Client: mongodb.GetMongodbClient(),
		Table: Table}
}

func (m *mDbBuyingEventDal) Add(data *model.BuyingEventModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection(m.Table)
	var _, err = collection.InsertOne(ctx, bson.D{
		{"Id", data.Id},
		{"Status", data.Status},
		{"AdvType", data.ProductType},
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"LevelName", data.LevelName},
		{"InWhatMinutes", data.InMinutes},
		{"TriggeredTime", data.TriggeredTime},
	})
	if err != nil {
		return err
	}
	return nil
}

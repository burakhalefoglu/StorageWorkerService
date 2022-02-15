package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbItemDal struct {
	Client *mongo.Client
	Table  string
}

func NewMDbItemDal(Table string) *mDbItemDal {
	return &mDbItemDal{Client: mongodb.GetMongodbClient(),
		Table: Table}
}

func (m *mDbItemDal) Add(data *model.ItemModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection(m.Table)
	var _, err = collection.InsertOne(ctx, bson.D{
		{"Id", data.Id},
		{"Count", data.Count},
		{"InventoryId", data.InventoryId},
		{"ItemType", data.ItemType},
	})
	if err != nil {
		return err
	}
	return nil
}

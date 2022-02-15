package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbInventoryDal struct {
	Client *mongo.Client
	Table  string
}

func NewMDbInventoryDal(Table string) *mDbInventoryDal {
	return &mDbInventoryDal{Client: mongodb.GetMongodbClient(),
		Table: Table}
}

func (m *mDbInventoryDal) Add(data *model.InventoryModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection(m.Table)
	var _, err = collection.InsertOne(ctx, bson.D{
		{"Id", data.Id},
		{"Status", data.Status},
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"CustomerId", data.MinorMine},
		{"CustomerId", data.ModerateMine},
		{"CustomerId", data.PreciousMine},
		{"CustomerId", data.CreatedAt},
	})
	if err != nil {
		return err
	}
	return nil
}

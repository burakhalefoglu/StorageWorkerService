package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbDInventoryDal struct {
	Client *mongo.Client
}

func MDbDInventoryDalConstructor() *mDbDInventoryDal {
	return &mDbDInventoryDal{Client: mongodb.GetMongodbClient()}
}

func (m *mDbDInventoryDal) Add(data *model.InventoryModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection("inventories")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"CustomerId", data.MinorMine},
		{"CustomerId", data.ModerateMine},
		{"CustomerId", data.PreciousMine},
		{"CustomerId", data.Items},
		{"CustomerId", data.Skills},
		{"CustomerId", data.TemporaryAbilities},
		{"CustomerId", data.CreatedAt},
	})
	if err != nil {
		return err
	}
	return nil
}

package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MDbDInventoryDal struct {
	Client *mongo.Client
}

func (m *MDbDInventoryDal) Add(data *model.InventoryModel) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("Client").Collection("Inventory")
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
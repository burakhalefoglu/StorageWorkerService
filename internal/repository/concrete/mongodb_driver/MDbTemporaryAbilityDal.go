package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbTemporaryAbilityDal struct {
	Client *mongo.Client
	Table  string
}

func NewMDbTemporaryAbilityDal(Table string) *mDbTemporaryAbilityDal {
	return &mDbTemporaryAbilityDal{Client: mongodb.GetMongodbClient(),
		Table: Table}
}

func (m *mDbTemporaryAbilityDal) Add(data *model.TemporaryAbilityModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection(m.Table)
	var _, err = collection.InsertOne(ctx, bson.D{
		{"Id", data.Id},
		{"Count", data.Count},
		{"InventoryId", data.InventoryId},
		{"AbilityType", data.AbilityType},
	})
	if err != nil {
		return err
	}
	return nil
}

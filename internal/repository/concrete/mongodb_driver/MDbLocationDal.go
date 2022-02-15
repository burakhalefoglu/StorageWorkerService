package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbLocationDal struct {
	Client *mongo.Client
	Table  string
}

func NewMDbLocationDal(Table string) *mDbLocationDal {
	return &mDbLocationDal{Client: mongodb.GetMongodbClient(),
		Table: Table}
}

func (m *mDbLocationDal) Add(data *model.LocationModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection(m.Table)
	var _, err = collection.InsertOne(ctx, bson.D{
		{"Id", data.Id},
		{"Status", data.Status},
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"Continent", data.Continent},
		{"Country", data.Country},
		{"City", data.City},
		{"Query", data.Query},
		{"Region", data.Region},
		{"Org", data.Org},
		{"CreatedAt", data.CreatedAt},
	})
	if err != nil {
		return err
	}
	return nil
}

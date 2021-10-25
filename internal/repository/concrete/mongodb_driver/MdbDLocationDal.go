package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MDbDLocationDal struct {
	Client *mongo.Client
}

func (m *MDbDLocationDal) Add(data *model.LocationModel) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("Client").Collection("Location")
	var _, err = collection.InsertOne(ctx, bson.D{
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
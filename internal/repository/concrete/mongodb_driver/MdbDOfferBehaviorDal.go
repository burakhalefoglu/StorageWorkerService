package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MDbDOfferBehaviorDal struct {
	Client *mongo.Client
}

func (m *MDbDOfferBehaviorDal) Add(data *model.OfferBehaviorModel) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("Client").Collection("OfferBehavior")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"DateTime", data.DateTime},
		{"OfferId", data.OfferId},
		{"IsBuyOffer", data.IsBuyOffer},

	})
	if err != nil {
		return err
	}
	return nil
}
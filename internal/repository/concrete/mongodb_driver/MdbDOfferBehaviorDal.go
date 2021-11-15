package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbDOfferBehaviorDal struct {
	Client *mongo.Client
}

func MDbDOfferBehaviorDalConstructor() *mDbDOfferBehaviorDal {
	return &mDbDOfferBehaviorDal{Client: mongodb.GetMongodbClient()}
}

func (m *mDbDOfferBehaviorDal) Add(data *model.OfferBehaviorModel) error{

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

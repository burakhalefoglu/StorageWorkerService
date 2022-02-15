package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbOfferBehaviorDal struct {
	Client *mongo.Client
	Table  string
}

func NewMDbOfferBehaviorDal(Table string) *mDbOfferBehaviorDal {
	return &mDbOfferBehaviorDal{Client: mongodb.GetMongodbClient(),
		Table: Table}
}

func (m *mDbOfferBehaviorDal) Add(data *model.OfferBehaviorModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection(m.Table)
	var _, err = collection.InsertOne(ctx, bson.D{
		{"Id", data.Id},
		{"Status", data.Status},
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

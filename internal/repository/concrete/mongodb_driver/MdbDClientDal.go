package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MDbDClientDal struct {
	Client *mongo.Client
}

func (m *MDbDClientDal) Add(data *model.ClientDataModel) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("Client").Collection("ClientModel")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"IsPaidClient", data.IsPaidClient},
		{"PaidTime", data.PaidTime},
		{"CreatedAt", data.CreatedAt},

	})
	if err != nil {
		return err
	}
	return nil
}

func (m *MDbDClientDal) UpdateByClientId(clientId string, data *model.ClientDataModel) error {
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.D{{"$set", bson.D{
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"IsPaidClient", data.IsPaidClient},
		{"PaidTime", data.PaidTime},
		{"CreatedAt", data.CreatedAt},
		{"UpdatedAt", time.Now()},
	}}}

	collection := m.Client.Database("Client").Collection("ClientModel")
	updateResult := collection.FindOneAndUpdate(ctx, bson.D{{
		"ClientId",clientId,
	}}, update)
	if updateResult.Err() != nil{
		return updateResult.Err()
	}
	return nil
}

func (m *MDbDClientDal) GetByClientId(clientId string)(*model.ClientDataModel, error){

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("Client").Collection("ClientModel")
	var result = collection.FindOne(ctx, bson.D{{
		"ClientId",clientId,
	}})

	var model = model.ClientDataModel{}
	if result.Err() != nil {
		return &model, result.Err()
	}
	var err = result.Decode(&model)
	if err != nil{
		return &model, err
	}
	return &model, nil
}

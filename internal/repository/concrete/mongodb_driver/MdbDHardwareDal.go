package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbDHardwareDal struct {
	Client *mongo.Client
}

func MDbDHardwareDalConstructor() *mDbDHardwareDal {
	return &mDbDHardwareDal{Client: mongodb.GetMongodbClient()}
}

func (m *mDbDHardwareDal) Add(data *model.HardwareModel) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("Client").Collection("Hardware")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"DateTime", data.DateTime},
		{"DeviceModel", data.DeviceModel},
		{"DeviceName", data.DeviceName},
		{"DeviceType", data.DeviceType},
		{"GraphicsDeviceName", data.GraphicsDeviceName},
		{"GraphicsDeviceType", data.GraphicsDeviceType},
		{"GraphicsDeviceVendor", data.GraphicsDeviceVendor},
		{"GraphicsDeviceVersion", data.GraphicsDeviceVersion},
		{"GraphicsMemorySize", data.GraphicsMemorySize},
		{"OperatingSystem", data.OperatingSystem},
		{"ProcessorCount", data.ProcessorCount},
		{"SystemMemorySize", data.SystemMemorySize},
		{"ProcessorType", data.ProcessorType},


	})
	if err != nil {
		return err
	}
	return nil
}

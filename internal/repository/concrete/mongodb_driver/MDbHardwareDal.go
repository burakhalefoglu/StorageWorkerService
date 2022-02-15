package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbHardwareDal struct {
	Client *mongo.Client
	Table  string
}

func NewMDbHardwareDal(Table string) *mDbHardwareDal {
	return &mDbHardwareDal{Client: mongodb.GetMongodbClient(),
		Table: Table}
}

func (m *mDbHardwareDal) Add(data *model.HardwareModel) error {

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

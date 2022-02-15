package mongodb_driver

import (
	"StorageWorkerService/internal/model"
	"StorageWorkerService/pkg/database/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mDbGameSessionDal struct {
	Client *mongo.Client
	Table  string
}

func NewMDbGameSessionDal(Table string) *mDbGameSessionDal {
	return &mDbGameSessionDal{Client: mongodb.GetMongodbClient(),
		Table: Table}
}

func (m *mDbGameSessionDal) Add(data *model.GameSessionModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("ClientDatabase").Collection(m.Table)
	var _, err = collection.InsertOne(ctx, bson.D{
		{"Id", data.Id},
		{"Status", data.Status},
		{"ProjectId", data.ProjectId},
		{"ClientId", data.ClientId},
		{"CustomerId", data.CustomerId},
		{"SessionTimeMinute", data.SessionTimeMinute},
		{"SessionStartTime", data.SessionStartTime},
		{"SessionFinishTime", data.SessionFinishTime},
	})
	if err != nil {
		return err
	}
	return nil
}

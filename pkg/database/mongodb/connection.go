package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func GetMongodbClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var url = os.Getenv("MONGODB_CONN")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil{
		//Log.SendPanicLog("MongoConnection", "ConnectMongodb", err)
		panic(err)
	}
	return client
}
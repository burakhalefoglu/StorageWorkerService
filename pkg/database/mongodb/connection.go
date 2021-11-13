package mongodb

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type Connection struct {
	Conn *mongo.Client
}
var Conn = Connection{
	Conn: ConnectMongodb(),
}

func ConnectMongodb() *mongo.Client {
	godotenv.Load()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var url = os.Getenv("MONGODB_CONN")
	log.Print(url)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil{
		panic(err)
	}
	return client
}
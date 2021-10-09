package endpoints

// contains MONGODB based functions

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client           *mongo.Client
	CONNECTIONURL    = "mongodb://localhost:27017"
	DATABASE         = "appointy"
	COLLECTION_USERS = "users"
	COLLECTION_POSTS = "posts"
)

func GetUserClient(Client *mongo.Client) *mongo.Collection {
	connection := Client.Database(DATABASE).Collection(COLLECTION_USERS)
	return connection
}

func GetPostsClient(Client *mongo.Client) *mongo.Collection {
	connection := Client.Database(DATABASE).Collection(COLLECTION_POSTS)
	return connection
}

func InitMongo() {
	Client, _ = mongo.NewClient(options.Client().ApplyURI(CONNECTIONURL))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb Connection Initialized...")
}

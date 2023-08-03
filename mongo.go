package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnectDB() *mongo.Client {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	fmt.Println(os.Getenv("DB_SERVER") + "://" + os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST"))
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(os.Getenv("DB_SERVER") + "://" + os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + "/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client

}

func MgoCollection(coll string, client *mongo.Client) *mongo.Collection {
	return client.Database("products").Collection(coll)
}

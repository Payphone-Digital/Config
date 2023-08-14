package config

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client          *mongo.Client
	collectionCache = map[string]*mongo.Collection{}
)

func MongoConnectDB(db string, coll string) (*mongo.Collection, *mongo.Client, error) {
	if client != nil {
		return GetCollection(db, coll), client, nil
	}
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("DB_SERVER") + "://" + os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + "/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)

	// Mengecek apakah koneksi sukses
	if client.Ping(ctx, nil) != nil {
		client.Disconnect(context.Background())
		client = nil
		return nil, client, err
	}

	return GetCollection(db, coll), client, err
}

func GetCollection(db string, collName string) *mongo.Collection {

	// Mengecek apakah koleksi sudah ada di cache, jika tidak, membuatnya
	if coll, ok := collectionCache[collName]; ok {
		return coll
	}

	coll := client.Database(db).Collection(collName)
	collectionCache[collName] = coll
	return coll
}

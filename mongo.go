package config

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client          *mongo.Client
	collectionCache = map[string]*mongo.Collection{}
	cacheMutex      sync.Mutex
	clientMutex     sync.Mutex
)

func MongoConnectDB(db, coll string) (*mongo.Collection, error) {
	clientMutex.Lock()
	defer clientMutex.Unlock()

	// Mencoba mengambil koleksi dari cache
	if coll, ok := collectionCache[coll]; ok {
		return coll, nil
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(os.Getenv("DB_SERVER") + "://" + os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + "/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPI).SetMinPoolSize(5).
		SetMaxPoolSize(100) // Atur ukuran maksimum dari connection pool

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if client == nil {
		var err error
		client, err = mongo.Connect(ctx, opts)
		if err != nil {
			return nil, err
		}

		// Menambahkan goroutine untuk menangani sinyal shutdown
		go func() {
			sigint := make(chan os.Signal, 1)
			signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)

			<-sigint
			// Menerima sinyal shutdown, menutup koneksi dengan aman
			if client != nil {
				client.Disconnect(context.Background())
			}
			os.Exit(0)
		}()
	}

	err := client.Ping(ctx, nil)
	if err != nil {
		client.Disconnect(context.Background())
		client = nil
		return nil, err
	}

	return GetCollection(db, coll), nil
}

func GetCollection(db, collName string) *mongo.Collection {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if coll, ok := collectionCache[collName]; ok {
		return coll
	}

	coll := client.Database(db).Collection(collName)
	collectionCache[collName] = coll
	return coll
}

func DbPointing(url, ds, cn string) (*mongo.Collection, string, error) {
	idDb := strings.Split(url, "/")
	colls := strings.Title(idDb[3]) + cn
	db, err := MongoConnectDB(strings.Title(idDb[2])+ds, colls)
	return db, colls, err
}

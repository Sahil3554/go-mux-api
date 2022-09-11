package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Client instance
var DB *mongo.Client

func ConnectDB() {
	// GetEnv()
	URI := os.Getenv("MONGO_URI")
	if URI == "" {
		GetEnv()
		URI = os.Getenv("MONGO_URI")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	DB = client
}

//getting database collections
func GetCollection(collectionName string) *mongo.Collection {
	if DB == nil {
		ConnectDB()
	}
	collection := DB.Database(os.Getenv("DB_NAME")).Collection(collectionName)
	return collection
}

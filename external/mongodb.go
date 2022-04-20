package external

import (
	"context"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoConn(document string) *mongo.Collection {
	var clientOptions *options.ClientOptions
	var client *mongo.Client
	var collection *mongo.Collection
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if os.Getenv("ENVIRONMENT") != "local" {
		serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
		clientOptions = options.Client().
			ApplyURI(fmt.Sprintf("mongodb+srv://ariebrainware:%s@cluster0.h2eai.mongodb.net/%s?retryWrites=true&w=majority", os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_DATABASE"))).
			SetServerAPIOptions(serverAPIOptions)
		client, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Error(err)
			panic("Failed to connect mongo")
		}
	}

	collection = client.Database("rama-shortner").Collection(document)
	err = client.Connect(ctx)
	if err != nil {
		log.Error(err)
		panic("Failed to connect mongo")
	}
	return collection
}

package mongo

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func init_viper() {
	viper.SetConfigFile(`src/configurations/mongo.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func GetDbName() string {
	init_viper()
	return viper.GetString(`database`)
}

func NewMongoClient() (*mongo.Client, *context.Context) {
	init_viper()

	mongodb_uri := viper.GetString(`mongodb_uri`)
	clientOptions := options.Client().ApplyURI(mongodb_uri)
	client, err := mongo.NewClient(clientOptions)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("users MongoDB!")

	return client, &ctx
}
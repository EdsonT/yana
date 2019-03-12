package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	cli *mongo.Client
	err error
	
)

func GetClient() *mongo.Client {
	// cli, err = mongo.NewClient(options.Client().ApplyURI("mongo://10.31.32.66:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cli, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://10.31.32.66:27017"))
	if err != nil {
		fmt.Println(err)
	}
	return cli
}

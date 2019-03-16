package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	opt  *options.ClientOptions
	cli  *mongo.Client
	auth options.Credential
	err  error
)
// NewMongoClient creates a mongo client connection with the necessary parameters
func NewMongoClient() *mongo.Client {
	opt = &options.ClientOptions{}
	auth = options.Credential{
		AuthSource: "yanaDb",
		Username:   "admin",
		Password:   "Setitiruces0",
	}
	// ec2-mongo
	opt.SetAuth(auth)
	opt.ApplyURI("mongodb://ec2-3-18-8-186.us-east-2.compute.amazonaws.com:27017")

	
	// localConfig
	// opt.ApplyURI("mongodb://10.31.32.66:27017")

	

	cli, err = mongo.NewClient(opt)

	err = cli.Connect(context.Background())

	return cli
}

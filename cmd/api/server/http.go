package server

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New() *gin.Engine {
	// New http server
	server := gin.New()

	// Create db connection
	dbcli := dbConnect()

	MapUrls(server, dbcli)

	return server
}

func dbConnect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	// Check connections
	if err := client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}

	return client
}

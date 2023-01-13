package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go/crud/controllers"
	"github.com/go/crud/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


var (
	server      *gin.Engine
	us          services.MovieService
	uc          controllers.MovieController
	ctx         context.Context
	moviec       *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	moviec = mongoclient.Database("moviedb").Collection("movies")
	us = services.NewMovieService(moviec, ctx)
	uc = controllers.New(us)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	uc.RegisterMovieRoutes(basepath)

	log.Fatal(server.Run(":9090"))

}

package common

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient mongo.Client
var StockMogoCollection *mongo.Collection

func init() {

	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接数据库
	MongoClient, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}

	// 判断服务是不是可用
	if err = MongoClient.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	StockMogoCollection = MongoClient.Database("mongotest").Collection("stock")
}

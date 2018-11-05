package main

import (
	"log"

	"github.com/mlvhub/learning-go/mongodb-rest-api/pkg"
	"github.com/mlvhub/learning-go/mongodb-rest-api/pkg/crypto"
	"github.com/mlvhub/learning-go/mongodb-rest-api/pkg/mongo"
	"github.com/mlvhub/learning-go/mongodb-rest-api/pkg/server"
)

func main() {
	mongoConfig := &root.MongoConfig{
		URL:       "mongodb://test:test@testcluster-shard-00-00-viokg.mongodb.net:27017,testcluster-shard-00-01-viokg.mongodb.net:27017,testcluster-shard-00-02-viokg.mongodb.net:27017/test?ssl=true&replicaSet=TestCluster-shard-0&authSource=admin",
		DBName:    "go_mongo_tut",
		UserTable: "users",
	}

	ms, err := mongo.NewSession(mongoConfig.URL)
	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}
	defer ms.Close()

	h := crypto.Crypto{}
	u := mongo.NewUserService(ms.Copy(), mongoConfig, &h)
	s := server.NewServer(u)

	s.Start()
}

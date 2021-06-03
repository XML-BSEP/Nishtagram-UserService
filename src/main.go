package main

import (
	"github.com/gin-gonic/gin"
	"user-service/infrastructure/mongo"
	"user-service/infrastructure/seeder"
)

func main() {

	mongoCli, ctx := mongo.NewMongoClient()
	db := mongo.GetDbName()
	println(db)
	seeder.SeedData(db, mongoCli, ctx)
	g := gin.Default()
	g.Run("localhost:8082")
}

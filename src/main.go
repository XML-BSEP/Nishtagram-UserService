package main

import (
	router2 "user-service/http/router"
	"user-service/infrastructure/mongo"
	"user-service/infrastructure/seeder"
	interactor2 "user-service/interactor"
)

func main() {

	mongoCli, ctx := mongo.NewMongoClient()
	db := mongo.GetDbName()
	seeder.SeedData(db, mongoCli, ctx)

	interactor := interactor2.NewInteractor(mongoCli)
	appHandler := interactor.NewAppHandler()

	router := router2.NewRouter(appHandler)
	router.Run("localhost:8082")
}

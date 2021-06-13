package main

import (
	"context"
	logger "github.com/jelena-vlajkov/logger/logger"
	router2 "user-service/http/router"
	"user-service/infrastructure/mongo"
	"user-service/infrastructure/seeder"
	interactor2 "user-service/interactor"

)

func main() {

	logger := logger.InitializeLogger("user-service", context.Background())

	mongoCli, ctx := mongo.NewMongoClient()
	db := mongo.GetDbName()
	seeder.SeedData(db, mongoCli, ctx)

	interactor := interactor2.NewInteractor(mongoCli, logger)
	appHandler := interactor.NewAppHandler()


	router := router2.NewRouter(appHandler, logger)


	router.RunTLS(":8082", "src/certificate/cert.pem", "src/certificate/key.pem")
}

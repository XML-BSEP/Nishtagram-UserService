package main

import (
	"context"
	logger "github.com/jelena-vlajkov/logger/logger"
	"os"
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

	if os.Getenv("DOCKER_ENV") == "" {

		err := router.RunTLS(":8082", "certificate/cert.pem", "certificate/key.pem")
		if err != nil {
			return
		}
	} else {
		err := router.Run(":8082")
		if err != nil {
			return 
		}
	}
}

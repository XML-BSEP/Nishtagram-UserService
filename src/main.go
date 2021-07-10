package main

import (
	"context"
	logger "github.com/jelena-vlajkov/logger/logger"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
	router2 "user-service/http/router"
	"user-service/infrastructure/grpc/service/user_service"
	"user-service/infrastructure/mongo"
	"user-service/infrastructure/saga"
	"user-service/infrastructure/saga_redisdb"
	"user-service/infrastructure/seeder"
	interactor2 "user-service/interactor"
)

func getNetListener(port uint) net.Listener {
	var domain string
	if os.Getenv("DOCKER_ENV") == "" {
		domain = "127.0.0.1"
	} else {
		domain = "userms"
	}
	domain = domain + ":" + strconv.Itoa(int(port))
	lis, err := net.Listen("tcp", domain)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return lis
}

func main() {

	logger := logger.InitializeLogger("user-service", context.Background())

	mongoCli, ctx := mongo.NewMongoClient()
	db := mongo.GetDbName()
	seeder.SeedData(db, mongoCli, ctx)
	sagaRedisClient := saga_redisdb.NewSagaRedis(logger)

	interactor := interactor2.NewInteractor(mongoCli, logger)
	appHandler := interactor.NewAppHandler()

	authSaga := saga.NewAuthSaga(interactor.NewProfileInfoUseCase(), sagaRedisClient)
	go authSaga.SagaAuth(context.Background())

	router := router2.NewRouter(appHandler, logger)

	lis := getNetListener(8075)
	userServiceImpl := interactor.NewUserServiceImpl()
	grpcServer := grpc.NewServer()
	user_service.RegisterUserDetailsServer(grpcServer, userServiceImpl)
	go func() {
		log.Fatalln(grpcServer.Serve(lis))
	}()

	if os.Getenv("DOCKER_ENV") == "" {

		err := router.RunTLS(":8082", "src/certificate/cert.pem", "src/certificate/key.pem")
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

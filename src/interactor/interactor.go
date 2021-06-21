package interactor

import (
	logger "github.com/jelena-vlajkov/logger/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"user-service/http/handler"
	"user-service/infrastructure/grpc/service/user_service/implementation"
	"user-service/repository"
	"user-service/usecase"
)

type interactor struct {
	db *mongo.Client
	logger *logger.Logger
}

type Interactor interface {

	NewProfileInfoRepository() repository.ProfileInfoRepository
	NewRequestVerificationRepository() repository.RequestVerificationRepository

	NewProfileInfoUseCase() usecase.ProfileInfoUseCase
	NewRequestVerificationUseCase() usecase.RequestVerificationUseCase

	NewProfileInfoHandler() handler.ProfileInfoHandler
	NewRequestVerificationHandler() handler.RequestVerificationHandler

	NewAppHandler() AppHandler

	NewUserServiceImpl() *implementation.UserServiceImpl
}


func NewInteractor(db *mongo.Client, logger *logger.Logger) Interactor {
	return &interactor{db: db, logger: logger}
}

func (i *interactor) NewProfileInfoRepository() repository.ProfileInfoRepository {
	return repository.NewProfileInfoRepository(i.db, i.logger)
}

func (i *interactor) NewRequestVerificationRepository() repository.RequestVerificationRepository {
	return repository.NewRequestVerificationRepository(i.db)
}

func (i *interactor) NewProfileInfoUseCase() usecase.ProfileInfoUseCase {
	return usecase.NewProfileInfoUseCase(i.NewProfileInfoRepository(), i.logger)
}

func (i *interactor) NewRequestVerificationUseCase() usecase.RequestVerificationUseCase {
	return  usecase.NewRequestVerificationUseCase(i.NewRequestVerificationRepository(), i.NewProfileInfoRepository())
}

func (i *interactor) NewProfileInfoHandler() handler.ProfileInfoHandler {
	return handler.NewProfileInfoHandler(i.NewProfileInfoUseCase(), i.logger)
}
func (i *interactor) NewRequestVerificationHandler() handler.RequestVerificationHandler {
	return handler.NewRequestVerificationHandler(i.NewRequestVerificationUseCase())
}

func (i *interactor) NewUserServiceImpl() *implementation.UserServiceImpl {
	return implementation.NewUserServiceImpl(i.NewProfileInfoUseCase())
}

type appHandler struct {
	handler.ProfileInfoHandler
	handler.RequestVerificationHandler

}

type AppHandler interface {
	handler.ProfileInfoHandler
	handler.RequestVerificationHandler

}

func (i *interactor) NewAppHandler() AppHandler{
	appHandler := &appHandler{}
	appHandler.ProfileInfoHandler = i.NewProfileInfoHandler()
	appHandler.RequestVerificationHandler = i.NewRequestVerificationHandler()


	return appHandler
}
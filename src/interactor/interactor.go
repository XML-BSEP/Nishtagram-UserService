package interactor

import (
	"go.mongodb.org/mongo-driver/mongo"
	"user-service/http/handler"
	"user-service/repository"
	"user-service/usecase"
)

type interactor struct {
	db *mongo.Client
}


type Interactor interface {
	NewProfileRepository() repository.ProfileRepository
	NewProfileInfoRepository() repository.ProfileInfoRepository

	NewProfileUseCase() usecase.ProfileUseCase
	NewProfileInfoUseCase() usecase.ProfileInfoUseCase


	NewProfileInfoHandler() handler.ProfileInfoHandler

	NewAppHandler() AppHandler
}


func NewInteractor(db *mongo.Client) Interactor {
	return &interactor{db: db}
}

func (i *interactor) NewProfileRepository() repository.ProfileRepository {
	return repository.NewProfileRepository(i.db)
}

func (i *interactor) NewProfileUseCase() usecase.ProfileUseCase {
	return usecase.NewProfileUseCase(i.NewProfileRepository())
}


func (i *interactor) NewProfileInfoRepository() repository.ProfileInfoRepository {
	return repository.NewProfileInfoRepository(i.db)
}

func (i *interactor) NewProfileInfoUseCase() usecase.ProfileInfoUseCase {
	return usecase.NewProfileInfoUseCase(i.NewProfileInfoRepository())
}

func (i *interactor) NewProfileInfoHandler() handler.ProfileInfoHandler {
	return handler.NewProfileInfoHandler(i.NewProfileInfoUseCase(), i.NewProfileUseCase())
}

type appHandler struct {
	handler.ProfileInfoHandler

}

type AppHandler interface {
	handler.ProfileInfoHandler

}

func (i *interactor) NewAppHandler() AppHandler{
	appHandler := &appHandler{}
	appHandler.ProfileInfoHandler = i.NewProfileInfoHandler()

	return appHandler
}
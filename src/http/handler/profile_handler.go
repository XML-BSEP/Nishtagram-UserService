package handler

import (
	"github.com/gin-gonic/gin"
	"user-service/usecase"
)

type profileHandler struct {
	ProfileUseCase usecase.ProfileUseCase
}

func (p profileHandler) GetProfileByUsername(ctx *gin.Context) {
	panic("implement me")
}

func (p profileHandler) GetAllProfiles(ctx *gin.Context) {
	panic("implement me")
}

func (p profileHandler) GetAllUserProfiles(ctx *gin.Context) {
	panic("implement me")
}

type ProfileHandler interface {
	GetProfileByUsername(ctx *gin.Context)
	GetAllProfiles(ctx *gin.Context)
	GetAllUserProfiles(ctx *gin.Context)

}

func NewProfileHandler(usecase usecase.ProfileUseCase) ProfileHandler{
	return &profileHandler{ProfileUseCase: usecase}
}
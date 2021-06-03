package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/usecase"
)

type profileInfoHandlder struct {
	ProfileInfoUseCase usecase.ProfileInfoUseCase
}

func (p *profileInfoHandlder) GetProfileInfoByUsername(ctx *gin.Context) {
	username := struct {
		Username string
	}{}

	decoder := json.NewDecoder(ctx.Request.Body)
	dec_err := decoder.Decode(&username)

	if dec_err != nil {
		ctx.JSON(http.StatusBadRequest, "username decoding error")
		ctx.Abort()
		return
	}

	profileInfo, err := p.ProfileInfoUseCase.GetByUsername(username.Username, ctx)

	if err != nil {
		ctx.JSON(http.StatusNotFound, "No users with that username")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": profileInfo})

}


func (p *profileInfoHandlder) GetById(ctx *gin.Context) {
	id := struct {
		Id string
	}{}

	decoder := json.NewDecoder(ctx.Request.Body)
	dec_err := decoder.Decode(&id)

	if dec_err != nil {
		ctx.JSON(http.StatusBadRequest, "username decoding error")
		ctx.Abort()
		return
	}

	profileInfo, err := p.ProfileInfoUseCase.GetById(id.Id, ctx)

	if err != nil {
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": profileInfo})
}

type ProfileInfoHandler interface {
	GetProfileInfoByUsername(ctx *gin.Context)

	GetById(ctx *gin.Context)

}

func NewProfileInfoHandler(usecase usecase.ProfileInfoUseCase) ProfileInfoHandler{
	return &profileInfoHandlder{ProfileInfoUseCase: usecase}
}

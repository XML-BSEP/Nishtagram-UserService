package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/dto"
	"user-service/usecase"
)

type profileInfoHandlder struct {
	ProfileInfoUseCase usecase.ProfileInfoUseCase
	ProfileUseCase usecase.ProfileUseCase
}

func (p *profileInfoHandlder) GetProfileUsernameImageById(ctx *gin.Context) {
	var id string
	err := json.NewDecoder(ctx.Request.Body).Decode(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Decoding error")
		ctx.Abort()
		return
	}

	profile, err1 := p.ProfileInfoUseCase.GetById(id, ctx)
	if err1 != nil {
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	var profileUIDTO dto.ProfileUsernameImageDTO
	profileUIDTO = dto.NewProfileUsernameImage(profile.Profile.Username, profile.ProfileImage)

	ctx.JSON(http.StatusOK, gin.H{"data": profileUIDTO})


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

func (p *profileInfoHandlder) IsPrivate(ctx *gin.Context) {
	var id string
	err := json.NewDecoder(ctx.Request.Body).Decode(&id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Decoding error")
		ctx.Abort()
		return
	}

	profile, err1 := p.ProfileInfoUseCase.GetById(id, ctx)

	if err1 != nil {
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	isPrivate, err2 := p.ProfileUseCase.IsProfilePrivate(profile.Profile.Username, ctx)
	if err2 != nil {
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": isPrivate})

}


type ProfileInfoHandler interface {
	GetProfileInfoByUsername(ctx *gin.Context)
	GetById(ctx *gin.Context)
	IsPrivate(ctx *gin.Context)
	GetProfileUsernameImageById(ctx *gin.Context)

}


func NewProfileInfoHandler(usecase usecase.ProfileInfoUseCase, profileUsecase usecase.ProfileUseCase) ProfileInfoHandler{
	return &profileInfoHandlder{ProfileInfoUseCase: usecase, ProfileUseCase: profileUsecase}
}

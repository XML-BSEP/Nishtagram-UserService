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
}



func (p *profileInfoHandlder) GetProfileUsernameImageById(ctx *gin.Context) {
	id := ctx.Request.URL.Query().Get("userId")

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
	username := ctx.Request.URL.Query().Get("username")

	profileInfo, err := p.ProfileInfoUseCase.GetByUsername(username, ctx)

	if err != nil {
		ctx.JSON(http.StatusNotFound, "No users with that username")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": profileInfo})

}

func (p *profileInfoHandlder) GetById(ctx *gin.Context) {
	id := ctx.Request.URL.Query().Get("userId")

	profileInfo, err := p.ProfileInfoUseCase.GetById(id, ctx)

	if err != nil {
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": profileInfo})
}

func (p *profileInfoHandlder) IsPrivate(ctx *gin.Context) {
	id := ctx.Request.URL.Query().Get("userId")

	profile, err1 := p.ProfileInfoUseCase.GetById(id, ctx)

	if err1 != nil {
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	isPrivate, err2 := p.ProfileInfoUseCase.IsProfilePrivate(profile.Profile.Username, ctx)
	if err2 != nil {
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": isPrivate})

}

func (p *profileInfoHandlder) GetUserById(ctx *gin.Context) {
	id := ctx.Request.URL.Query().Get("userId")

	profileDTO, error := p.ProfileInfoUseCase.GetUserById(id, ctx)
	if error != nil{
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": profileDTO})
}

func (p *profileInfoHandlder) GetUserProfileById(ctx *gin.Context) {

	id := ctx.Request.URL.Query().Get("userId")


	profileUserDTO, error := p.ProfileInfoUseCase.GetUserProfileById(id, ctx)
	if error != nil{
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": profileUserDTO})


}

type ProfileInfoHandler interface {
	GetProfileInfoByUsername(ctx *gin.Context)
	GetById(ctx *gin.Context)
	IsPrivate(ctx *gin.Context)
	GetProfileUsernameImageById(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	GetUserProfileById(ctx *gin.Context)
	SaveNewUser(ctx *gin.Context)
}

func (p *profileInfoHandlder) SaveNewUser(ctx *gin.Context) {
	var newUserDTO dto.NewUserDTO

	err := json.NewDecoder(ctx.Request.Body).Decode(&newUserDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Decoding error")
		ctx.Abort()
		return
	}

	exists, _ := p.ProfileInfoUseCase.Exists(newUserDTO.Username, newUserDTO.Email, ctx)

	if exists {
		ctx.JSON(400, gin.H{"message" : "User already exists"})
		return
	}

	newUserProfile := dto.NewUserDTOtoEntity(newUserDTO)

	error := p.ProfileInfoUseCase.SaveNewUser(newUserProfile, ctx)
	if error != nil {
		ctx.JSON(http.StatusNotFound, "Failed to save")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, "idegasnamax")

}


func NewProfileInfoHandler(usecase usecase.ProfileInfoUseCase) ProfileInfoHandler{
	return &profileInfoHandlder{ProfileInfoUseCase: usecase}
}

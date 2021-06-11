package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"net/http"
	"strings"
	"user-service/dto"
	"user-service/infrastructure/mapper"
	validator2 "user-service/infrastructure/validator"
	"user-service/usecase"
)

type profileInfoHandlder struct {
	ProfileInfoUseCase usecase.ProfileInfoUseCase
}

func (p *profileInfoHandlder) SearchPublicUser(ctx *gin.Context) {
	search := ctx.Request.URL.Query().Get("search")
	/*policy := bluemonday.UGCPolicy()
		email := strings.TrimSpace(policy.Sanitize(req.Email))*/
	policy := bluemonday.UGCPolicy()
	search = strings.TrimSpace(policy.Sanitize(search))

	users, err := p.ProfileInfoUseCase.SearchPublicUsers(search, ctx)
	if err != nil {
		ctx.JSON(404, "User does not exists")
		return
	}

	var usersDTO []dto.UserDTO
	for _, user := range users {
		usersDTO = append(usersDTO, dto.NewUserDTOfromEntity(*user))
	}

	if len(usersDTO) == 0 {
		ctx.JSON(404, "User does not exists")
		return
	}

	ctx.JSON(200, usersDTO)

}

func (p *profileInfoHandlder) SearchUser(ctx *gin.Context) {
	search := ctx.Request.URL.Query().Get("search")

	policy := bluemonday.UGCPolicy()
	search = strings.TrimSpace(policy.Sanitize(search))

	users, err := p.ProfileInfoUseCase.SearchUser(search, ctx)
	if err != nil {
		ctx.JSON(404, "User does not exists")
		return
	}

	var usersDTO []dto.SearchUserDTO
	for _, user := range users {
		usersDTO = append(usersDTO, dto.NewSearchUserDTOFromEntity(*user))
	}

	if len(usersDTO) == 0 {
		ctx.JSON(404, "User does not exists")
		return
	}

	ctx.JSON(200, usersDTO)

}

func (p *profileInfoHandlder) IsPrivatePostMethod(ctx *gin.Context) {

	var privacyCheck dto.PrivacyCheckDto

	fmt.Println(ctx.Request.Body)

	if err := json.NewDecoder(ctx.Request.Body).Decode(&privacyCheck); err != nil {
		ctx.JSON(500, "Error decoding body")
		return
	}

	policy := bluemonday.UGCPolicy()
	privacyCheck.Id = strings.TrimSpace(policy.Sanitize(privacyCheck.Id))


	isPrivate, err := p.ProfileInfoUseCase.IsPrivateById(privacyCheck.Id, ctx)

	if err != nil {
		ctx.JSON(400, gin.H{"message" : "Can not get profile"})
		return
	}

	ctx.JSON(200, gin.H{"is_private" : isPrivate})

}

func (p *profileInfoHandlder) GetProfileInfoById(ctx *gin.Context) {

	id := ctx.Request.URL.Query().Get("userId")

	policy := bluemonday.UGCPolicy()
	id = strings.TrimSpace(policy.Sanitize(id))

	val, err := p.ProfileInfoUseCase.GetById(id, ctx)

	if err != nil {
		ctx.JSON(400, gin.H{"message" : "Can not get profile info"})
		return
	}

	if p.ProfileInfoUseCase.IsBanned(val, ctx) {
		ctx.JSON(400, gin.H{"message" : "User is baned"})
		return
	}

	profileInfoDto := mapper.ProfileInfoToProfileInfoDto(val)

	ctx.JSON(200, profileInfoDto)
}

func (p *profileInfoHandlder) GetProfileUsernameImageById(ctx *gin.Context) {
	id := ctx.Request.URL.Query().Get("userId")

	policy := bluemonday.UGCPolicy()
	id = strings.TrimSpace(policy.Sanitize(id))

	profile, err1 := p.ProfileInfoUseCase.GetById(id, ctx)
	if err1 != nil {
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	var profileUIDTO dto.ProfileUsernameImageDTO
	profileUIDTO = dto.NewProfileUsernameImage(profile.Profile.Username, profile.ProfileImage)

	ctx.JSON(http.StatusOK, profileUIDTO)


}

func (p *profileInfoHandlder) GetProfileInfoByUsername(ctx *gin.Context) {
	username := ctx.Request.URL.Query().Get("username")

	policy := bluemonday.UGCPolicy()
	username = strings.TrimSpace(policy.Sanitize(username))


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

	policy := bluemonday.UGCPolicy()
	id = strings.TrimSpace(policy.Sanitize(id))

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

	policy := bluemonday.UGCPolicy()
	id = strings.TrimSpace(policy.Sanitize(id))


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

	policy := bluemonday.UGCPolicy()
	id = strings.TrimSpace(policy.Sanitize(id))


	//DEKODE OVDEE
	profileDTO, error := p.ProfileInfoUseCase.GetUserById(id, ctx)
	if error != nil{
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, profileDTO)
}

func (p *profileInfoHandlder) GetUserProfileById(ctx *gin.Context) {

	id := ctx.Request.URL.Query().Get("userId")
	policy := bluemonday.UGCPolicy()
	id = strings.TrimSpace(policy.Sanitize(id))


	//DEKODE OVDEE
	profileUserDTO, error := p.ProfileInfoUseCase.GetUserProfileById(id, ctx)
	if error != nil{
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, profileUserDTO)


}

func (p *profileInfoHandlder) SaveNewUser(ctx *gin.Context) {
	var newUserDTO dto.NewUserDTO


	err := json.NewDecoder(ctx.Request.Body).Decode(&newUserDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Decoding error")
		ctx.Abort()
		return
	}

	//sanitizer
	policy := bluemonday.UGCPolicy()
	newUserDTO.ID = strings.TrimSpace(policy.Sanitize(newUserDTO.ID))
	newUserDTO.Name = strings.TrimSpace(policy.Sanitize(newUserDTO.Name))
	newUserDTO.Surname = strings.TrimSpace(policy.Sanitize(newUserDTO.Surname))
	newUserDTO.Email = strings.TrimSpace(policy.Sanitize(newUserDTO.Email))
	newUserDTO.Address = strings.TrimSpace(policy.Sanitize(newUserDTO.Address))
	newUserDTO.Phone = strings.TrimSpace(policy.Sanitize(newUserDTO.Phone))
	newUserDTO.Birthday = strings.TrimSpace(policy.Sanitize(newUserDTO.Birthday))
	newUserDTO.Gender = strings.TrimSpace(policy.Sanitize(newUserDTO.Gender))
	newUserDTO.Web = strings.TrimSpace(policy.Sanitize(newUserDTO.Web))
	newUserDTO.Bio = strings.TrimSpace(policy.Sanitize(newUserDTO.Bio))
	newUserDTO.Username = strings.TrimSpace(policy.Sanitize(newUserDTO.Username))
	newUserDTO.Image = strings.TrimSpace(policy.Sanitize(newUserDTO.Image))

	if newUserDTO.Birthday == "" {
		ctx.JSON(400, gin.H{"message" : "Enter birthday!"})
		return
	}

	if strings.Contains(newUserDTO.Username, " ") {
		ctx.JSON(400, gin.H{"message" : "Username is not in valid format!"})
		return
	}


	exists, _ := p.ProfileInfoUseCase.Exists(newUserDTO.Username, newUserDTO.Email, ctx)

	if exists {
		ctx.JSON(400, gin.H{"message" : "User already exists"})
		return
	}

	if newUserDTO.Image != "" {
		mediaToAttach, err := p.ProfileInfoUseCase.EncodeBase64(newUserDTO.Image, newUserDTO.ID, context.Background())
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "error while decoding base64")
			ctx.Abort()
			return
		}
		newUserDTO.Image = mediaToAttach
	}

	newUserProfile := dto.NewUserDTOtoEntity(newUserDTO)

	customValidator := validator2.NewCustomValidator()
	translator, _ := customValidator.RegisterEnTranslation()
	errValidation := customValidator.Validator.Struct(newUserProfile)
	errs := customValidator.TranslateError(errValidation, translator)
	errorsString := customValidator.GetErrorsString(errs)

	if errValidation != nil {
		ctx.JSON(400, gin.H{"message" : errorsString[0]})
		return
	}

	error := p.ProfileInfoUseCase.SaveNewUser(newUserProfile, ctx)
	if error != nil {
		ctx.JSON(http.StatusNotFound, "Failed to save")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, "idegasnamax")

}

func (p *profileInfoHandlder) GetAllPublicProfiles(ctx *gin.Context) {
	//DEKODE OVDEE
	users, err := p.ProfileInfoUseCase.GetAllPublicProfiles(ctx)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(200, gin.H{"data" : users})
	return
}

func (p *profileInfoHandlder) EditUser(ctx *gin.Context) {
	var newUserDTO dto.NewUserDTO

	err := json.NewDecoder(ctx.Request.Body).Decode(&newUserDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Decoding error")
		ctx.Abort()
		return
	}

	//sanitizer
	policy := bluemonday.UGCPolicy()
	newUserDTO.ID = strings.TrimSpace(policy.Sanitize(newUserDTO.ID))
	newUserDTO.Name = strings.TrimSpace(policy.Sanitize(newUserDTO.Name))
	newUserDTO.Surname = strings.TrimSpace(policy.Sanitize(newUserDTO.Surname))
	newUserDTO.Email = strings.TrimSpace(policy.Sanitize(newUserDTO.Email))
	newUserDTO.Address = strings.TrimSpace(policy.Sanitize(newUserDTO.Address))
	newUserDTO.Phone = strings.TrimSpace(policy.Sanitize(newUserDTO.Phone))
	newUserDTO.Birthday = strings.TrimSpace(policy.Sanitize(newUserDTO.Birthday))
	newUserDTO.Gender = strings.TrimSpace(policy.Sanitize(newUserDTO.Gender))
	newUserDTO.Web = strings.TrimSpace(policy.Sanitize(newUserDTO.Web))
	newUserDTO.Bio = strings.TrimSpace(policy.Sanitize(newUserDTO.Bio))
	newUserDTO.Username = strings.TrimSpace(policy.Sanitize(newUserDTO.Username))
	newUserDTO.Image = strings.TrimSpace(policy.Sanitize(newUserDTO.Image))

	if newUserDTO.Birthday == "" {
		ctx.JSON(400, gin.H{"message" : "Enter birthday!"})
		return
	}

	if strings.Contains(newUserDTO.Username, " ") {
		ctx.JSON(400, gin.H{"message" : "Username is not in valid format!"})
		return
	}

	customValidator := validator2.NewCustomValidator()
	translator, _ := customValidator.RegisterEnTranslation()
	errValidation := customValidator.Validator.Struct(newUserDTO)
	errs := customValidator.TranslateError(errValidation, translator)
	errorsString := customValidator.GetErrorsString(errs)

	if errValidation != nil {
		ctx.JSON(400, gin.H{"message" : errorsString[0]})
		return
	}

	error := p.ProfileInfoUseCase.EditUser(newUserDTO, ctx)
	if error != nil {
		ctx.JSON(http.StatusNotFound, "Failed to edit")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, "edit works")

}

type ProfileInfoHandler interface {
	GetProfileInfoByUsername(ctx *gin.Context)
	GetById(ctx *gin.Context)
	IsPrivate(ctx *gin.Context)
	GetProfileUsernameImageById(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	GetUserProfileById(ctx *gin.Context)
	SaveNewUser(ctx *gin.Context)
	GetAllPublicProfiles (ctx *gin.Context)
	EditUser(ctx *gin.Context)
	GetProfileInfoById(ctx *gin.Context)
	SearchUser(ctx *gin.Context)
	IsPrivatePostMethod(ctx *gin.Context)
	SearchPublicUser(ctx *gin.Context)
}
func NewProfileInfoHandler(usecase usecase.ProfileInfoUseCase) ProfileInfoHandler{
	return &profileInfoHandlder{ProfileInfoUseCase: usecase}
}

package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	logger "github.com/jelena-vlajkov/logger/logger"
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
	logger *logger.Logger
}



func (p *profileInfoHandlder) SearchPublicUser(ctx *gin.Context) {
	p.logger.Logger.Println("Handling SEARCH PUBLIC USERS")
	search := ctx.Request.URL.Query().Get("search")

	policy := bluemonday.UGCPolicy()
	search = strings.TrimSpace(policy.Sanitize(search))

	users, err := p.ProfileInfoUseCase.SearchPublicUsers(search, ctx)
	if err != nil {
		p.logger.Logger.Errorf("user does not exists, error: %v\n", err)
		ctx.JSON(404, "User does not exists")
		return
	}

	var usersDTO []dto.UserDTO
	for _, user := range users {
		usersDTO = append(usersDTO, dto.NewUserDTOfromEntity(*user))
	}

	if len(usersDTO) == 0 {
		p.logger.Logger.Errorf("user does not exists, error: %v\n", err)
		ctx.JSON(404, "User does not exists")
		return
	}

	ctx.JSON(200, usersDTO)

}

func (p *profileInfoHandlder) SearchUser(ctx *gin.Context) {
	p.logger.Logger.Println("Handling SEARCH USERS")
	search := ctx.Request.URL.Query().Get("search")

	policy := bluemonday.UGCPolicy()
	search = strings.TrimSpace(policy.Sanitize(search))

	users, err := p.ProfileInfoUseCase.SearchUser(search, ctx)
	if err != nil {
		p.logger.Logger.Errorf("user does not exists, error: %v\n", err)
		ctx.JSON(404, "User does not exists")
		return
	}

	var usersDTO []dto.SearchUserDTO
	for _, user := range users {
		usersDTO = append(usersDTO, dto.NewSearchUserDTOFromEntity(*user))
	}

	if len(usersDTO) == 0 {
		p.logger.Logger.Errorf("user does not exists, error: %v\n", err)
		ctx.JSON(404, "User does not exists")
		return
	}

	ctx.JSON(200, usersDTO)

}

func (p *profileInfoHandlder) IsPrivatePostMethod(ctx *gin.Context) {
	p.logger.Logger.Println("Handling IS PRIVATE POST METHOD")
	var privacyCheck dto.PrivacyCheckDto

	fmt.Println(ctx.Request.Body)

	if err := json.NewDecoder(ctx.Request.Body).Decode(&privacyCheck); err != nil {
		p.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
		ctx.JSON(500, "Error decoding body")
		return
	}

	policy := bluemonday.UGCPolicy()
	privacyCheck.Id = strings.TrimSpace(policy.Sanitize(privacyCheck.Id))


	isPrivate, err := p.ProfileInfoUseCase.IsPrivateById(privacyCheck.Id, ctx)

	if err != nil {
		p.logger.Logger.Errorf("can not get profile, error: %v\n", err)
		ctx.JSON(400, gin.H{"message" : "Can not get profile"})
		return
	}

	ctx.JSON(200, gin.H{"is_private" : isPrivate})

}

func (p *profileInfoHandlder) GetProfileInfoById(ctx *gin.Context) {
	p.logger.Logger.Println("Handling GETTING PROFILE INFO BY ID")
	id := ctx.Request.URL.Query().Get("userId")

	policy := bluemonday.UGCPolicy()
	id = strings.TrimSpace(policy.Sanitize(id))

	val, err := p.ProfileInfoUseCase.GetById(id, ctx)

	if err != nil {
		p.logger.Logger.Errorf("can not get profile info, error: %v\n", err)
		ctx.JSON(400, gin.H{"message" : "Can not get profile info"})
		return
	}

	if p.ProfileInfoUseCase.IsBanned(val, ctx) {
		p.logger.Logger.Errorf("user is baned, error: %v\n", err)
		ctx.JSON(400, gin.H{"message" : "User is baned"})
		return
	}

	profileInfoDto := mapper.ProfileInfoToProfileInfoDto(val)

	ctx.JSON(200, profileInfoDto)
}

func (p *profileInfoHandlder) GetProfileUsernameImageById(ctx *gin.Context) {
	p.logger.Logger.Println("Handling GETTING PROFILE USERNAME AND IMAGE BY ID")

	id := ctx.Request.URL.Query().Get("userId")

	policy := bluemonday.UGCPolicy()
	id = strings.TrimSpace(policy.Sanitize(id))

	profile, err1 := p.ProfileInfoUseCase.GetById(id, ctx)
	if err1 != nil {
		p.logger.Logger.Errorf("no users with that id, error: %v\n", err1)
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	var profileUIDTO dto.ProfileUsernameImageDTO
	profileUIDTO = dto.NewProfileUsernameImage(profile.Profile.Username, profile.ProfileImage)

	ctx.JSON(http.StatusOK, profileUIDTO)


}

func (p *profileInfoHandlder) GetProfileInfoByUsername(ctx *gin.Context) {
	p.logger.Logger.Println("Handling GETTING PROFILE INFO BY USERNAME")

	username := ctx.Request.URL.Query().Get("username")

	policy := bluemonday.UGCPolicy()
	username = strings.TrimSpace(policy.Sanitize(username))


	profileInfo, err := p.ProfileInfoUseCase.GetByUsername(username, ctx)

	if err != nil {
		p.logger.Logger.Errorf("no users with that username, error: %v\n", err)
		ctx.JSON(http.StatusNotFound, "No users with that username")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": profileInfo})

}

func (p *profileInfoHandlder) GetById(ctx *gin.Context) {
	p.logger.Logger.Println("Handling GETTING PROFILE BY ID")
	id := ctx.Request.URL.Query().Get("userId")

	policy := bluemonday.UGCPolicy()
	id = strings.TrimSpace(policy.Sanitize(id))

	profileInfo, err := p.ProfileInfoUseCase.GetById(id, ctx)

	if err != nil {
		p.logger.Logger.Errorf("no users with that id, error: %v\n", err)
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": profileInfo})
}

func (p *profileInfoHandlder) IsPrivate(ctx *gin.Context) {
	p.logger.Logger.Println("Handling IS PROFILE PRIVATE")

	id := ctx.Request.URL.Query().Get("userId")

	policy := bluemonday.UGCPolicy()
	id = strings.TrimSpace(policy.Sanitize(id))


	profile, err1 := p.ProfileInfoUseCase.GetById(id, ctx)

	if err1 != nil {
		p.logger.Logger.Errorf("no users with that id, error: %v\n", err1)
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	isPrivate, err2 := p.ProfileInfoUseCase.IsProfilePrivate(profile.Profile.Username, ctx)
	if err2 != nil {
		p.logger.Logger.Errorf("no users with that id, error: %v\n", err2)
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": isPrivate})

}

func (p *profileInfoHandlder) GetUserById(ctx *gin.Context) {
	p.logger.Logger.Println("Handling GETTING USER BY ID")
	id := ctx.Request.URL.Query().Get("userId")

	policy := bluemonday.UGCPolicy()
	id = strings.TrimSpace(policy.Sanitize(id))


	//DEKODER OVDE
	profileDTO, error := p.ProfileInfoUseCase.GetUserById(id, ctx)
	if error != nil{
		p.logger.Logger.Errorf("no users with that id, error: %v\n", error)
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, profileDTO)
}

func (p *profileInfoHandlder) GetUserProfileById(ctx *gin.Context) {
	p.logger.Logger.Println("Handling GETTING USER PROFILE ID")

	id := ctx.Request.URL.Query().Get("userId")
	policy := bluemonday.UGCPolicy()
	id = strings.TrimSpace(policy.Sanitize(id))


	//DEKODE OVDEE
	profileUserDTO, error := p.ProfileInfoUseCase.GetUserProfileById(id, ctx)
	if error != nil{
		p.logger.Logger.Errorf("no users with that id, error: %v\n", error)
		ctx.JSON(http.StatusNotFound, "No users with that id")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, profileUserDTO)


}

func (p *profileInfoHandlder) SaveNewUser(ctx *gin.Context) {
	p.logger.Logger.Println("Handling SAVING NEW USER")

	var newUserDTO dto.NewUserDTO


	err := json.NewDecoder(ctx.Request.Body).Decode(&newUserDTO)
	if err != nil {
		p.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
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

	if newUserDTO.ID == "" || newUserDTO.Name == "" || newUserDTO.Surname == "" || newUserDTO.Email == "" || newUserDTO.Address == "" || newUserDTO.Phone == "" || newUserDTO.Birthday  == "" ||
		newUserDTO.Gender == "" || newUserDTO.Web == "" || newUserDTO.Bio  == "" ||newUserDTO.Username == "" {
		p.logger.Logger.Errorf("fields are empty or xss attack happened")
		ctx.JSON(400, gin.H{"message" : "Fields are empty or xss attack happened"})
		return
	}

	if newUserDTO.Birthday == "" {
		ctx.JSON(400, gin.H{"message" : "Enter birthday!"})
		return
	}

	if strings.Contains(newUserDTO.Username, " ") {
		p.logger.Logger.Errorf("username is not in valid format!")
		ctx.JSON(400, gin.H{"message" : "Username is not in valid format!"})
		return
	}



	exists, _ := p.ProfileInfoUseCase.Exists(newUserDTO.Username, newUserDTO.Email, ctx)

	if exists {
		p.logger.Logger.Errorf("user already exists")
		ctx.JSON(400, gin.H{"message" : "User already exists"})
		return
	}

	if newUserDTO.Image != "" {
		mediaToAttach, err := p.ProfileInfoUseCase.EncodeBase64(newUserDTO.Image, newUserDTO.ID, context.Background())
		if err != nil {
			p.logger.Logger.Errorf("error while decoding base64, error: %v\n", err)
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
		p.logger.Logger.Errorf("error while validating, error: %v\n", errorsString[0])
		ctx.JSON(400, gin.H{"message" : errorsString[0]})
		return
	}

	error := p.ProfileInfoUseCase.SaveNewUser(newUserProfile, ctx)
	if error != nil {
		p.logger.Logger.Errorf("error while saving new user, error: %v\n", error)
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
	p.logger.Logger.Println("Handling EDITING USER")

	var newUserDTO dto.NewUserDTO

	err := json.NewDecoder(ctx.Request.Body).Decode(&newUserDTO)
	if err != nil {
		p.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
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

	if newUserDTO.ID == "" || newUserDTO.Name == "" || newUserDTO.Surname == "" || newUserDTO.Email == "" || newUserDTO.Address == "" || newUserDTO.Phone == "" || newUserDTO.Birthday  == "" ||
		newUserDTO.Gender == "" || newUserDTO.Web == "" || newUserDTO.Bio  == "" ||newUserDTO.Username == "" || newUserDTO.Image == "" {
		p.logger.Logger.Errorf("fields are empty or xss attack happeed")
		ctx.JSON(400, gin.H{"message" : "Field are empty or xss attack happened"})
		return
	}


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
		p.logger.Logger.Errorf("error while decoding json, error: %v\n", errorsString[0])
		ctx.JSON(400, gin.H{"message" : errorsString[0]})
		return
	}

	error := p.ProfileInfoUseCase.EditUser(newUserDTO, ctx)
	if error != nil {
		p.logger.Logger.Errorf("error while decoding json, error: %v\n", error)
		ctx.JSON(http.StatusNotFound, "Failed to edit")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, "edit works")

}

func (p *profileInfoHandlder) ChangePrivacyAndTaggin(ctx *gin.Context) {
	var changeDTO dto.PrivacyTaggingDTO

	err := json.NewDecoder(ctx.Request.Body).Decode(&changeDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Decoding error")
		ctx.Abort()
		return
	}

	error := p.ProfileInfoUseCase.ChangePrivacyAndTagging(changeDTO, ctx)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, "Failed to change privacy")
		ctx.Abort()
		return
	}

	ctx.JSON(200, "User updated")


}

func (p *profileInfoHandlder) GetPrivacyAndTagging(ctx *gin.Context) {
	id := ctx.Request.URL.Query().Get("userId")

	dto := p.ProfileInfoUseCase.GetPrivacyAndTagging(id, ctx)

	ctx.JSON(200, dto)

}

func (p *profileInfoHandlder) BanProfile(ctx *gin.Context) {

	var banProfile dto.BanProfileDTO
	err := json.NewDecoder(ctx.Request.Body).Decode(&banProfile)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Decoding error")
		ctx.Abort()
		return
	}

	result := p.ProfileInfoUseCase.BanUser(banProfile.ProfileId, ctx)
	if result == false {
		ctx.JSON(http.StatusBadRequest, "Failed to ban")
		ctx.Abort()
		return
	}

	ctx.JSON(200, "Profile banned")

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
	ChangePrivacyAndTaggin(ctx *gin.Context)
	GetPrivacyAndTagging(ctx *gin.Context)
	BanProfile(ctx *gin.Context)
}
func NewProfileInfoHandler(usecase usecase.ProfileInfoUseCase, logger *logger.Logger) ProfileInfoHandler{
	return &profileInfoHandlder{ProfileInfoUseCase: usecase, logger: logger}
}

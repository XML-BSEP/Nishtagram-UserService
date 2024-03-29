package usecase

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	logger "github.com/jelena-vlajkov/logger/logger"
	"io/ioutil"
	"os"
	"strings"
	"user-service/domain"
	"user-service/domain/enum"
	"user-service/dto"
	"user-service/gateway"
	"user-service/repository"
)

type profileInfoUseCase struct {
	ProfileInfoRepository repository.ProfileInfoRepository
	logger *logger.Logger
}



type ProfileInfoUseCase interface {
	GetByUsername(username string, ctx context.Context) (domain.ProfileInfo, error)
	GetAllProfiles(ctx context.Context) ([]domain.ProfileInfo, error)
	GetAllUserProfiles(ctx context.Context) ([]domain.ProfileInfo, error)
	GetById(id string, ctx context.Context) (*domain.ProfileInfo, error)
	GetUserById(id string, ctx context.Context) (dto.UserDTO, error)
	GetUserProfileById(id string, ctx context.Context) (dto.UserProfileDTO, error)
    IsProfilePrivate(username string, ctx context.Context) (bool, error)
	SaveNewUser(user domain.ProfileInfo, ctx context.Context) error
	Exists(username string, email string, ctx context.Context) (bool, error)
	EncodeBase64(media string, userId string, ctx context.Context) (string, error)
	GetAllPublicProfiles(ctx context.Context) ([]dto.UserDTO, error)
	DecodeBase64(media string, userId string, ctx context.Context) (string, error)
	EditUser(newUser dto.NewUserDTO, ctx context.Context) error
	IsBanned(user *domain.ProfileInfo, ctx context.Context) bool
	SearchUser(search string, ctx context.Context) ([]*domain.ProfileInfo, error)
	IsPrivateById(id string, ctx context.Context) (bool, error)
	SearchPublicUsers(search string, ctx context.Context) ([]*domain.ProfileInfo, error)
	ChangePrivacyAndTagging(taggingDTO dto.PrivacyTaggingDTO, ctx context.Context) error
	GetPrivacyAndTagging(profileId string, ctx context.Context) dto.PrivacyTaggingDTO
	BanUser(profileId string, ctx context.Context) bool
	IsInfluencerAndPrivate(profileId string, ctx context.Context) (*dto.InfluencerPrivateDTO, error)

}

func (p *profileInfoUseCase) IsPrivateById(id string, ctx context.Context) (bool, error) {
	p.logger.Logger.Infof("is private by id %v\n", id)
	return p.ProfileInfoRepository.IsPrivateById(id, ctx)
}

func (p *profileInfoUseCase) IsBanned(user *domain.ProfileInfo, ctx context.Context) bool {
	p.logger.Logger.Infof("is banned %v\n", user.ID)
	if user.Profile.PrivacyPermission == 2 {
		return true
	}
	return false
}

func (p *profileInfoUseCase) Exists(username string, email string, ctx context.Context) (bool, error) {
	p.logger.Logger.Infof("does exist by username %v, and email %v\n", username, email)
	return p.ProfileInfoRepository.Exists(username, email, ctx)
}

func (p *profileInfoUseCase) GetByUsername(username string, ctx context.Context) (domain.ProfileInfo, error) {
	p.logger.Logger.Infof("getting by username %v\n", username)
	return p.ProfileInfoRepository.GetByUsername(username, ctx)
}

func (p *profileInfoUseCase) GetAllProfiles(ctx context.Context) ([]domain.ProfileInfo, error) {
	 profiles, err := p.ProfileInfoRepository.GetAllProfiles(ctx)

	 if err != nil {
	 	return nil, err
	 }

	 var profileInfos []domain.ProfileInfo
	 for _, p := range profiles {
	 	if p.Profile.PrivacyPermission != 2 {
	 		profileInfos = append(profileInfos, p)
		}
	 }


	 return profileInfos, nil
}

func (p *profileInfoUseCase) GetAllUserProfiles(ctx context.Context) ([]domain.ProfileInfo, error) {
	profiles, err := p.ProfileInfoRepository.GetAllUserProfiles(ctx)

	if err != nil {
		return nil, err
	}

	var profileInfos []domain.ProfileInfo
	for _, p := range profiles {
		if p.Profile.PrivacyPermission != 2 {
			profileInfos = append(profileInfos, p)
		}
	}


	return profileInfos, nil
}

func (p *profileInfoUseCase) GetById(id string, ctx context.Context) (*domain.ProfileInfo, error) {
	p.logger.Logger.Infof("getting by id %v\n", id)
	user, _ :=  p.ProfileInfoRepository.GetById(id, ctx)

	if user.Profile.PrivacyPermission == 2 {
		return nil, nil
	}
	if user.ProfileImage != "" {
		encodedImage, _ := p.DecodeBase64(user.ProfileImage, user.ID, ctx)
		user.ProfileImage = encodedImage
	}


	return user, nil
}

func (p *profileInfoUseCase) GetUserById(id string, ctx context.Context) (dto.UserDTO, error) {
	p.logger.Logger.Infof("getting user by id %v\n", id)
	profile, _ := p.ProfileInfoRepository.GetUserById(id, ctx)

	if profile.Profile.PrivacyPermission == 2 {
		return dto.UserDTO{}, nil
	}

	var encodedImage string
	if profile.ProfileImage != "" {
		encodedImage, _ = p.DecodeBase64(profile.ProfileImage, profile.ID, ctx)
	}else {encodedImage = profile.ProfileImage }

	userDTO := dto.NewSimplyUserDTO(profile.Person.Name, profile.Person.Surname, profile.Email, profile.Person.Address,
		profile.Person.Phone, profile.Person.DateOfBirth.Format("02-Jan-2006"), profile.Person.Gender, profile.WebPage, profile.Biography,
		profile.Profile.Username, encodedImage, profile.Profile.PrivacyPermission.String())
	userDTO.Category = profile.Category.String()

	return userDTO, nil
}

func (p *profileInfoUseCase) GetUserProfileById(id string, ctx context.Context) (dto.UserProfileDTO, error) {
	p.logger.Logger.Infof("gettin user profile %v\n", id)
	profile, _ := p.ProfileInfoRepository.GetUserById(id, ctx)

	if profile.Profile.PrivacyPermission == 2 {
		return dto.UserProfileDTO{}, nil
	}

	var encodedImage string
	if profile.ProfileImage != "" {
		encodedImage, _ = p.DecodeBase64(profile.ProfileImage, profile.ID, ctx)
	}else {encodedImage = profile.ProfileImage }

	userDTO := dto.NewSimplyUserDTO(profile.Person.Name, profile.Person.Surname, profile.Email, profile.Person.Address,
		profile.Person.Phone, profile.Person.DateOfBirth.Format("02-Jan-2006"), profile.Person.Gender, profile.WebPage, profile.Biography,
		profile.Profile.Username, encodedImage, profile.Profile.PrivacyPermission.String())

	var private bool
	if profile.Profile.PrivacyPermission == 0 {
		private = true
	}else {
		private = false
	}
	userProfileDto := dto.NewUserProfileDTO(userDTO, &private)

	return  userProfileDto, nil
}

func (p *profileInfoUseCase) IsProfilePrivate(username string, ctx context.Context) (bool, error) {
	p.logger.Logger.Infof("is private by username %v\n", username)
	return  p.ProfileInfoRepository.IsProfilePrivate(username, ctx)
}

func (p *profileInfoUseCase) SaveNewUser(user domain.ProfileInfo, ctx context.Context) error {
	p.logger.Logger.Infof("saving new user with username %v\n", user.Profile.Username)
	return p.ProfileInfoRepository.SaveNewUser(user, ctx)
}

func (p *profileInfoUseCase) EncodeBase64(media string, userId string, ctx context.Context) (string, error) {
	p.logger.Logger.Infof("encoding base64 image for userId %v\n", userId)
	workingDirectory, _ := os.Getwd()
	if !strings.HasSuffix(workingDirectory, "src") {
		firstPart := strings.Split(workingDirectory, "src")
		value := firstPart[0] + "src"
		workingDirectory = value
		os.Chdir(workingDirectory)
	}

	path1 := "./assets/"
	err := os.Chdir(path1)
	if err != nil {
		p.logger.Logger.Errorf("error while encoding base64 image for userId %v, error: %v\n", userId, err)
	}
	err = os.Mkdir(userId, 0755)
	fmt.Println(err)

	err = os.Chdir(userId)
	fmt.Println(err)

	s := strings.Split(media, ",")
	a := strings.Split(s[0], "/")
	format := strings.Split(a[1], ";")
	dec, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		p.logger.Logger.Errorf("error while encoding base64 image for userId %v, error: %v\n", userId, err)
	}
	uuid := uuid.NewString()
	f, err := os.Create(uuid + "." + format[0])
	if err != nil {
		p.logger.Logger.Errorf("error while encoding base64 image for userId %v, error: %v\n", userId, err)
	}

	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		p.logger.Logger.Errorf("error while encoding base64 image for userId %v, error: %v\n", userId, err)
	}
	if err := f.Sync(); err != nil {
		p.logger.Logger.Errorf("error while encoding base64 image for userId %v, error: %v\n", userId, err)
	}

	os.Chdir(workingDirectory)
	return userId + "/" + uuid + "." + format[0], nil
}

func (p *profileInfoUseCase) GetAllPublicProfiles(ctx context.Context) ([]dto.UserDTO, error) {
	p.logger.Logger.Infof("getting all public profiles")
	users, err := p.ProfileInfoRepository.GetAllPublicProfiles(ctx)
	if err != nil {
		p.logger.Logger.Errorf("error while getting public profiles, error %v\n", err)
	}

	var usersNotBanned[] domain.ProfileInfo
	for _, u := range users {
		if u.Profile.PrivacyPermission != 2 {
			usersNotBanned = append(usersNotBanned, u)
		}
	}

	var usersDTO []dto.UserDTO
	for _, user := range usersNotBanned {
		if user.ProfileImage != ""{
			var img string
			img, _ = p.DecodeBase64(user.ProfileImage, user.ID, ctx)
			user.ProfileImage = img
		}
		usersDTO = append(usersDTO, dto.NewUserDTOfromEntity(user))
	}


	return usersDTO, nil


}

func (p *profileInfoUseCase) DecodeBase64(media string, userId string, ctx context.Context) (string, error) {
	p.logger.Logger.Infof("decoding base64 for image %v and user %v\n", media, userId)

	workingDirectory, _ := os.Getwd()
	if !strings.HasSuffix(workingDirectory, "src") {
		firstPart := strings.Split(workingDirectory, "src")
		value := firstPart[0] + "src"
		workingDirectory = value
		os.Chdir(workingDirectory)
	}

	path1 := "./assets/"
	err := os.Chdir(path1)
	fmt.Println(err)

	err = os.Chdir(userId)
	fmt.Println(err)

	spliced := strings.Split(media, "/")
	var f *os.File
	if len(spliced) > 1 {
		f, _ = os.Open(spliced[1])
	} else {
		f, _ = os.Open(spliced[0])
	}


	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)


	encoded := base64.StdEncoding.EncodeToString(content)

	os.Chdir(workingDirectory)

	return "data:image/jpg;base64," + encoded, nil
}


func (p *profileInfoUseCase) EditUser(newUser dto.NewUserDTO, ctx context.Context) error {
	p.logger.Logger.Infof("editing post for user, username %v\n", newUser.Username)
	_, err := p.GetUserById(newUser.ID, ctx)
	if err != nil {
		p.logger.Logger.Errorf("error while getting user by id, error %v\n", err)
		return err
	}

	var editedUser domain.ProfileInfo
	editedUser = dto.NewUserDTOtoEntity(newUser)
	if editedUser.ProfileImage != "" {
		image, err := p.EncodeBase64(editedUser.ProfileImage, editedUser.ID, ctx)
		if err != nil {
			p.logger.Logger.Errorf("error while encoding edited user image, error %v\n", err)
			return err
		}
		editedUser.ProfileImage = image
	}

	errRepo := p.ProfileInfoRepository.EditUser(editedUser, ctx)
	if err != nil {
		p.logger.Logger.Errorf("error while editting user, error %v\n", err)
		return errRepo
	}
	
	return nil


}


func (p *profileInfoUseCase) SearchUser(search string, ctx context.Context) ([]*domain.ProfileInfo, error) {
	p.logger.Logger.Infof("searching users, search = %v\n", search)

	var notBannedUsers []*domain.ProfileInfo
	searchedUsers, error := p.ProfileInfoRepository.SearchUser(search, ctx)
	if error != nil {
		p.logger.Logger.Errorf("getting searchhed users, error %v\n", error)
		return nil, error
	}
	for _, user := range searchedUsers {
		if user.Profile.PrivacyPermission.String() != "Banned" {
			var encodedImage string
			if user.ProfileImage != "" {
				encodedImage, _ = p.DecodeBase64(user.ProfileImage, user.ID, ctx)
			}else {encodedImage = user.ProfileImage }
			user.ProfileImage = encodedImage;
			notBannedUsers = append(notBannedUsers, user)
		}
	}



	return notBannedUsers, nil
}

func (p *profileInfoUseCase) SearchPublicUsers(search string, ctx context.Context) ([]*domain.ProfileInfo, error) {
	p.logger.Logger.Infof("searching public users, search = %v\n", search)

	var publics []*domain.ProfileInfo
	searchedUsers, error := p.ProfileInfoRepository.SearchUser(search, ctx)
	if error != nil {
		p.logger.Logger.Errorf("getting searched public users, error %v\n", error)
		return nil, error
	}
	for _, user := range searchedUsers {
		if user.Profile.PrivacyPermission.String() != "Banned" && user.Profile.PrivacyPermission.String() == "Public" {
			publics = append(publics, user)
		}
	}

	return publics, nil
}

func (p *profileInfoUseCase) ChangePrivacyAndTagging(taggingDTO dto.PrivacyTaggingDTO, ctx context.Context) error {

	var privacy enum.PrivacyPermission
	if taggingDTO.PrivacyPermission == "Private"{
		privacy = enum.PrivacyPermission(0)
	}else {
		privacy = enum.PrivacyPermission(1)
	}

	err := p.ProfileInfoRepository.ChangePrivacyAndTagging(privacy, taggingDTO.AllowTagging, taggingDTO.ProfileId, ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *profileInfoUseCase) GetPrivacyAndTagging(profileId string, ctx context.Context) dto.PrivacyTaggingDTO {
	profile,err := p.ProfileInfoRepository.GetById(profileId, ctx)
	if err != nil {
		return dto.PrivacyTaggingDTO{}
	}

	var privacyTaggingDTO dto.PrivacyTaggingDTO
	privacyTaggingDTO.AllowTagging = profile.Profile.AllowTagging
	privacyTaggingDTO.PrivacyPermission = profile.Profile.PrivacyPermission.String()
	privacyTaggingDTO.ProfileId = profileId

	return privacyTaggingDTO

}

func (p *profileInfoUseCase) BanUser(profileId string, ctx context.Context) bool {
	profile, _ := p.GetUserProfileById(profileId, ctx)
	response := p.ProfileInfoRepository.BanProfile(profileId, ctx)

	gateway.BanFollow(ctx, profileId)
	gateway.DeleteProfileInfo(ctx, profile.User.Username)

	return response
}

func (p *profileInfoUseCase) IsInfluencerAndPrivate(profileId string, ctx context.Context) (*dto.InfluencerPrivateDTO, error) {

	var influencerPrivateDTO dto.InfluencerPrivateDTO

	isInfluencer, err := p.ProfileInfoRepository.IsInfluencer(profileId, ctx)
	if err != nil {
		return nil, err
	}

	isPrivate, err := p.ProfileInfoRepository.IsPrivateById(profileId, ctx)
	if err != nil {
		return nil, err
	}

	influencerPrivateDTO.IsInfluencer = isInfluencer
	influencerPrivateDTO.IsPrivate = isPrivate

	return &influencerPrivateDTO, err

}

func NewProfileInfoUseCase(repo repository.ProfileInfoRepository, logger *logger.Logger) ProfileInfoUseCase {
	return &profileInfoUseCase{ ProfileInfoRepository: repo, logger: logger}
}


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
	"user-service/dto"
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
	 return p.ProfileInfoRepository.GetAllProfiles(ctx)
}

func (p *profileInfoUseCase) GetAllUserProfiles(ctx context.Context) ([]domain.ProfileInfo, error) {
	return p.ProfileInfoRepository.GetAllUserProfiles(ctx)
}

func (p *profileInfoUseCase) GetById(id string, ctx context.Context) (*domain.ProfileInfo, error) {
	p.logger.Logger.Infof("getting by id %v\n", id)
	user, _ :=  p.ProfileInfoRepository.GetById(id, ctx)
	if user.ProfileImage != "" {
		encodedImage, _ := p.DecodeBase64(user.ProfileImage, user.ID, ctx)
		user.ProfileImage = encodedImage
	}
	return user, nil
}

func (p *profileInfoUseCase) GetUserById(id string, ctx context.Context) (dto.UserDTO, error) {
	p.logger.Logger.Infof("getting user by id %v\n", id)
	profile, _ := p.ProfileInfoRepository.GetUserById(id, ctx)

	var encodedImage string
	if profile.ProfileImage != "" {
		encodedImage, _ = p.DecodeBase64(profile.ProfileImage, profile.ID, ctx)
	}else {encodedImage = profile.ProfileImage }

	userDTO := dto.NewSimplyUserDTO(profile.Person.Name, profile.Person.Surname, profile.Email, profile.Person.Address,
		profile.Person.Phone, profile.Person.DateOfBirth.Format("02-Jan-2006"), profile.Person.Gender, profile.WebPage, profile.Biography,
		profile.Profile.Username, encodedImage, profile.Profile.PrivacyPermission.String())


	return userDTO, nil
}

func (p *profileInfoUseCase) GetUserProfileById(id string, ctx context.Context) (dto.UserProfileDTO, error) {
	p.logger.Logger.Infof("gettin user profile %v\n", id)
	profile, _ := p.ProfileInfoRepository.GetUserById(id, ctx)

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

	var usersDTO []dto.UserDTO
	for _, user := range users {
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

func NewProfileInfoUseCase(repo repository.ProfileInfoRepository, logger *logger.Logger) ProfileInfoUseCase {
	return &profileInfoUseCase{ ProfileInfoRepository: repo, logger: logger}
}


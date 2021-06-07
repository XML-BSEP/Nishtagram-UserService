package usecase

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strings"
	"user-service/domain"
	"user-service/dto"
	"user-service/repository"
)

type profileInfoUseCase struct {
	ProfileInfoRepository repository.ProfileInfoRepository
}

type ProfileInfoUseCase interface {
	GetByUsername(username string, ctx context.Context) (domain.ProfileInfo, error)
	GetAllProfiles(ctx context.Context) ([]domain.ProfileInfo, error)
	GetAllUserProfiles(ctx context.Context) ([]domain.ProfileInfo, error)
	GetById(id string, ctx context.Context) (domain.ProfileInfo, error)
	GetUserById(id string, ctx context.Context) (dto.UserDTO, error)
	GetUserProfileById(id string, ctx context.Context) (dto.UserProfileDTO, error)
    IsProfilePrivate(username string, ctx context.Context) (bool, error)
	SaveNewUser(user domain.ProfileInfo, ctx context.Context) error
	Exists(username string, email string, ctx context.Context) (bool, error)
	EncodeBase64(media string, userId string, ctx context.Context) (string, error)
}

func (p *profileInfoUseCase) Exists(username string, email string, ctx context.Context) (bool, error) {
	return p.ProfileInfoRepository.Exists(username, email, ctx)
}

func (p *profileInfoUseCase) GetByUsername(username string, ctx context.Context) (domain.ProfileInfo, error) {
	return p.ProfileInfoRepository.GetByUsername(username, ctx)
}

func (p *profileInfoUseCase) GetAllProfiles(ctx context.Context) ([]domain.ProfileInfo, error) {
	 return p.ProfileInfoRepository.GetAllProfiles(ctx)
}

func (p *profileInfoUseCase) GetAllUserProfiles(ctx context.Context) ([]domain.ProfileInfo, error) {
	return p.ProfileInfoRepository.GetAllUserProfiles(ctx)
}

func (p *profileInfoUseCase) GetById(id string, ctx context.Context) (domain.ProfileInfo, error) {
	return  p.ProfileInfoRepository.GetById(id, ctx)
}

func (p *profileInfoUseCase) GetUserById(id string, ctx context.Context) (dto.UserDTO, error) {
	return  p.ProfileInfoRepository.GetUserById(id, ctx)
}

func (p *profileInfoUseCase) GetUserProfileById(id string, ctx context.Context) (dto.UserProfileDTO, error) {
	return  p.ProfileInfoRepository.GetUserProfileById(id, ctx)
}


func (p *profileInfoUseCase) IsProfilePrivate(username string, ctx context.Context) (bool, error) {
	return  p.ProfileInfoRepository.IsProfilePrivate(username, ctx)
}

func (p *profileInfoUseCase) SaveNewUser(user domain.ProfileInfo, ctx context.Context) error {
	return p.ProfileInfoRepository.SaveNewUser(user, ctx)
}

func (p *profileInfoUseCase) EncodeBase64(media string, userId string, ctx context.Context) (string, error) {
	workingDirectory, _ := os.Getwd()
	path1 := "../assets"
	err := os.Chdir(path1)
	if err != nil {
		fmt.Println(err)
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
		panic(err)
	}
	uuid := uuid.NewString()
	f, err := os.Create(uuid + "." + format[0])
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}

	os.Chdir(workingDirectory)
	return userId + "/" + uuid + "." + format[0], nil
}


func NewProfileInfoUseCase(repo repository.ProfileInfoRepository) ProfileInfoUseCase {
	return &profileInfoUseCase{ ProfileInfoRepository: repo}
}


package usecase

import (
	"context"
	"user-service/domain"
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

func NewProfileInfoUseCase(repo repository.ProfileInfoRepository) ProfileInfoUseCase {
	return &profileInfoUseCase{ ProfileInfoRepository: repo}
}


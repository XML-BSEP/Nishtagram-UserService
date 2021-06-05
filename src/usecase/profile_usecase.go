package usecase

import (
	"context"
	"user-service/domain"
	"user-service/repository"
)

type profileUseCase struct {
	ProfileRepository repository.ProfileRepository
}



type ProfileUseCase interface {
	GetByUsername(username string, ctx context.Context) (domain.Profile, error)
	GetAllProfiles(ctx context.Context) ([]domain.Profile, error)
	GetAllUserProfiles(ctx context.Context) ([]domain.Profile, error)
	IsProfilePrivate(username string, ctx context.Context) (bool, error)
}

func (p *profileUseCase) GetByUsername(username string, ctx context.Context) (domain.Profile, error) {
	return p.ProfileRepository.GetByUsername(username, ctx)
}

func (p *profileUseCase) GetAllProfiles(ctx context.Context) ([]domain.Profile, error) {
	return p.ProfileRepository.GetAllProfiles(ctx)
}

func (p *profileUseCase) GetAllUserProfiles(ctx context.Context) ([]domain.Profile, error) {
	return p.ProfileRepository.GetAllUserProfiles(ctx)
}

func (p *profileUseCase) IsProfilePrivate(username string, ctx context.Context) (bool, error) {
	return p.ProfileRepository.IsProfilePrivate(username, ctx)
}

func NewProfileUseCase(repo repository.ProfileRepository) ProfileUseCase {
	return &profileUseCase{ ProfileRepository: repo}
}


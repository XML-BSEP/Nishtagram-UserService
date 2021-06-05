package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"user-service/domain"
	"user-service/domain/enum"
	"user-service/dto"
)

type profileInfoRepository struct {
	collection *mongo.Collection
	db *mongo.Client
}



type ProfileInfoRepository interface {
	GetByUsername(username string, ctx context.Context) (domain.ProfileInfo, error)
	GetAllProfiles(ctx context.Context) ([]domain.ProfileInfo, error)
	GetAllUserProfiles(ctx context.Context) ([]domain.ProfileInfo, error)
	GetById(id string, ctx context.Context) (domain.ProfileInfo, error)
	GetUserById(id string, ctx context.Context) (dto.UserDTO, error)
	GetUserProfileById(id string, ctx context.Context) (dto.UserProfileDTO, error)
}


func (p *profileInfoRepository) GetByUsername(username string, ctx context.Context) (domain.ProfileInfo, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var profile domain.ProfileInfo
	err := p.collection.FindOne(ctx, bson.M{"profile.username" : username}).Decode(&profile)
	if err != nil {
		return profile, err
	}
	return profile, nil
}

func (p *profileInfoRepository) GetAllProfiles(ctx context.Context) ([]domain.ProfileInfo, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	profiles , err := p.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var allProfiles []domain.ProfileInfo
	if err = profiles.All(ctx, &allProfiles); err != nil {
		return nil, err
	}

	return allProfiles, nil
}

func (p *profileInfoRepository) GetAllUserProfiles(ctx context.Context) ([]domain.ProfileInfo, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	profiles , err := p.collection.Find(ctx, bson.M{"profile.type" : enum.ProfileType(0).String()})
	if err != nil {
		return nil, err
	}

	var allProfiles []domain.ProfileInfo
	if err = profiles.All(ctx, &allProfiles); err != nil {
		return nil, err
	}

	return allProfiles, nil
}

func (p *profileInfoRepository) GetById(id string, ctx context.Context) (domain.ProfileInfo, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var profile domain.ProfileInfo
	err := p.collection.FindOne(ctx, bson.M{"person._id" : id}).Decode(&profile)
	if err != nil {
		return profile, err
	}
	return profile, nil
}

func (p *profileInfoRepository) GetUserById(id string, ctx context.Context) (dto.UserDTO, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var profile domain.ProfileInfo
	var userDTO dto.UserDTO
	err := p.collection.FindOne(ctx, bson.M{"person._id" : id}).Decode(&profile)
	if err != nil {
		return userDTO, err
	}

	userDTO = dto.NewSimplyUserDTO(profile.Person.Name, profile.Person.Surname, profile.Email, profile.Person.Address,
		profile.Person.Phone, profile.Person.DateOfBirth.Format("02-Jan-2006"), profile.Person.Gender, profile.WebPage, profile.Biography,
		profile.Profile.Username, profile.ProfileImage)

	return userDTO, nil

}

func (p *profileInfoRepository) GetUserProfileById(id string, ctx context.Context) (dto.UserProfileDTO, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var profile domain.ProfileInfo
	err := p.collection.FindOne(ctx, bson.M{"person._id" : id}).Decode(&profile)
	if err != nil {
		return dto.UserProfileDTO{}, err
	}

	userDTO, err1 := p.GetUserById(id, ctx)
	if err1 != nil {
		return dto.UserProfileDTO{}, err
	}

	var isPrivate bool
	if profile.Profile.PrivacyPermission == 0 {
		isPrivate = true
	}

	if profile.Profile.PrivacyPermission == 1 {
		isPrivate = false
	}

	userProfileDTO := dto.NewUserProfileDTO(userDTO, &isPrivate)

	return userProfileDTO, nil

}

func NewProfileInfoRepository(db *mongo.Client) ProfileInfoRepository {
	return &profileInfoRepository {
		db : db,
		collection: db.Database("user_db").Collection("profiles_info"),
	}
}

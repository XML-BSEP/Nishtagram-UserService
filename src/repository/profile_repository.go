package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"user-service/domain"
	"user-service/domain/enum"
)

type profileRepository struct {
	collection *mongo.Collection
	db *mongo.Client
}


type ProfileRepository interface {
	GetByUsername(username string, ctx context.Context) (domain.Profile, error)
	GetAllProfiles(ctx context.Context) ([]domain.Profile, error)
	GetAllUserProfiles(ctx context.Context) ([]domain.Profile, error)
	IsProfilePrivate(username string, ctx context.Context) (bool, error)

}

func (p *profileRepository) GetByUsername(username string, ctx context.Context) (domain.Profile, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var profile domain.Profile
	err := p.collection.FindOne(ctx, bson.M{"username" : username}).Decode(&profile)
	if err != nil {
		return profile, err
	}
	return profile, nil

}

func (p *profileRepository) GetAllProfiles(ctx context.Context) ([]domain.Profile, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	profiles , err := p.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var allProfiles []domain.Profile
	if err = profiles.All(ctx, &allProfiles); err != nil {
		return nil, err
	}

	return allProfiles, nil
}

func (p *profileRepository) GetAllUserProfiles(ctx context.Context) ([]domain.Profile, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	profiles , err := p.collection.Find(ctx, bson.M{"type" : enum.ProfileType(0)})
	if err != nil {
		return nil, err
	}

	var allProfiles []domain.Profile
	if err = profiles.All(ctx, &allProfiles); err != nil {
		return nil, err
	}

	return allProfiles, nil
}

func (p *profileRepository) IsProfilePrivate(username string, ctx context.Context) (bool, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()


	var profile domain.Profile
	err := p.collection.FindOne(ctx, bson.M{"username" : username}).Decode(&profile)
	if err != nil {
		return false, err
	}

	if profile.PrivacyPermission.String() == "Private" {
		return true, nil
	}

	return false, nil

}

func NewProfileRepository(db *mongo.Client) ProfileRepository {
	return &profileRepository {
		db : db,
		collection: db.Database("user_db").Collection("profiles"),
	}
}
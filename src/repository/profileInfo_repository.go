package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"user-service/domain"
	"user-service/domain/enum"
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

func NewProfileInfoRepository(db *mongo.Client) ProfileInfoRepository {
	return &profileInfoRepository {
		db : db,
		collection: db.Database("user_db").Collection("profiles_info"),
	}
}

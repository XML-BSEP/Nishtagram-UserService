package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"user-service/domain"
	"user-service/domain/enum"
)

type requestVerificationRepository struct {
	collection *mongo.Collection
	db *mongo.Client
}



type RequestVerificationRepository interface {
	SaveNewRequestVerification(verification domain.RequestVerification, ctx context.Context) (bool, error)
	GetAllRequestVerificationForWaiting(ctx context.Context) *[]domain.RequestVerification
	GetAllRequestVerification(ctx context.Context) *[]domain.RequestVerification
	ApproveRequestVerification(verificationId string, ctx context.Context) (bool, error)
	RejectRequestVerification(verificationId string, ctx context.Context) (bool, error)
	ExistsRequestForProfile(profileId string, ctx context.Context) (bool,error)
	GetByProfileId(profileId string, ctx context.Context) (domain.RequestVerification, error)
}


func (r *requestVerificationRepository) ExistsRequestForProfile(profileId string, ctx context.Context) (bool,error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()


	var profile domain.RequestVerification
	err := r.collection.FindOne(ctx, bson.M{"profile_id" : profileId}).Decode(&profile)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *requestVerificationRepository) SaveNewRequestVerification(verification domain.RequestVerification, ctx context.Context) (bool, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()


	_, err := r.collection.InsertOne(ctx, verification)
	if err != nil {
		return false, err
	}

	return true, nil


}

func (r *requestVerificationRepository) GetAllRequestVerificationForWaiting(ctx context.Context) *[]domain.RequestVerification {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	verifications, err := r.collection.Find(ctx, bson.M{"state" : enum.VerificationState(0)})
	if err != nil {
		return nil
	}

	var allVerifications []domain.RequestVerification
	if err = verifications.All(ctx, &allVerifications); err != nil {
		return nil
	}

	return &allVerifications
}

func (r *requestVerificationRepository) GetAllRequestVerification(ctx context.Context) *[]domain.RequestVerification {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	verifications, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil
	}

	var allVerifications []domain.RequestVerification
	if err = verifications.All(ctx, &allVerifications); err != nil {
		return nil
	}

	return &allVerifications

}

func (r *requestVerificationRepository) ApproveRequestVerification(verificationId string, ctx context.Context) (bool, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	verificationToApprove := bson.M{"_id" : verificationId}
	updateVerification := bson.M{"$set" : bson.M {
		"state" : enum.VerificationState(1),
	}}

	_, err := r.collection.UpdateOne(ctx, verificationToApprove, updateVerification)
	if err != nil {
		return  false, err
	}

	return true, nil

}

func (r *requestVerificationRepository) GetByProfileId(profileId string, ctx context.Context) (domain.RequestVerification, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var req domain.RequestVerification
	err := r.collection.FindOne(ctx, bson.M{"profile_id" : profileId}).Decode(&req)

	if err != nil {
		return req, err
	}
	return req, err
}


func (r *requestVerificationRepository) RejectRequestVerification(verificationId string, ctx context.Context) (bool, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	verificationToApprove := bson.M{"_id" : verificationId}
	updateVerification := bson.M{"$set" : bson.M {
		"state" : enum.VerificationState(2),
	}}

	_, err := r.collection.UpdateOne(ctx, verificationToApprove, updateVerification)
	if err != nil {
		return  false, err
	}

	return true, nil
}


func NewRequestVerificationRepository(db *mongo.Client) RequestVerificationRepository {
	return &requestVerificationRepository {
		db : db,
		collection: db.Database("user_db").Collection("request_verifications"),
	}
}


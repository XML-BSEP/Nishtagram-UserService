package dto

type RequestVerificationToChangeStateDTO struct {
	ID string `bson:"_id" json:"id"`
	ProfileId string `bson:"profile_id" json:"profile_id"`
}
package domain

import "user-service/domain/enum"

type RequestVerification struct {
	ID string `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	Surname string `bson:"surname" json:"surname"`
	Category enum.Category `bson:"category" json:"category"`
	Image string `bson:"image" json:"image"`
}

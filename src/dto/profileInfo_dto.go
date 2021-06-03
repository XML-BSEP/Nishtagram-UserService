package dto

import (
	"user-service/domain/enum"
)

type ProfileInfoDTO struct {
	Email string `bson:"email" json:"email"`
	Biography string `bson:"biography" json:"biography"`
	WebPage string `bson:"web_page" json:"web_page"`
	Category enum.Category `json:"category" bson:"category"`
	ProfileImage string `bson:"profile_image" json:"profile_image"`
	Person PersonDTO `bson:"person" json:"person"`
	Profile ProfileDTO `bson:"profile" json:"profile"`

}



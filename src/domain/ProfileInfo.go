package domain

import "user-service/domain/enum"

type ProfileInfo struct {
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Biography string `bson:"biography" json:"biography"`
	WebPage string `bson:"web_page" json:"web_page"`
	Category enum.Category `json:"category" bson:"category"`
	ProfileImage string `bson:"profile_image" json:"profile_image"`
	Person Person `bson:"person" json:"person"`
	Profile Profile `bson:"profile" json:"profile"`
}

package domain

import "user-service/domain/enum"

type ProfileInfo struct {
	ID string `bson:"_id" json:"id"`
	Email string `bson:"email" json:"email" validate:"required,email"`
	Biography string `bson:"biography" json:"biography"`
	WebPage string `bson:"web_page" json:"web_page"`
	Category enum.Category `json:"category" bson:"category"`
	ProfileImage string `bson:"profile_image" json:"profile_image"`
	Person Person `bson:"person" json:"person"`
	Profile Profile `bson:"profile" json:"profile"`
}

func NewProfileInfo(email string, bio string, web string, cat enum.Category, image string, person Person, profile Profile) ProfileInfo {
	return ProfileInfo{Email: email, Biography: bio, WebPage: web, Category: cat, ProfileImage: image, Person: person, Profile: profile}
}

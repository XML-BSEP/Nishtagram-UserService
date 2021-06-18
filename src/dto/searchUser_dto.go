package dto

import (
	"user-service/domain"
)

type SearchUserDTO struct {
	Id		 string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Surname  string `bson:"surname" json:"surname"`
	Username string `bson:"username" json:"username"`
	Private bool `bson:"private" json:"private"`
	Image	string `json:"image" bson:"image"`
}


func NewSearchUserDTOFromEntity(profile domain.ProfileInfo) SearchUserDTO{

	var private bool
	if profile.Profile.PrivacyPermission == 0 {
		private = true
	}else {
		private = false
	}

	return SearchUserDTO{Name:profile.Person.Name, Surname: profile.Person.Surname,
   Username: profile.Profile.Username, Image: profile.ProfileImage, Private: private, Id: profile.ID}
}


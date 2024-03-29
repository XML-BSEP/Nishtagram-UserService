package dto

import (
	"user-service/domain"
	"user-service/domain/enum"
)

type UserDTO struct {
	Name     string `bson:"name" json:"name"`
	Surname  string `bson:"surname" json:"surname"`
	Email    string `bson:"email" json:"email"`
	Address  string `bson:"address" json:"address"`
	Phone    string `bson:"phone" json:"phone"`
	Birthday string `bson:"birthday" json:"birthday"`
	Gender   string `bson:"gender" json:"gender"`
	Web      string `bson:"web" json:"web"`
	Bio      string `bson:"bio" json:"bio"`
	Username string `bson:"username" json:"username"`
	Image    string `bson:"image" json:"image"`
	Private bool `bson:"private" json:"private"`
	Category string `bson:"category" json:"category"`
}

type UserIdsDto struct{
	Ids []string `json:"ids"`
}

func NewSimplyUserDTO(name string, surname string, email string, address string, phone string, birthday string,
	gender enum.Gender, web string, bio string, username string, image string, private string) UserDTO{

	var genderString string
	if gender == 0 {
		genderString = "Male"
	}else if gender == 1 {
		genderString = "Female"
	}else {
		genderString = "Other"
	}

	var privacy bool
	if private == "Private" {
		privacy = true
	} else {
		privacy = false
	}

	return UserDTO{
		Name: name,
		Surname: surname,
		Email: email,
		Address: address,
		Phone: phone,
		Birthday: birthday,
		Gender: genderString,
		Web: web,
		Bio: bio,
		Username: username,
		Image: image,
		Private: privacy,
	}
}

func NewUserDTOfromEntity(profile domain.ProfileInfo) UserDTO{

	var private bool
	if profile.Profile.PrivacyPermission == 0 {
		private = true
	}else {
		private = false
	}

	return UserDTO{Name:profile.Person.Name, Surname: profile.Person.Surname, Email: profile.Email, Address: profile.Person.Address,
		Phone: profile.Person.Phone, Birthday: profile.Person.DateOfBirth.String(), Gender: profile.Person.Gender.String(), Web: profile.WebPage,
	Bio: profile.WebPage, Username: profile.Profile.Username, Image: profile.ProfileImage, Private: private}
}


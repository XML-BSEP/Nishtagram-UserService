package dto

import (
	"log"
	"time"
	"user-service/domain"
	"user-service/domain/enum"
)

type NewUserDTO struct {
	ID       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name" validate:"required,name"`
	Surname  string `bson:"surname" json:"surname" validate:"required,surname"`
	Email    string `bson:"email" json:"email" validate:"required,email"`
	Address  string `bson:"address" json:"address"`
	Phone    string `bson:"phone" json:"phone" validate:"required,phone"`
	Birthday string `bson:"birthday" json:"birthday"`
	Gender   string `bson:"gender" json:"gender"`
	Web      string `bson:"web" json:"web"`
	Bio      string `bson:"bio" json:"bio"`
	Username string `bson:"username" json:"username" validate:"required,username"`
	Image    string `bson:"image" json:"image"`
	Private bool `bson:"private" json:"private"`

}

func NewUserDTOtoEntity(newUserDto NewUserDTO) domain.ProfileInfo {

	var gender enum.Gender
	if newUserDto.Gender == "Male" {
		gender = enum.Gender(0)
	}else if newUserDto.Gender == "Female" {
		gender = enum.Gender(1)
	}else {
		gender = enum.Gender(2)
	}

	birth, err := time.Parse("2006-01-02T15:04:05.000Z", newUserDto.Birthday)
	if err != nil {
		log.Fatal(err)
	}

	var privacy enum.PrivacyPermission
	if newUserDto.Private == true {
		privacy = enum.PrivacyPermission(0)
	}else {
		privacy = enum.PrivacyPermission(1)
	}


	person := domain.NewPerson(newUserDto.Name, newUserDto.Surname, gender, birth, newUserDto.Address, newUserDto.Phone)
	profile := domain.NewProfile(newUserDto.Username, privacy, true, true, enum.ProfileType(0))
	profileInfo := domain.NewProfileInfo(newUserDto.Email, newUserDto.Bio, newUserDto.Web, enum.Category(6), newUserDto.Image, person, profile)
	profileInfo.ID = newUserDto.ID
	return profileInfo


}

package dto

import (
	"time"
	"user-service/domain/enum"
)

type PersonDTO struct {
	ID string `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	Surname string `bson:"surname" json:"surname"`
	Gender	enum.Gender `bson:"gender" json:"gender"`
	DateOfBirth time.Time `bson:"date_of_birth" json:"date_of_birth"`
	Address string `bson:"address" json:"address"`
	Phone string `bson:"phone" json:"phone"`
}

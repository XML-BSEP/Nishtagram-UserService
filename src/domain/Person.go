package domain

import (
	"time"
	"user-service/domain/enum"
)

type Person struct {

	Name string `bson:"name" json:"name"`
	Surname string `bson:"surname" json:"surname"`
	Gender	enum.Gender `bson:"gender" json:"gender"`
	DateOfBirth time.Time `bson:"date_of_birth" json:"date_of_birth"`
	Address string `bson:"address" json:"address"`
	Phone string `bson:"phone" json:"phone"`

}

func NewPerson(name string, surname string, gender enum.Gender, birth time.Time, address string, phone string) Person {
	return Person{Name: name, Surname: surname, Gender: gender, DateOfBirth: birth, Address: address, Phone: phone}
}

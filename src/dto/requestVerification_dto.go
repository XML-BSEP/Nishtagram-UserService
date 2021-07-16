package dto



type RequestVerificationDTO struct {
	Name string `bson:"name" json:"name"`
	Surname string `bson:"surname" json:"surname"`
	Category string `bson:"category" json:"category"`
	Image string `bson:"image" json:"image"`
	ProfileId string `bson:"profile_id" json:"profile_id"`
	State string `bson:"state" json:"state"`
	Id string `bson:"id" json:"id"`
}



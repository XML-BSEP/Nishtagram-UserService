package dto

type AdminDTO struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Person PersonDTO `bson:"person" json:"person"`
}

package dto

type userProfileDTO struct {
	Username string `bson:"username" json:"username"`
	FullName string `bson:"full_name" json:"full_name"`
	
}

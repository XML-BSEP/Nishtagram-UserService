package dto

type ProfileUsernameImageDTO struct {
	Username string `bson:"username" json:"username"`
	Image string `bson:"image" json:"image"`
}

func NewProfileUsernameImage(username string, image string) ProfileUsernameImageDTO{
	return ProfileUsernameImageDTO {
		Username: username,
		Image : image,
	}
}



package dto

type ProfileUsernameImageDTO struct {
	Username string `bson:"username" json:"username"`
	ProfilePhoto string `bson:"profile_photo" json:"profile_photo"`
}

func NewProfileUsernameImage(username string, image string) ProfileUsernameImageDTO{
	return ProfileUsernameImageDTO {
		Username: username,
		ProfilePhoto : image,
	}
}



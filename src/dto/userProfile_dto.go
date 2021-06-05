package dto

type UserProfileDTO struct {
	/*user : User;
	  followers : UserInFeed[];
	  following : Following[];
	  posts : PostInProfile[];
	  private : boolean;*/

	User UserDTO `bson:"user" json:"user"`
	Private *bool `bson:"private" json:"private" binding:"exists"`
	Followers []string `bson:"followers" json:"followers"`
	Following []string `bson:"following" json:"following"`

}

func NewUserProfileDTO(user UserDTO, private *bool) UserProfileDTO {
	return UserProfileDTO{User: user, Private: private}
}


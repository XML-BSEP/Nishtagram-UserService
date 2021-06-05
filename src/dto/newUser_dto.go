package dto

type NewUserDTO struct {
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
	Password string `bson:"password" json:"password"`
	ConfirmPassword string `bson:"confirm_password" json:"confirm_password"`
}

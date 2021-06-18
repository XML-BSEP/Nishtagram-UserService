package domain

type Admin struct {
	ID string `bson:"_id" json:"id"`
	Username string `bson:"username" json:"username"`
	Person Person `bson:"person" json:"person"`
}

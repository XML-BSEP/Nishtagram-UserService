package dto

type ProfileInfoDto struct {
	Id string `json:"id"`
	Usernmae string `json:"usernmae"`
	Address string `json:"address"`
	Bio string `json:"bio"`
	WebSite string `json:"web_site"`
	ProfileImage string `json:"profile_image"`
	IsPrivate bool `json:"is_private"`
}

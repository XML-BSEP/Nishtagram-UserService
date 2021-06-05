package domain

import "user-service/domain/enum"

type Profile struct {
	Username string `bson:"username" json:"username"`
	PrivacyPermission enum.PrivacyPermission `bson:"privacy_permission" json:"privacy_permission"`
	AllowTagging bool `bson:"allow_tagging" json:"allow_tagging"`
	AllowNotification bool `bson:"allow_notification" json:"allow_notification"`
	Type enum.ProfileType `bson:"type" json:"type"`
}

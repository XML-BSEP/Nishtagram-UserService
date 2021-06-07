package mapper

import (
	"user-service/domain"
	"user-service/domain/enum"
	"user-service/dto"
)

func ProfileInfoToProfileInfoDto(profileInfo *domain.ProfileInfo) dto.ProfileInfoDto {
	return dto.ProfileInfoDto{
		Id: profileInfo.ID,
		Address: profileInfo.Person.Address,
		Bio: profileInfo.Biography,
		WebSite: profileInfo.WebPage,
		ProfileImage: profileInfo.ProfileImage,
		IsPrivate: privacyPermissionToPrivate(profileInfo.Profile.PrivacyPermission),
	}
}

func privacyPermissionToPrivate(permission enum.PrivacyPermission) bool {
	if permission == 0 {
		return true
	}

	return false
}

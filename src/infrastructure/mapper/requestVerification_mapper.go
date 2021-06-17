package mapper

import (
	"user-service/domain"
	"user-service/domain/enum"
	"user-service/dto"
)

func NewVerificationRequestDTOtoEntity(dto dto.RequestVerificationDTO) domain.RequestVerification {

	var reqVerification domain.RequestVerification
	reqVerification.Name = dto.Name
	reqVerification.Surname = dto.Surname
	reqVerification.Image = dto.Image
	reqVerification.State = enum.VerificationState(0)

	var category enum.Category
	if dto.Category == "Influencer" {
		category = enum.Category(0)
	}else if dto.Category == "Sports" {
		category = enum.Category(1)
	}else if dto.Category == "NewMedia" {
		category = enum.Category(2)
	}else if dto.Category == "Business" {
		category = enum.Category(3)
	}else if dto.Category == "Brand" {
		category = enum.Category(4)
	}else if dto.Category == "Organization" {
		category = enum.Category(5)
	}

	reqVerification.Category = category
	reqVerification.ProfileId = dto.ProfileId
	return reqVerification

}

func RequestVerificationToDTO(verification domain.RequestVerification) dto.RequestVerificationDTO {
	return dto.RequestVerificationDTO{Name: verification.Name,
				Surname: verification.Surname,
				Category: verification.Category.String(),
				Image: verification.Image,
				ProfileId: verification.ProfileId,
				State: verification.State.String(),
	}

}
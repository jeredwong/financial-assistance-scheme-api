package mapper

import (
	"github.com/jeredwong/financial-scheme-manager/internal/constants"
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
)

func ApplicationDTOToModel(applicationDTO dto.ApplicationDTO) models.Application {
	return models.Application{
		Status: constants.ApplicationStatus(applicationDTO.Status),
		ApplicantId: applicationDTO.Applicant.Id,
		SchemeId: applicationDTO.Scheme.Id,
	}
}

func ApplicationModelToDTO(application models.Application) dto.ApplicationDTO {
	return dto.ApplicationDTO{
		Id: application.ID,
		Status: string(application.Status),
		ApplicationDate: application.ApplicationDate,
		Applicant: dto.ApplicantDTO{},
		Scheme: dto.SchemeDTO{},
	}
}

// TODO: refactor nested mapping of DTOs in mappers instead of within handlers?
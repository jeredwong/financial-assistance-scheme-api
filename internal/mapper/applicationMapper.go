package mapper

import (
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
)

// what db needs
func ApplicationDTOToModel(applicationDTO dto.ApplicationDTO) models.Application {
	return models.Application{
		ApplicantId: applicationDTO.Applicant.Id,
		SchemeId: applicationDTO.Scheme.Id,
	}
}

// what client needs
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
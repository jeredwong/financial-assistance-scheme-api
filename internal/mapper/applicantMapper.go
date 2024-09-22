package mapper

import (
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
)

func ApplicantDTOToModel(applicantDTO dto.ApplicantDTO) models.Applicant{
	return models.Applicant{
		Name:             applicantDTO.Name,
        DateOfBirth:      applicantDTO.DateOfBirth,
        Sex:              models.Sex(applicantDTO.Sex),
        MaritalStatus:    models.MaritalStatus(applicantDTO.MaritalStatus),
        EmploymentStatus: models.EmploymentStatus(applicantDTO.EmploymentStatus),
	}
}

func ApplicantModelToDTO(applicant models.Applicant) dto.ApplicantDTO {
    return dto.ApplicantDTO{
        Name:             applicant.Name,
        DateOfBirth:      applicant.DateOfBirth,
        Sex:              string(applicant.Sex),
        MaritalStatus:    string(applicant.MaritalStatus),
        EmploymentStatus: string(applicant.EmploymentStatus),
        // TODO: implementation of household members mapping
    }
}

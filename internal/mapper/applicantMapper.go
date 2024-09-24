package mapper

import (
	"github.com/jeredwong/financial-scheme-manager/internal/constants"
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
)

func ApplicantDTOToModel(applicantDTO dto.ApplicantDTO) models.Applicant{
    applicant := models.Applicant{
		Name:             applicantDTO.Name,
        DateOfBirth:      applicantDTO.DateOfBirth,
        Sex:              constants.Sex(applicantDTO.Sex),
        MaritalStatus:    constants.MaritalStatus(applicantDTO.MaritalStatus),
        EmploymentStatus: constants.EmploymentStatus(applicantDTO.EmploymentStatus),
    }
    return applicant
}

func ApplicantDTOsToModels(applicantDTOs []dto.ApplicantDTO) []models.Applicant{
    applicants := make([]models.Applicant, len(applicantDTOs))
    for i, applicant := range(applicantDTOs) {
        applicants[i] = ApplicantDTOToModel(applicant)
    }
    return applicants
}

func ApplicantModelToDTO(applicant models.Applicant) dto.ApplicantDTO {
    applicantDTO := dto.ApplicantDTO {
        Name:             applicant.Name,
        DateOfBirth:      applicant.DateOfBirth,
        Sex:              string(applicant.Sex),
        MaritalStatus:    string(applicant.MaritalStatus),
        EmploymentStatus: string(applicant.EmploymentStatus),
    }
    return applicantDTO
}

func ApplicantModelsToDTOs(applicants []models.Applicant) []dto.ApplicantDTO {
    applicantDTOs := make([]dto.ApplicantDTO, len(applicants))
    for i, applicant := range(applicants) {
        applicantDTOs[i] = ApplicantModelToDTO(applicant)
    }
    return applicantDTOs
}

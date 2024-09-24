package mapper

import (
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
)

func ApplicantDTOToModel(applicantDTO dto.ApplicantDTO) models.Applicant{
    applicant := models.Applicant{
		Name:             applicantDTO.Name,
        DateOfBirth:      applicantDTO.DateOfBirth,
        Sex:              models.Sex(applicantDTO.Sex),
        MaritalStatus:    models.MaritalStatus(applicantDTO.MaritalStatus),
        EmploymentStatus: models.EmploymentStatus(applicantDTO.EmploymentStatus),
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
    // householdMembers, err := services.HouseholdMemberService.GetHouseholdMembersByApplicantId(applicant.ID)
    // // TODO: implement more graceful error handling 
    // if err != nil {
    //     log.Printf("error retrieving household members")

    // }
    // var householdMemberDTOs []dto.HouseholdMemberDTO
    // for _, householdMember := range(householdMembers) {
    //     householdMemberDTOs = append(householdMemberDTOs, mapper.HouseholdMemberModelToDTO(householdMember))
    // }
    // applicantDTO.HouseholdMembers = householdMemberDTOs
    return applicantDTO
}

func ApplicantModelsToDTOs(applicants []models.Applicant) []dto.ApplicantDTO {
    applicantDTOs := make([]dto.ApplicantDTO, len(applicants))
    for i, applicant := range(applicants) {
        applicantDTOs[i] = ApplicantModelToDTO(applicant)
    }
    return applicantDTOs
}

package mapper

import (
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
)

func HouseholdMemberDTOToModel(memberDTO dto.HouseholdMemberDTO) models.HouseholdMember {
	return models.HouseholdMember{
		// ID:               memberDTO.ID,
		// ApplicantID:      memberDTO.ApplicantID,
		Name:             memberDTO.Name,
		DateOfBirth:      memberDTO.DateOfBirth,
		Sex:              models.Sex(memberDTO.Sex),
		Relationship:     models.Relationship(memberDTO.Relationship),
		EmploymentStatus: models.EmploymentStatus(memberDTO.EmploymentStatus),
		SchoolLevel:      models.SchoolLevel(memberDTO.SchoolLevel),
		// CreatedAt:        memberDTO.CreatedAt,
		// UpdatedAt:        memberDTO.UpdatedAt,
	}
}

func HouseholdMemberDTOstoModels(memberDTOs []dto.HouseholdMemberDTO) []models.HouseholdMember {
	members := make([]models.HouseholdMember, len(memberDTOs))
	for i, memberDTO := range(memberDTOs) {
		members[i] = HouseholdMemberDTOToModel(memberDTO)
	}
	return members
}

func HouseholdMemberModelToDTO(member models.HouseholdMember) dto.HouseholdMemberDTO {
	return dto.HouseholdMemberDTO{
		// ID:               member.ID,
		// ApplicantID:      member.ApplicantID,
		Name:             member.Name,
		DateOfBirth:      member.DateOfBirth,
		Sex:              string(member.Sex),
		Relationship:     string(member.Relationship),
		EmploymentStatus: string(member.EmploymentStatus),
		SchoolLevel:      string(member.SchoolLevel),
		// CreatedAt:        member.CreatedAt,
		// UpdatedAt:        member.UpdatedAt,
	}
}

func HouseholdMemberModelsToDTOs(members []models.HouseholdMember) []dto.HouseholdMemberDTO {
	memberDTOs := make([]dto.HouseholdMemberDTO, len(members))
	for i, member := range(members) {
		memberDTOs[i] = HouseholdMemberModelToDTO(member)
	}
	return memberDTOs
}
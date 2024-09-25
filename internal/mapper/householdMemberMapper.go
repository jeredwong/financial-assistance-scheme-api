package mapper

import (
	"github.com/jeredwong/financial-scheme-manager/internal/constants"
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
)

func HouseholdMemberDTOToModel(memberDTO dto.HouseholdMemberDTO) models.HouseholdMember {
	return models.HouseholdMember{
		ID: memberDTO.Id,
		Name:             memberDTO.Name,
		DateOfBirth:      memberDTO.DateOfBirth,
		Sex:              constants.Sex(memberDTO.Sex),
		Relationship:     constants.Relationship(memberDTO.Relationship),
		EmploymentStatus: constants.EmploymentStatus(memberDTO.EmploymentStatus),
		SchoolLevel:      constants.SchoolLevel(memberDTO.SchoolLevel),
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
		Id: member.ID,
		Name:             member.Name,
		DateOfBirth:      member.DateOfBirth,
		Sex:              string(member.Sex),
		Relationship:     string(member.Relationship),
		EmploymentStatus: string(member.EmploymentStatus),
		SchoolLevel:      string(member.SchoolLevel),
	}
}

func HouseholdMemberModelsToDTOs(members []models.HouseholdMember) []dto.HouseholdMemberDTO {
	memberDTOs := make([]dto.HouseholdMemberDTO, len(members))
	for i, member := range(members) {
		memberDTOs[i] = HouseholdMemberModelToDTO(member)
	}
	return memberDTOs
}
package mapper

import (
	"github.com/jeredwong/financial-scheme-manager/internal/constants"
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
)

func SchemeCriteriaDTOToModel(criteriaDTO dto.SchemeCriteriaDTO) models.SchemeCriteria {
	return models.SchemeCriteria{
		ID: criteriaDTO.Id,
		MaritalStatus: constants.MaritalStatus(criteriaDTO.MaritalStatus),
		EmploymentStatus: constants.EmploymentStatus(criteriaDTO.EmploymentStatus),
		HouseholdStatus: criteriaDTO.HouseholdStatus,
	}
}

func SchemeCriteriaDTOsToModels(criteriaDTOs []dto.SchemeCriteriaDTO) []models.SchemeCriteria {
	criterion := make([]models.SchemeCriteria, len(criteriaDTOs)) 
	for i, criteriaDTO := range(criteriaDTOs) {
		criterion[i] = SchemeCriteriaDTOToModel(criteriaDTO)
	}
	return criterion
}

func SchemeCriteriaModelToDTO(criteria models.SchemeCriteria) dto.SchemeCriteriaDTO {
	return dto.SchemeCriteriaDTO{
		Id: criteria.ID,
		MaritalStatus: string(criteria.MaritalStatus),
		EmploymentStatus: string(criteria.MaritalStatus),
		HouseholdStatus: criteria.HouseholdStatus,
	}
}

func SchemeCriteriaModelsToDTOs(criterion []models.SchemeCriteria) []dto.SchemeCriteriaDTO {
	criteriaDTOs := make([]dto.SchemeCriteriaDTO, len(criterion))
	for i, criteria := range(criterion) {
		criteriaDTOs[i] = SchemeCriteriaModelToDTO(criteria)
	}
	return criteriaDTOs
}
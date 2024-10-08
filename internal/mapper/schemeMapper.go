package mapper

import (
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
)

func SchemeDTOToModel(schemeDTO dto.SchemeDTO) models.Scheme {
	return models.Scheme{
		ID: schemeDTO.Id,
		Name: schemeDTO.Name,
		Description: schemeDTO.Description,
	}
}

func SchemeDTOsToModels(schemeDTOs []dto.SchemeDTO) []models.Scheme {
	schemes := make([]models.Scheme, len(schemeDTOs))
	for i, schemeDTO := range(schemeDTOs) {
		schemes[i] = SchemeDTOToModel(schemeDTO)
	}
	return schemes
}

func SchemeModelToDTO(scheme models.Scheme) dto.SchemeDTO {
	schemeDTO := dto.SchemeDTO{}
	schemeDTO.Id = scheme.ID
	schemeDTO.Name = string(scheme.Name)
	schemeDTO.Description = string(scheme.Description)

	return schemeDTO
}

func SchemeModelsToDTOs(schemes []models.Scheme) []dto.SchemeDTO {
	schemeDTOs := make([]dto.SchemeDTO, len(schemes))
	for i, scheme := range(schemes) {
		schemeDTOs[i] = SchemeModelToDTO(scheme)
	}
	return schemeDTOs
}
package mapper

import (
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
)

func BenefitDTOToModel(benefitDTO dto.BenefitDTO) models.Benefit {
	return models.Benefit{
		Name: benefitDTO.Name,
		Description: benefitDTO.Description,
		Amount: benefitDTO.Amount,
	}
}

func BenefitDTOsToModels(benefitDTOs []dto.BenefitDTO) []models.Benefit {
	benefits := make([]models.Benefit, len(benefitDTOs)) 
	for i, benefitDTO := range(benefitDTOs) {
		benefits[i] = BenefitDTOToModel(benefitDTO)
	}
	return benefits
}

func BenefitModelToDTO(benefit models.Benefit) dto.BenefitDTO {
	return dto.BenefitDTO{
		Name: benefit.Name,
		Description: benefit.Description,
		Amount: benefit.Amount,
	}
}

func BenefitModelsToDTOs(benefits []models.Benefit) []dto.BenefitDTO {
	benefitDTOs := make([]dto.BenefitDTO, len(benefits))
	for i, benefit := range(benefits) {
		benefitDTOs[i] = BenefitModelToDTO(benefit)
	}
	return benefitDTOs
}
package services

import (
	"github.com/google/uuid"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"github.com/jeredwong/financial-scheme-manager/internal/repository"
)

type BenefitService interface {
	GetBenefitsBySchemeId(schemeId uuid.UUID) ([]models.Benefit, error)
	GetBenefitsByCriteriaId(criteriaId uuid.UUID) ([]models.Benefit, error)
}

type benefitService struct {
	benefitRepo repository.BenefitRepository
}

func NewBenefitService (benefitRepo repository.BenefitRepository) *benefitService {
	return &benefitService{benefitRepo: benefitRepo}
}

func (s *benefitService) GetBenefitsBySchemeId(schemeId uuid.UUID) ([]models.Benefit, error) {
	return s.benefitRepo.GetBenefitsBySchemeId(schemeId)
}

func (s *benefitService) GetBenefitsByCriteriaId(criteriaId uuid.UUID) ([]models.Benefit, error) {
	return s.benefitRepo.GetBenefitsByCriteriaId(criteriaId)
}

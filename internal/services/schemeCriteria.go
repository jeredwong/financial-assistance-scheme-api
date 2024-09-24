package services

import (
	"github.com/google/uuid"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"github.com/jeredwong/financial-scheme-manager/internal/repository"
)

type SchemeCriteriaService interface {
	GetSchemeCriteriaBySchemeId(schemeId uuid.UUID) ([]models.SchemeCriteria, error)
}

type schemeCriteriaService struct {
	schemeCriteriaRepo repository.SchemeCriteriaRepository
}

func NewSchemeCriteriaService (schemeCriteriaRepo repository.SchemeCriteriaRepository) *schemeCriteriaService {
	return &schemeCriteriaService{schemeCriteriaRepo: schemeCriteriaRepo}
}

func (s *schemeCriteriaService) GetSchemeCriteriaBySchemeId(schemeId uuid.UUID) ([]models.SchemeCriteria, error) {
	return s.schemeCriteriaRepo.GetSchemeCriteriaBySchemeId(schemeId)
}

package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"gorm.io/gorm"
)


type BenefitRepository interface {
	GetBenefitsBySchemeId(schemeId uuid.UUID) ([]models.Benefit, error)
	GetBenefitsByCriteriaId(criteriaId uuid.UUID) ([]models.Benefit, error)
}

type gormBenefitRepository struct {
	db *gorm.DB
}

func NewGormBenefitRepository(db *gorm.DB) BenefitRepository {
	return &gormBenefitRepository{db: db}
}

func (r *gormBenefitRepository) GetBenefitsBySchemeId(schemeId uuid.UUID) ([]models.Benefit, error) {
	var benefits []models.Benefit 
	result := r.db.Where("scheme_id = ?", schemeId).Find(&benefits)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("retrieved %d scheme benefits", len(benefits))
	return benefits, nil
}

func (r *gormBenefitRepository) GetBenefitsByCriteriaId(criteriaId uuid.UUID) ([]models.Benefit, error) {
	var benefits []models.Benefit 
	result := r.db.Where("criteria_id = ?", criteriaId).Find(&benefits)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("retrieved %d scheme benefits", len(benefits))
	return benefits, nil
}
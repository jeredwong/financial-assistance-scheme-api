package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"gorm.io/gorm"
)


type SchemeCriteriaRepository interface {
	GetSchemeCriteriaBySchemeId(schemeId uuid.UUID) ([]models.SchemeCriteria, error)
}

type gormSchemeCriteriaRepository struct {
	db *gorm.DB
}

func NewGormSchemeCriteriaRepository(db *gorm.DB) SchemeCriteriaRepository {
	return &gormSchemeCriteriaRepository{db: db}
}

func (r *gormSchemeCriteriaRepository) GetSchemeCriteriaBySchemeId(schemeId uuid.UUID) ([]models.SchemeCriteria, error) {
	var schemeCriterion []models.SchemeCriteria 
	result := r.db.Where("scheme_id = ?", schemeId).Find(&schemeCriterion)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("retrieved %d scheme criteria", len(schemeCriterion))
	return schemeCriterion, nil
}

package repository

import (
	"log"

	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"gorm.io/gorm"
)

type SchemeRepository interface {
	GetAllSchemes(page, pageSize int) ([]models.Scheme, int64, error) 
}

type gormSchemeRepository struct {
	db *gorm.DB
}

func NewGormSchemeRepository(db *gorm.DB) SchemeRepository {
	return &gormSchemeRepository{db: db}
}

func (r *gormSchemeRepository) GetAllSchemes(page, pageSize int) ([]models.Scheme, int64, error) {
	var schemes []models.Scheme
	var totalItems int64

	offset := (page - 1) * pageSize 

	err := r.db.Model(&models.Scheme{}) .Count(&totalItems).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset(offset).Limit(pageSize).Find(&schemes).Error
	if err != nil {
		return nil, 0, err
	}

	log.Printf("retrieved %d schemes", len(schemes))

	return schemes, totalItems, nil
}
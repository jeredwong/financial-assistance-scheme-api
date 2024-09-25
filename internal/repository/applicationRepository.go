package repository

import (
	"log"

	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"gorm.io/gorm"
)

type ApplicationRepository interface {
	GetAllApplications(page, pageSize int) ([]models.Application, int64, error)
	CreateApplication(application *models.Application) error
}

type gormApplicationRepository struct {
	db *gorm.DB
}

func NewGormApplicationRepository(db *gorm.DB) ApplicationRepository {
	return &gormApplicationRepository{db: db}
}

func (r *gormApplicationRepository) GetAllApplications(page, pageSize int) ([]models.Application, int64, error) {
	var applications []models.Application
    var totalItems int64

    offset := (page - 1) * pageSize

	// get total count 
    err := r.db.Model(&models.Application{}).Count(&totalItems).Error
    if err != nil {
        return nil, 0, err
    }

	// retrieve paginated data 
    err = r.db.Offset(offset).Limit(pageSize).Find(&applications).Error
    if err != nil {
        return nil, 0, err
    }

	log.Printf("retrieved %d applications", len(applications))

    return applications, totalItems, nil
}

func (r *gormApplicationRepository) CreateApplication(application *models.Application) error {
	return r.db.Create(application).Error
}
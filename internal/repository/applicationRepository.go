package repository

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"gorm.io/gorm"
)

type ApplicationRepository interface {
	GetAllApplications(page, pageSize int) ([]models.Application, int64, error)
	GetApplicantById(applicantId uuid.UUID) (models.Applicant, error)
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

func (r *gormApplicationRepository) GetApplicantById(applicantId uuid.UUID) (models.Applicant, error) {
	var applicant models.Applicant
	result := r.db.First(&applicant, "id = ?", applicantId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Applicant{}, errors.New("applicant not found")
		}
		return models.Applicant{}, result.Error
	}
	return applicant, nil
}

func (r *gormApplicationRepository) CreateApplication(application *models.Application) error {
	return r.db.Create(application).Error
}
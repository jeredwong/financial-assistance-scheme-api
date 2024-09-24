package repository

import (
	"log"

	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"gorm.io/gorm"
)

type ApplicantRepository interface {
	List(page, pageSize int) ([]models.Applicant, int64, error)
	Create(applicant *models.Applicant) error
}

type gormApplicantRepository struct {
	db *gorm.DB
}

func NewGormApplicantRepository(db *gorm.DB) ApplicantRepository {
	return &gormApplicantRepository{db: db}
}

func (r *gormApplicantRepository) List(page, pageSize int) ([]models.Applicant, int64, error) {
    var applicants []models.Applicant
    var totalItems int64

    offset := (page - 1) * pageSize

	// get total count 
    err := r.db.Model(&models.Applicant{}).Count(&totalItems).Error
    if err != nil {
        return nil, 0, err
    }

	// retrieve paginated data 
    err = r.db.Offset(offset).Limit(pageSize).Find(&applicants).Error
    if err != nil {
        return nil, 0, err
    }

	log.Printf("retrieved %d applicants", len(applicants))

    return applicants, totalItems, nil
}

func (r *gormApplicantRepository) Create(applicant *models.Applicant) error {
	return r.db.Create(applicant).Error
}

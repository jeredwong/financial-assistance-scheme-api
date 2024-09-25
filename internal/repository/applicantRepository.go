package repository

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"gorm.io/gorm"
)

type ApplicantRepository interface {
	GetAllApplicants(page, pageSize int) ([]models.Applicant, int64, error)
    // TODO: return pointer or copy
    GetApplicantById(applicantId uuid.UUID) (models.Applicant, error) 
	CreateApplicant(applicant *models.Applicant) error
}

type gormApplicantRepository struct {
	db *gorm.DB
}

func NewGormApplicantRepository(db *gorm.DB) ApplicantRepository {
	return &gormApplicantRepository{db: db}
}

func (r *gormApplicantRepository) GetAllApplicants(page, pageSize int) ([]models.Applicant, int64, error) {
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

func (r *gormApplicantRepository) GetApplicantById(applicantId uuid.UUID) (models.Applicant, error) {
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

func (r *gormApplicantRepository) CreateApplicant(applicant *models.Applicant) error {
	return r.db.Create(applicant).Error
}

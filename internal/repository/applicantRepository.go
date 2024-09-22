package repository

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
)

type ApplicantRepository interface {
	Create(applicant *models.Applicant) error
	GetById(id uuid.UUID) (*models.Applicant, error)
	// Update(applicant *models.Applicant) error
	// Delete(id uuid.UUID) error
	List(limit, offset int) ([]models.Applicant, error)
}

type gormApplicantRepository struct {
	db *gorm.DB
}

func NewGormApplicantRepository(db *gorm.DB) ApplicantRepository {
	return &gormApplicantRepository{db: db}
}

func (r *gormApplicantRepository) Create(applicant *models.Applicant) error {
	return r.db.Create(applicant).Error
}

func (r *gormApplicantRepository) GetById(id uuid.UUID) (*models.Applicant, error) {
	var applicant models.Applicant
	err := r.db.Preload("HouseholdMembers").First(&applicant, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("applicant not found")
		}
		return nil, err
	}
	return &applicant, nil
}

func (r *gormApplicantRepository) List(limit, offset int) ([]models.Applicant, error) {
	var applicants []models.Applicant
	err := r.db.Limit(limit).Offset(offset).Find(&applicants).Error
	return applicants, err
}

// func (r *gormApplicantRepository) Update(ap)


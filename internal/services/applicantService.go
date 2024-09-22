package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"github.com/jeredwong/financial-scheme-manager/internal/repository"
)

type ApplicantService interface {
	CreateApplicant(applicant *models.Applicant) error
	GetApplicantById(id uuid.UUID) (*models.Applicant, error)
	ListApplicants(page, pageSize int) ([]models.Applicant, error)
}

type applicantService struct {
	applicantRepo repository.ApplicantRepository
}

func NewApplicantService(applicantRepo repository.ApplicantRepository) ApplicantService {
	return &applicantService{applicantRepo: applicantRepo}
}

func (s *applicantService) CreateApplicant(applicant *models.Applicant) error {
	if err := validateApplicant(applicant); err != nil {
		return err
	}
	return s.applicantRepo.Create(applicant)
}

func (s *applicantService) GetApplicantById(id uuid.UUID) (*models.Applicant, error) {
	return  s.applicantRepo.GetById(id)
}

func (s *applicantService) ListApplicants(page, pageSize int) ([]models.Applicant, error) {
	offset := (page - 1) * pageSize
	return s.applicantRepo.List(pageSize, offset)
}

// helper functions
func validateApplicant(applicant *models.Applicant) error {
	if applicant.Name == "" {
		return errors.New("applicant name is required")
	}
	if applicant.DateOfBirth.After(time.Now()) {
		return errors.New("date of birth cannot be in the future ")
	}
	// TODO: more validation
	return nil
}

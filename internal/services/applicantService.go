package services

import (
	"errors"
	"math"
	"time"

	"github.com/jeredwong/financial-scheme-manager/internal/constants"
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"github.com/jeredwong/financial-scheme-manager/internal/repository"
)

type ApplicantService interface {
	ListApplicants(query dto.PaginationQuery) (dto.PaginatedResponse, error)
	CreateApplicant(applicant *models.Applicant) error
}

type applicantService struct {
	applicantRepo repository.ApplicantRepository
}

func NewApplicantService(applicantRepo repository.ApplicantRepository) ApplicantService {
	return &applicantService{applicantRepo: applicantRepo}
}

func (s *applicantService) ListApplicants(query dto.PaginationQuery) (dto.PaginatedResponse, error) {
	// default values
	if query.Page == 0 {
		query.Page = constants.DefaultPage
	}
	if query.PageSize == 0 {
		query.PageSize = constants.DefaultPageSize
	}
	if query.PageSize > constants.MaxPageSize {
		query.PageSize = constants.MaxPageSize
	}

	applicants, totalItems, err := s.applicantRepo.List(query.Page, query.PageSize)
	if err != nil {
		return dto.PaginatedResponse{}, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(query.PageSize)))

	return dto.PaginatedResponse{
		Data:       applicants,
        TotalItems: totalItems,
        TotalPages: totalPages,
        Page:       query.Page,
        PageSize:   query.PageSize,
    }, nil
}

func (s *applicantService) CreateApplicant(applicant *models.Applicant) error {
	if err := validateApplicant(applicant); err != nil {
		return err
	}
	return s.applicantRepo.Create(applicant)
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

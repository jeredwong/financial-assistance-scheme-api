package services

import (
	"math"

	"github.com/jeredwong/financial-scheme-manager/internal/constants"
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"github.com/jeredwong/financial-scheme-manager/internal/repository"
)

type ApplicationService interface {
	GetAllApplications(query dto.PaginationQuery) (dto.PaginatedResponse, error)
	CreateApplication(applicant *models.Application) error
}

type applicationService struct {
	applicationRepo repository.ApplicationRepository
}

func NewApplicationService(applicationRepo repository.ApplicationRepository) ApplicationService {
	return &applicationService{applicationRepo: applicationRepo}
}

func (s *applicationService) GetAllApplications(query dto.PaginationQuery) (dto.PaginatedResponse, error) {
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

	applications, totalItems, err := s.applicationRepo.GetAllApplications(query.Page, query.PageSize)
	if err != nil {
		return dto.PaginatedResponse{}, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(query.PageSize)))

	return dto.PaginatedResponse{
		Data:       applications,
        TotalItems: totalItems,
        TotalPages: totalPages,
        Page:       query.Page,
        PageSize:   query.PageSize,
    }, nil
}

func (s *applicationService) CreateApplication(application *models.Application) error {
	if err := validateApplication(application); err != nil {
		return err
	}
	return s.applicationRepo.CreateApplication(application)
}

// helper functions
func validateApplication(application *models.Application) error {
	// TODO: application validation 
	return nil
}

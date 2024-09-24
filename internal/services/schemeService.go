package services

import (
	"math"

	"github.com/jeredwong/financial-scheme-manager/internal/constants"
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/repository"
)

type SchemeService interface {
	GetAllSchemes(query dto.PaginationQuery) (dto.PaginatedResponse, error)
}

type schemeService struct {
	schemeRepo repository.SchemeRepository
}

func NewSchemeService(schemeRepo repository.SchemeRepository) SchemeService{
	return &schemeService{schemeRepo: schemeRepo}
}

func (s *schemeService) GetAllSchemes(query dto.PaginationQuery) (dto.PaginatedResponse, error) {
	if query.Page == 0 {
		query.Page = constants.DefaultPage
	}
	if query.PageSize == 0 {
		query.PageSize = constants.DefaultPageSize
	}
	if query.PageSize > constants.MaxPageSize {
		query.PageSize = constants.MaxPageSize
	}

	schemes, totalItems, err := s.schemeRepo.GetAllSchemes(query.Page, query.PageSize)
	if err != nil {
		return dto.PaginatedResponse{}, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(query.PageSize)))

	return dto.PaginatedResponse{
		Data: 		schemes,
		TotalItems:	totalItems,
		TotalPages:	totalPages,	
		Page:		query.Page,
		PageSize: 	query.PageSize,
	}, nil
}
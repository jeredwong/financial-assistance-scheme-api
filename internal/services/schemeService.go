package services

import (
	"fmt"
	"log"
	"math"

	"github.com/google/uuid"
	"github.com/jeredwong/financial-scheme-manager/internal/constants"
	"github.com/jeredwong/financial-scheme-manager/internal/dto"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"github.com/jeredwong/financial-scheme-manager/internal/repository"
)

type SchemeService interface {
	GetAllSchemes(query dto.PaginationQuery) (dto.PaginatedResponse, error)
	GetSchemeById(schemeId uuid.UUID) (models.Scheme, error)
	GetEligibleSchemesForApplicant(applicantId uuid.UUID) ([]models.Scheme, error)
}

type schemeService struct {
	schemeRepo repository.SchemeRepository
	applicantRepo repository.ApplicantRepository
	householdMemberRepo repository.HouseholdMemberRepository
}

func NewSchemeService(
	schemeRepo repository.SchemeRepository,
	applicantRepo repository.ApplicantRepository,
	householdMemberRepo repository.HouseholdMemberRepository) SchemeService{
	return &schemeService{
		schemeRepo: schemeRepo,
		applicantRepo: applicantRepo,
		householdMemberRepo: householdMemberRepo,
	}
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

func (s *schemeService) GetSchemeById(schemeId uuid.UUID) (models.Scheme, error) {
	return s.schemeRepo.GetSchemeById(schemeId)
}

func (s *schemeService) GetEligibleSchemesForApplicant(applicantId uuid.UUID) ([]models.Scheme, error)  {
	applicant, err := s.applicantRepo.GetApplicantById(applicantId)
	if err != nil {
		return nil, fmt.Errorf("error fetching applicant: %w", err)
	}
	log.Printf("getting eligible schemes for applicant: %s", applicantId)

	householdMembers, err := s.householdMemberRepo.GetByApplicantId(applicantId)
	if err != nil {
		return nil, fmt.Errorf("error fetching household members: %w", err)
	}

	eligibleSchemes, err := s.schemeRepo.GetEligibleSchemesForApplicant(applicant, householdMembers)
	if err != nil {
		return nil, err
	}

	return eligibleSchemes, nil
}
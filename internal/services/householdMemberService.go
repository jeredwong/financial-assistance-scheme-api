package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"github.com/jeredwong/financial-scheme-manager/internal/repository"
)

type HouseholdMemberService interface {
	CreateHouseholdMember(householdMember *models.HouseholdMember) error
	GetHouseholdMembersByApplicantId(applicantId uuid.UUID) ([]models.HouseholdMember, error)
}

type householdMemberService struct {
	householdMemberRepo repository.HouseholdMemberRepository
}

func NewHouseholdMemberService(householdMemberRepo repository.HouseholdMemberRepository) HouseholdMemberService {
	return &householdMemberService{householdMemberRepo: householdMemberRepo}
}

func (s *householdMemberService) CreateHouseholdMember(householdMember *models.HouseholdMember) error {
	if err := validateHouseholdMember(householdMember); err != nil {
		return err
	}
	return s.householdMemberRepo.Create(householdMember)
}

func (s *householdMemberService) GetHouseholdMembersByApplicantId(applicantId uuid.UUID) ([]models.HouseholdMember, error) {
	return s.householdMemberRepo.GetByApplicantId(applicantId)
}

// helper function
func validateHouseholdMember(householdMember *models.HouseholdMember) error {
	if householdMember.Name == "" {
		return errors.New("applicant name is required")
	}
	if householdMember.DateOfBirth.After(time.Now()) {
		return errors.New("date of birth cannot be in the future ")
	}
	// TODO: more validation
	return nil
}
package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"gorm.io/gorm"
)

type SchemeRepository interface {
	GetAllSchemes(page, pageSize int) ([]models.Scheme, int64, error) 
	GetSchemeById(schemeId uuid.UUID) (models.Scheme, error) 
	GetEligibleSchemesForApplicant(applicant models.Applicant, householdMember []models.HouseholdMember) ([]models.Scheme, error)
}

type gormSchemeRepository struct {
	db *gorm.DB
}

func NewGormSchemeRepository(db *gorm.DB) SchemeRepository {
	return &gormSchemeRepository{db: db}
}

func (r *gormSchemeRepository) GetAllSchemes(page, pageSize int) ([]models.Scheme, int64, error) {
	var schemes []models.Scheme
	var totalItems int64

	offset := (page - 1) * pageSize 

	err := r.db.Model(&models.Scheme{}) .Count(&totalItems).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset(offset).Limit(pageSize).Find(&schemes).Error
	if err != nil {
		return nil, 0, err
	}

	log.Printf("retrieved %d schemes", len(schemes))

	return schemes, totalItems, nil
}

func (r *gormSchemeRepository) GetSchemeById(schemeId uuid.UUID) (models.Scheme, error) {

    var scheme models.Scheme
    result := r.db.First(&scheme, "id = ?", schemeId)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return models.Scheme{}, errors.New("scheme not found")
        }
        return models.Scheme{}, result.Error
    }
    return scheme, nil
}

func (r *gormSchemeRepository) GetEligibleSchemesForApplicant(applicant models.Applicant, householdMembers []models.HouseholdMember) ([]models.Scheme, error) {
	var schemes []models.Scheme
    
    // Fetch all schemes with their criteria
    err := r.db.Preload("Criteria").Find(&schemes).Error
    if err != nil {
        return nil, fmt.Errorf("error fetching schemes: %w", err)
    }

    var eligibleSchemes []models.Scheme
    for _, scheme := range schemes {
        if isSchemeEligible(scheme, applicant, householdMembers) {
            eligibleSchemes = append(eligibleSchemes, scheme)
        }
    }

    return eligibleSchemes, nil
}

// helper functions
func isSchemeEligible(scheme models.Scheme, applicant models.Applicant, householdMembers []models.HouseholdMember) bool {
    for _, criteria := range scheme.Criteria {
        if isCriteriamet(criteria, applicant, householdMembers) {
            return true
        }
    }
    return false
}

func isCriteriamet(criteria models.SchemeCriteria, applicant models.Applicant, householdMembers []models.HouseholdMember) bool {
    // Check marital status
    if criteria.MaritalStatus != "" && criteria.MaritalStatus != applicant.MaritalStatus {
        return false
    }

    // Check employment status
    if criteria.EmploymentStatus != "" && criteria.EmploymentStatus != applicant.EmploymentStatus {
        return false
    }

    // Check household status
    if criteria.HouseholdStatus != nil {
        var householdCriteria map[string]string
        if err := json.Unmarshal(criteria.HouseholdStatus, &householdCriteria); err != nil {
            return false
        }
        if !isHouseholdEligible(householdMembers, householdCriteria) {
            return false
        }
    }

    return true
}

func isHouseholdEligible(householdMembers []models.HouseholdMember, criteria map[string]string) bool {
    for key, value := range criteria {
        eligible := false
        for _, member := range householdMembers {
            switch key {
            case "schoolLevel":
                if string(member.SchoolLevel) == value {
                    eligible = true
                    break
                }
            // Add more cases here for other household status criteria
            }
            if eligible {
                break
            }
        }
        if !eligible {
            return false
        }
    }
    return true
}
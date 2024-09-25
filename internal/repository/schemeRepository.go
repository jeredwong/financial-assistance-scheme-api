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

func (r *gormSchemeRepository) GetEligibleSchemesForApplicant (applicant models.Applicant, householdMembers []models.HouseholdMember) ([]models.Scheme, error) {
	var eligibleSchemes []models.Scheme

	query := r.db.Distinct("schemes.*").
		Joins("JOIN scheme_criteria ON schemes.id = scheme_criteria.scheme_id").
		Where("(scheme_criteria.marital_status IS NULL OR scheme_criteria.marital_status = ?) AND " +
		"(scheme_criteria.employment_status IS NULL OR scheme_criteria.employement_status = ?)",
		applicant.MaritalStatus, applicant.EmploymentStatus)

	err := query.Find(&eligibleSchemes).Error
	if err != nil {
		return nil, fmt.Errorf("error fetching eligible schemes: %w", err)
	}

	var finalEligibleSchemes []models.Scheme
	for _, scheme := range(eligibleSchemes) {
		var criteria models.SchemeCriteria
		if err := r.db.Where("scheme_id = ?", scheme.ID).First(&criteria).Error; err != nil {
			continue
		}

		if isEligibleBasedOnHouseholdStatus(householdMembers, criteria.HouseholdStatus) {
			finalEligibleSchemes = append(finalEligibleSchemes, scheme)
		}
	}

	return finalEligibleSchemes, nil
}

// helper functions 
func isEligibleBasedOnHouseholdStatus(householdMembers []models.HouseholdMember, householdStatus json.RawMessage) bool {
	var criteria map[string]string
	if err := json.Unmarshal(householdStatus, &criteria); err != nil {
		log.Printf("Failed to read household status")
		return false
	}
	
	// key: criteria type; value: criteria value 
	for key, value := range(criteria) {
		eligible := false
		// for each criteria, check if ANY member meets it 
		for _, member := range(householdMembers) {
			eligible = true
			switch key {
			case "schoolLevel":
				if string(member.SchoolLevel) != value {
					eligible = false
					continue
				}
			}
			if eligible {
				break
			}
		}
		if !eligible {
			return false // if any criterion is not met, the scheme is not eligible
		}
	}
	return true
}
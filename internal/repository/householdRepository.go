package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/jeredwong/financial-scheme-manager/internal/models"
	"gorm.io/gorm"
)

type HouseholdMemberRepository interface {
	Create(householdMember *models.HouseholdMember) error
	GetByApplicantId(applicantId uuid.UUID) ([]models.HouseholdMember, error)
}

type gormHouseholdMemberRepository struct {
	db *gorm.DB
}

func NewGormHouseholdMemberRepository(db *gorm.DB) HouseholdMemberRepository {
	return &gormHouseholdMemberRepository{db: db}
}

func (r *gormHouseholdMemberRepository) Create(householdMember *models.HouseholdMember) error {
	return r.db.Create(householdMember).Error
}

func (r *gormHouseholdMemberRepository) GetByApplicantId(applicantId uuid.UUID) ([]models.HouseholdMember, error) {
    var householdMembers []models.HouseholdMember
	result := r.db.Where("applicant_id = ?", applicantId).Find(&householdMembers)
    if result.Error != nil {
        return nil, result.Error
    }
	log.Printf("retrieved %d household members", len(householdMembers))
    return householdMembers, nil
}
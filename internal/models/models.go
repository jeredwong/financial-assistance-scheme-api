package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jeredwong/financial-scheme-manager/internal/constants"
)

type Applicant struct {
	// gorm.Model
	ID               uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name             string    `gorm:"type:varchar(255);not null"`
	DateOfBirth      time.Time `gorm:"type:date;not null"`
	Sex              constants.Sex       `gorm:"type:sex;not null"`
	MaritalStatus    constants.MaritalStatus    `gorm:"type:marital_status;not null"`
	EmploymentStatus constants.EmploymentStatus `gorm:"type:employment_status;not null"`
	CreatedAt        time.Time        `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time        `gorm:"default:CURRENT_TIMESTAMP"`
	// HouseholdMembers []HouseholdMember `gorm:"foreignKey:ApplicantID"`
	// Applications		[]Application
}

type HouseholdMember struct {
	// gorm.Model
	ID               uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ApplicantID      uuid.UUID `gorm:"type:uuid;not null"`
	Name             string    `gorm:"type:varchar(255);not null"`
	DateOfBirth      time.Time `gorm:"type:date;not null"`
	Sex              constants.Sex       `gorm:"type:sex;not null"`
	Relationship     constants.Relationship     `gorm:"type:relationship;not null"`
	EmploymentStatus constants.EmploymentStatus `gorm:"type:employment_status;not null"`
	SchoolLevel      constants.SchoolLevel      `gorm:"type:school_level;not null"`
	CreatedAt        time.Time        `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time        `gorm:"default:CURRENT_TIMESTAMP"`

	// Applicant Applicant `gorm:"foreignKey:ApplicantID"`
}

type Scheme struct {
	ID			uuid.UUID	`gorm:"type:uuid;primary_key;default;uuid_generate_v4()"`
	Name		string		`gorm:"type:varchar(255);not null"`
	Description	string		`gorm:"type:text"`
	CreatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP"` 
}

type SchemeCriteria struct {
	ID 					uuid.UUID 			`gorm:"type:uuid;primary_key;default;uuid_generate_v4()"`
	SchemeID			uuid.UUID			`gorm:"type:uuid;not null"`
	MaritalStatus		constants.MaritalStatus		`gorm:"type:marital_status;not null"`
	EmploymentStatus	constants.EmploymentStatus	`gorm:"type:employment_status;not null"`
	HouseholdStatus		json.RawMessage		`gorm:"type:json"`
	// CriteriaType		string				`gorm:"type:varchar(50);not null"`
	// CriteriaValue		json.RawMessage		`gorm:"type:json;not null"`
	CreatedAt			time.Time			`gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt			time.Time			`gorm:"default:CURRENT_TIMESTAMP"`
}

type Benefit struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	SchemeID    uuid.UUID `gorm:"type:uuid;not null"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Amount      float64   `gorm:"type:decimal(10,2)"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
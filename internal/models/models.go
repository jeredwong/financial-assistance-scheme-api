package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jeredwong/financial-scheme-manager/internal/constants"
)

type Applicant struct {
	// gorm.Model
	ID               uuid.UUID 					`gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name             string    					`gorm:"type:varchar(255);not null"`
	DateOfBirth      time.Time 					`gorm:"type:date;not null"`
	Sex              constants.Sex       		`gorm:"type:sex;not null"`
	MaritalStatus    constants.MaritalStatus    `gorm:"type:marital_status;not null"`
	EmploymentStatus constants.EmploymentStatus `gorm:"type:employment_status;not null"`
	CreatedAt        time.Time        			`gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time        			`gorm:"default:CURRENT_TIMESTAMP"`
	// HouseholdMembers []HouseholdMember `gorm:"foreignKey:ApplicantID"`
	// Applications		[]Application
}

type HouseholdMember struct {
	// gorm.Model
	ID               uuid.UUID 					`gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ApplicantID      uuid.UUID 					`gorm:"type:uuid;not null"`
	Name             string    					`gorm:"type:varchar(255);not null"`
	DateOfBirth      time.Time 					`gorm:"type:date;not null"`
	Sex              constants.Sex       		`gorm:"type:sex;not null"`
	Relationship     constants.Relationship     `gorm:"type:relationship;not null"`
	EmploymentStatus constants.EmploymentStatus `gorm:"type:employment_status;not null"`
	SchoolLevel      constants.SchoolLevel      `gorm:"type:school_level;not null"`
	CreatedAt        time.Time        			`gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time        			`gorm:"default:CURRENT_TIMESTAMP"`

	// Applicant Applicant `gorm:"foreignKey:ApplicantID"`
}

type Scheme struct {
	ID			uuid.UUID	`gorm:"type:uuid;primary_key;default;uuid_generate_v4()"`
	Name		string		`gorm:"type:varchar(255);not null"`
	Description	string		`gorm:"type:text"`
	CreatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP"` 
	Criteria []SchemeCriteria `gorm:"foreignKey:SchemeID"`
}

type SchemeCriteria struct {
	ID 					uuid.UUID 					`gorm:"type:uuid;primary_key;default;uuid_generate_v4()"`
	SchemeID			uuid.UUID					`gorm:"type:uuid;not null"`
	MaritalStatus		constants.MaritalStatus		`gorm:"type:marital_status;not null"`
	EmploymentStatus	constants.EmploymentStatus	`gorm:"type:employment_status;not null"`
	HouseholdStatus		json.RawMessage				`gorm:"type:json"`
	CreatedAt			time.Time					`gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt			time.Time					`gorm:"default:CURRENT_TIMESTAMP"`
	Benefits []Benefit `gorm:"foreignKey:CriteriaID"`
}

type Benefit struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	SchemeID    uuid.UUID `gorm:"type:uuid;not null"`
	CriteriaID	uuid.UUID `gorm:"type:uuid;not null"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Amount      float64   `gorm:"type:decimal(10,2)"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Application struct {
	ID 				uuid.UUID					`gorm:"type:uuid;primary_key;default:uuid_gernerate_v4()"`
	ApplicantId 	uuid.UUID					`gorm:"type:uuid; not null"`
	SchemeId 		uuid.UUID					`gorm:"type:uuid;not null"`
	Status 			constants.ApplicationStatus	`gorm:"type:application_status;default:pending;not null"`
	ApplicationDate time.Time 					`gorm:"default:CURRENT_TIMESTAMP;not null"`
	CreatedAt   	time.Time 					`gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   	time.Time 					`gorm:"default:CURRENT_TIMESTAMP"`
}
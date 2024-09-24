package models

import (
	"time"

	"github.com/google/uuid"
	// "gorm.io/gorm"
)

type EmploymentStatus string
type MaritalStatus string
type Sex string
type Relationship string
type SchoolLevel string
type ApplicationStatus string

const (
	EmploymentStatusEmployed	EmploymentStatus = "employed" 
	EmploymentStatusUnemployed	EmploymentStatus = "unemployed"

	MaritalStatusSingle		MaritalStatus = "single"
	MaritalStatusMarried	MaritalStatus = "married"
	MaritalStatusWidowed	MaritalStatus = "widowed"
	MaritalStatusDivorced	MaritalStatus = "divorced"

	SexMale		Sex = "male" 
	SexFemale	Sex = "female"
	SexOther	Sex = "other"

	RelationshipSpouse  Relationship = "spouse"
	RelationshipChild   Relationship = "child"
	RelationshipParent  Relationship = "parent"
	RelationshipSibling Relationship = "sibling"
	RelationshipOther   Relationship = "other"

	SchoolLevelPreschool SchoolLevel = "preschool"
	SchoolLevelPrimary   SchoolLevel = "primary"
	SchoolLevelSecondary SchoolLevel = "secondary"
	SchoolLevelTertiary  SchoolLevel = "tertiary"
	SchoolLevelNone      SchoolLevel = "none"

	ApplicationStatusPending  ApplicationStatus = "pending"
	ApplicationStatusApproved ApplicationStatus = "approved"
	ApplicationStatusRejected ApplicationStatus = "rejected"
)

type Applicant struct {
	// gorm.Model
	ID               uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name             string    `gorm:"type:varchar(255);not null"`
	DateOfBirth      time.Time `gorm:"type:date;not null"`
	Sex              Sex       `gorm:"type:sex;not null"`
	MaritalStatus    MaritalStatus    `gorm:"type:marital_status;not null"`
	EmploymentStatus EmploymentStatus `gorm:"type:employment_status;not null"`
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
	Sex              Sex       `gorm:"type:sex;not null"`
	Relationship     Relationship     `gorm:"type:relationship;not null"`
	EmploymentStatus EmploymentStatus `gorm:"type:employment_status;not null"`
	SchoolLevel      SchoolLevel      `gorm:"type:school_level;not null"`
	CreatedAt        time.Time        `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time        `gorm:"default:CURRENT_TIMESTAMP"`

	// Applicant Applicant `gorm:"foreignKey:ApplicantID"`
}
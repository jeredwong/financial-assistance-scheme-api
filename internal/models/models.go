package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
	gorm.Model
	ID					uuid.UUID
	Name				string
	DateOfBirth			time.Time
	Sex					Sex
	MaritalStatus		MaritalStatus
	EmploymentStatus	EmploymentStatus
	CreatedAt			time.Time
	UpdatedAt			time.Time
	HouseholdMembers	[]HouseholdMember
	// Applications		[]Application
}

type HouseholdMember struct {
	ID					uuid.UUID
	ApplicantID			uuid.UUID
	Name				string
	DateOfBirth			time.Time
	Sex					Sex
	Relationship		Relationship
	EmploymentStatus	EmploymentStatus
	SchoolLevel			SchoolLevel
	CreatedAt			time.Time
	UpdatedAt			time.Time
	Applicant			Applicant
}
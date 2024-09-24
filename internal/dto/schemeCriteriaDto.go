package dto

import (
	"encoding/json"
)

type SchemeCriteriaDTO struct {
	MaritalStatus		string			`gorm:"type:marital_status;not null"`
	EmploymentStatus	string			`gorm:"type:employment_status;not null"`
	HouseholdStatus		json.RawMessage	`gorm:"type:json"`
	// CriteriaType	string			`json:"criteria_type"`
	// CriteriaValue	json.RawMessage	`json:"criteria_value"`
}
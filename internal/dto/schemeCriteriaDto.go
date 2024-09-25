package dto

import (
	"encoding/json"

	"github.com/google/uuid"
)

type SchemeCriteriaDTO struct {
	Id					uuid.UUID			`json:"id"`
	MaritalStatus		string				`json:"marital_status"`
	EmploymentStatus	string				`json:"employment_status"`
	HouseholdStatus		json.RawMessage		`json:"household_status"`
	Benefits			[]BenefitDTO		`json:"benefits"`
}
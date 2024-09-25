package dto

import "github.com/google/uuid"

type SchemeDTO struct {
	Id			uuid.UUID			`json:"id"`
	Name		string				`json:"name"`
	Description	string				`json:"description"`
	Criteria	[]SchemeCriteriaDTO	`json:"criteria"`
}
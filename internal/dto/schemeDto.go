package dto

type SchemeDTO struct {
	Name		string				`json:"name"`
	Description	string				`json:"description"`
	Criteria	[]SchemeCriteriaDTO	`json:"criteria"`
	Benefits	[]BenefitDTO		`json:"benefits"`
}
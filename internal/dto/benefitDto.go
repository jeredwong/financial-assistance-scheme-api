package dto

import "github.com/google/uuid"

type BenefitDTO struct {
	Id			uuid.UUID			`json:"id"`
	Name		string	`gorm:"type:varchar(255);not null"`
	Description	string	`gorm:"type:text"`
	Amount		float64	`gorm:"type:numeric(10,2)"`
}
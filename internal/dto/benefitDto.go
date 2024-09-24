package dto

type BenefitDTO struct {
	Name		string	`gorm:"type:varchar(255);not null"`
	Description	string	`gorm:"type:text"`
	Amount		float64	`gorm:"type:numeric(10,2)"`
}
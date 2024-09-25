package dto

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationDTO struct {
	Id				uuid.UUID					`json:"id"`
	// ApplicantId		uuid.UUID					`json:"applicant_id"`
	// SchemeId		uuid.UUID					`json:"scheme_id"`
	Status			string						`json:"status"`
	ApplicationDate	time.Time					`json:"application_date"`
	// TODO: can come up with a less detailed applicantDTO and schemeDTO (eg. simpleApplicantDTO, simpleSchemeDTO)
	Applicant		ApplicantDTO				`json:"applicant"`
	Scheme			SchemeDTO					`json:"scheme"`
}
package dto

import (
	"time"

	"github.com/google/uuid"
)

type ApplicantDTO struct {
	Id					uuid.UUID				`json:"id"`
	Name 				string					`json:"name"`
	DateOfBirth 		time.Time				`json:"date_of_birth"`
	Sex 				string 					`json:"sex"`
	MaritalStatus		string 					`json:"marital_status"`
	EmploymentStatus	string 					`json:"employment_status"`
	HouseholdMembers	[]HouseholdMemberDTO 	`json:"household_members"`
}
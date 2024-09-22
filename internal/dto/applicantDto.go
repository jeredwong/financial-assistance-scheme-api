package dto

import "time"

type ApplicantDTO struct {
	Name 				string					`json:"name"`
	DateOfBirth 		time.Time				`json:"date_of_birth"`
	Sex 				string 					`json:"sex"`
	MaritalStatus		string 					`json:"marital_status"`
	EmploymentStatus	string 					`json:"employment_status"`
	HouseholdMembers	[]HouseholdMemberDTO 	`json:"household_members`
}

type HouseholdMemberDTO struct {
	Name 				string					`json:"name"`
	DateOfBirth 		time.Time				`json:"date_of_birth"`
	Sex 				string 					`json:"sex"`
	Relationship		string					`json:"relationship"`
	EmploymentStatus	string 					`json:"employment_status"`
	SchoolLevel			string 					`json:"school_level"`

}
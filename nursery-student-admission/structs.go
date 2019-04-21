package main

const (
	STUDENT_APPLIED                         int = 0
	APPLICATION_REJECTED_BY_SCHOOL          int = 1
	APPLICATION_ACCEPTED_BY_SCHOOL          int = 2
	APPLICATION_REJECTED_BY_EDUCATION_BOARD int = 3
	STUDENT_ADMITTED                        int = 4
)

type StudentDetails struct {
	StudentId            string `json:"studentId"`
	Name                 string `json:"name"`
	DOB                  string `json:"dob"`
	Gender               string `json:"gender"`
	FatherName           string `json:"fatherName"`
	MotherName           string `json:"motherName"`
	Address              string `json:"address"`
	City                 string `json:"city"`
	Status               int    `json:"status"`
	SchoolID             string `json:"schoolId"`
	SchoolRemark         string `json:"schoolRemark"`
	EducationBoardID     string `json:"educationBoardId"`
	EducationBoardRemark string `json:"educationBoardRemark"`
}

type SchoolDetails struct {
	SchoolId         string `json:"schoolId"`
	Name             string `json:"name"`
	PrincipalName    string `json:"principalName"`
	Address          string `json:"address"`
	EducationBoardID string `json:"educationBoardID"`
}

type EducationBoardDetails struct {
	EducationBoardId string `json:"educationBoardId"`
	Name             string `json:"name"`
	HeadOfficer      string `json:"headOfficer"`
	Address          string `json:"address"`
}

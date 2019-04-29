package main

import "time"

type Gender string

type StudentStatus int

type StudentTransferApplicationType int

type StudentTransferApplicationStatus int

type AssetType string

const (
	STUDENT                      AssetType = "Student"
	STUDENT_TRANSFER_APPLICATION AssetType = "StudentTransferApplication"
)

const (
	MALE        Gender = "MALE"
	FEMALE      Gender = "FEMALE"
	TRANSGENDER Gender = "TRANSGENDER"
)

const (
	ADMISSION_PENDING  StudentStatus = -1
	MEMBER             StudentStatus = 0
	TRANSFERRED        StudentStatus = 1
	EDUCATION_COMPLETE StudentStatus = 2
)

const (
	LEAVE_REQUEST     StudentTransferApplicationType = 0
	ADMISSION_REQUEST StudentTransferApplicationType = 1
)

const (
	PENDING                    StudentTransferApplicationStatus = 0
	REJECTED_BY_SCHOOL         StudentTransferApplicationStatus = 1
	REJECTED_BY_CBSE           StudentTransferApplicationStatus = 2
	LEAVE_REQUEST_ACCEPTED     StudentTransferApplicationStatus = 3
	ADMISSION_REQUEST_ACCEPTED StudentTransferApplicationStatus = 4
)

type Student struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	DOB           string        `json:"dob"`
	Gender        Gender        `json:"gender"`
	FatherName    string        `json:"fatherName"`
	MotherName    string        `json:"motherName"`
	Status        StudentStatus `json:"status"`
	DateOfJoining time.Time     `json:"dateOfJoining"`
	DateOfLeaving time.Time     `json:"dateOfLeaving"`
	Doctype       AssetType     `json:"doctype"`
}

type StudentTransferApplication struct {
	ID                string                           `json:"id"`
	StudentID         string                           `json:"studentId"`
	DateOfRequest     time.Time                        `json:"dateOfRequest"`
	LastModified      time.Time                        `json:"lastModified"`
	CBSERemark        string                           `json:"cbseRemark"`
	SchoolRemark      string                           `json:"schoolRemark"`
	ApplicationType   StudentTransferApplicationType   `json:"applicationType"`
	ApplicationStatus StudentTransferApplicationStatus `json:"applicationStatus"`
	Doctype           AssetType                        `json:"doctype"`
}

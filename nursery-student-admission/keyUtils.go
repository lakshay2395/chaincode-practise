package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func GetStudentKey(stub shim.ChaincodeStubInterface, studentId string) (string, error) {
	key, err := stub.CreateCompositeKey("Student", []string{studentId})
	return key, err
}

func GetSchoolKey(stub shim.ChaincodeStubInterface, schoolId string) (string, error) {
	key, err := stub.CreateCompositeKey("School", []string{schoolId})
	return key, err
}

func GetEducationBoardKey(stub shim.ChaincodeStubInterface, educationBoardId string) (string, error) {
	key, err := stub.CreateCompositeKey("EducationBoard", []string{educationBoardId})
	return key, err
}

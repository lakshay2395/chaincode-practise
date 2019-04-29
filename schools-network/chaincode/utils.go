package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func ValidateArgumentCount(expected int, found int) error {
	if expected != found {
		return errors.New(fmt.Sprintf("Argument count does not match. Expected %d , found %d ", expected, found))
	}
	return nil
}

func GetStudentKey(stub shim.ChaincodeStubInterface, studentId string) (string, error) {
	key, err := stub.CreateCompositeKey("Student", []string{studentId})
	return key, err
}

func GetStudentTransferApplicationKey(stub shim.ChaincodeStubInterface, studentTransferApplicationId string) (string, error) {
	key, err := stub.CreateCompositeKey("StudentTransferApplication", []string{studentTransferApplicationId})
	return key, err
}

func GetState(stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	data, err := stub.GetState(id)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("No record exists for given id")
	}
	return data, nil
}

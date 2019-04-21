package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var CC, STUB = getInitializedChaincodeAndStub()

func getInitializedChaincodeAndStub() (*NurseryStudentAdmissionChaincode, *shim.MockStub) {
	hwcc := new(NurseryStudentAdmissionChaincode)
	stub := shim.NewMockStub("MockStub", hwcc)
	stub.MockInit("1", nil)
	return hwcc, stub
}

func TestInit(t *testing.T) {
	hwcc := new(NurseryStudentAdmissionChaincode)
	stub := shim.NewMockStub("MockStub", hwcc)
	stub.MockInit("1", nil)
}

func TestApply(t *testing.T) {
	studentKey, err := GetStudentKey(STUB, "001")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	schoolKey, err := GetSchoolKey(STUB, "001")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	educationBoardKey, err := GetEducationBoardKey(STUB, "001")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	response := STUB.MockInvoke("2", [][]byte{
		[]byte("apply"),
		[]byte("001"),
		[]byte("Lakshay Bhambri"),
		[]byte("23-06-1995"),
		[]byte("Gender"),
		[]byte("Vinay Bhambri"),
		[]byte("Vimmi Bhambri"),
		[]byte("1/550 G.T. Road MS Park Shahdara Delhi - 110032"),
		[]byte("Delhi"),
		[]byte(schoolKey),
		[]byte(educationBoardKey),
	})
	fmt.Println(response.String())
	bytes, err := STUB.GetState(studentKey)
	if err != nil {
		fmt.Println("Error : " + err.Error())
		t.FailNow()
	}
	if len(bytes) == 0 {
		fmt.Println("Error : No data found for provided student id")
		t.FailNow()
	}
	fmt.Println(string(bytes))
}

func TestAcceptApplicationAsSchool(t *testing.T) {
	studentKey, err := GetStudentKey(STUB, "001")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	response := STUB.MockInvoke("2", [][]byte{
		[]byte("acceptApplicationAsSchool"),
		[]byte("001"),
		[]byte("Accepted from School"),
	})
	fmt.Println(response.String())
	bytes, err := STUB.GetState(studentKey)
	if err != nil {
		fmt.Println("Error : " + err.Error())
		t.FailNow()
	}
	if len(bytes) == 0 {
		fmt.Println("Error : No data found for provided student id")
		t.FailNow()
	}
	fmt.Println(string(bytes))
}

func TestAcceptApplicationAsEducationBoard(t *testing.T) {
	studentKey, err := GetStudentKey(STUB, "001")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	response := STUB.MockInvoke("2", [][]byte{
		[]byte("acceptApplicationAsEducationBoard"),
		[]byte("001"),
		[]byte("Accepted from EducationBoard"),
	})
	fmt.Println(response.String())
	bytes, err := STUB.GetState(studentKey)
	if err != nil {
		fmt.Println("Error : " + err.Error())
		t.FailNow()
	}
	if len(bytes) == 0 {
		fmt.Println("Error : No data found for provided student id")
		t.FailNow()
	}
	fmt.Println(string(bytes))
}

// func TestGetPendingListForEducationBoard(t *testing.T) {
// 	educationBoardKey, err := GetEducationBoardKey(STUB, "001")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		t.FailNow()
// 	}
// 	response := STUB.MockInvoke("2", [][]byte{
// 		[]byte("getPendingListForEducationBoard"),
// 		[]byte(educationBoardKey),
// 	})
// 	fmt.Println(response.String())
// }

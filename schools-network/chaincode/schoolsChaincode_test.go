package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	uuid "github.com/satori/go.uuid"
)

var CC, STUB = getInitializedChaincodeAndStub()

var StudentID string

var ApplicationID string

func UUID() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}

func getInitializedChaincodeAndStub() (*SchoolsChaincode, *shim.MockStub) {
	scc := new(SchoolsChaincode)
	scc.testing = true
	stub := shim.NewMockStub("MockStub", scc)
	stub.MockInit("1", nil)
	return scc, stub
}

func TestAddStudentDetails(t *testing.T) {
	response := STUB.MockInvoke(UUID(), GetBytesArrayFromStringArray([]string{
		"AddStudentDetails",
		"Lakshay Bhambri",
		"23-06-1995",
		"MALE",
		"Vinay Bhambri",
		"Vimmi Bhambri"}))
	student := Student{}
	err := json.Unmarshal(response.Payload, &student)
	if err != nil {
		fmt.Println("Error :: ", err.Error())
		t.FailNow()
	}
	StudentID = student.ID
	fmt.Println("*****************************")
	if response.Status != 200 {
		fmt.Println("Error : " + string(response.Message))
		t.FailNow()
	} else {
		fmt.Println("Success : " + string(response.Payload))
	}
	fmt.Println("*****************************")
}

func TestUpdateStudentDetails(t *testing.T) {
	response := STUB.MockInvoke(UUID(), GetBytesArrayFromStringArray([]string{
		"UpdateStudentDetails",
		StudentID,
		"Lakshay Bhambri",
		"24-06-1995",
		"MALE",
		"Vinay Bhambri",
		"Vimmi Bhambri"}))
	fmt.Println("*****************************")
	if response.Status != 200 {
		fmt.Println("Error : " + string(response.Message))
		t.FailNow()
	} else {
		fmt.Println("Success : " + string(response.Payload))
	}
	fmt.Println("*****************************")
}

func TestGetStudentDetails(t *testing.T) {
	response := STUB.MockInvoke(UUID(), GetBytesArrayFromStringArray([]string{
		"GetStudentDetails",
		StudentID}))
	fmt.Println("*****************************")
	if response.Status != 200 {
		fmt.Println("Error : " + string(response.Message))
		t.FailNow()
	} else {
		fmt.Println("Success : " + string(response.Payload))
	}
	fmt.Println("*****************************")
}

// func TestDeleteStudentDetails(t *testing.T) {
// 	response := STUB.MockInvoke(UUID(), GetBytesArrayFromStringArray([]string{
// 		"DeleteStudentDetails",
// 		StudentID}))
// 	fmt.Println("*****************************")
// 	if response.Status != 200 {
// 		fmt.Println("Error : " + string(response.Message))
// 		t.FailNow()
// 	} else {
// 		fmt.Println("Success : " + string(response.Payload))
// 	}
// 	fmt.Println("*****************************")
// }

func TestInitiateSchoolLeavingRequest(t *testing.T) {
	response := STUB.MockInvoke(UUID(), GetBytesArrayFromStringArray([]string{
		"InitiateSchoolLeavingRequest",
		StudentID,
		"Start School Leaving Request"}))
	fmt.Println("*****************************")
	if response.Status != 200 {
		fmt.Println("Error : " + string(response.Message))
		t.FailNow()
	} else {
		fmt.Println("Success : " + string(response.Payload))
		application := StudentTransferApplication{}
		err := json.Unmarshal(response.Payload, &application)
		if err != nil {
			fmt.Println("Error :: ", err.Error())
			t.FailNow()
		}
		ApplicationID = application.ID
	}
	fmt.Println("*****************************")
}

func TestRejectSchoolLeavingApplication(t *testing.T) {
	response := STUB.MockInvoke(UUID(), GetBytesArrayFromStringArray([]string{
		"RejectSchoolLeavingApplication",
		ApplicationID,
		"Student Still Studying"}))
	fmt.Println("*****************************")
	if response.Status != 200 {
		fmt.Println("Error : " + string(response.Message))
		t.FailNow()
	} else {
		fmt.Println("Success : " + string(response.Payload))
	}
	fmt.Println("*****************************")
}

func TestResubmitSchoolLeavingRequest(t *testing.T) {
	response := STUB.MockInvoke(UUID(), GetBytesArrayFromStringArray([]string{
		"ResubmitSchoolLeavingRequest",
		ApplicationID,
		"Defence Personnel Transfer"}))
	fmt.Println("*****************************")
	if response.Status != 200 {
		fmt.Println("Error : " + string(response.Message))
		t.FailNow()
	} else {
		fmt.Println("Success : " + string(response.Payload))
	}
	fmt.Println("*****************************")
}

func TestAcceptSchoolLeavingApplication(t *testing.T) {
	response := STUB.MockInvoke(UUID(), GetBytesArrayFromStringArray([]string{
		"AcceptSchoolLeavingApplication",
		ApplicationID,
		"Ok Accepted"}))
	fmt.Println("*****************************")
	if response.Status != 200 {
		fmt.Println("Error : " + string(response.Message))
		t.FailNow()
	} else {
		fmt.Println("Success : " + string(response.Payload))
	}
	fmt.Println("*****************************")
}

func TestInitiateSchoolAdmissionRequest(t *testing.T) {
	response := STUB.MockInvoke(UUID(), GetBytesArrayFromStringArray([]string{
		"InitiateSchoolAdmissionRequest",
		StudentID,
		"Lakshay Bhambri",
		"23-06-1995",
		"MALE",
		"Vinay Bhambri",
		"Vimmi Bhambri",
		"Start School Admission Request"}))
	fmt.Println("*****************************")
	if response.Status != 200 {
		fmt.Println("Error : " + string(response.Message))
		t.FailNow()
	} else {
		fmt.Println("Success : " + string(response.Payload))
		application := StudentTransferApplication{}
		err := json.Unmarshal(response.Payload, &application)
		if err != nil {
			fmt.Println("Error :: ", err.Error())
			t.FailNow()
		}
		ApplicationID = application.ID
	}
	fmt.Println("*****************************")
}

func TestRejectSchoolAdmissionApplication(t *testing.T) {
	response := STUB.MockInvoke(UUID(), GetBytesArrayFromStringArray([]string{
		"RejectSchoolAdmissionApplication",
		ApplicationID,
		"Student Still Studying"}))
	fmt.Println("*****************************")
	if response.Status != 200 {
		fmt.Println("Error : " + string(response.Message))
		t.FailNow()
	} else {
		fmt.Println("Success : " + string(response.Payload))
	}
	fmt.Println("*****************************")
}

func TestResubmitSchoolAdmissionRequest(t *testing.T) {
	response := STUB.MockInvoke(UUID(), GetBytesArrayFromStringArray([]string{
		"ResubmitSchoolAdmissionRequest",
		ApplicationID,
		"Defence Personnel Transfer"}))
	fmt.Println("*****************************")
	if response.Status != 200 {
		fmt.Println("Error : " + string(response.Message))
		t.FailNow()
	} else {
		fmt.Println("Success : " + string(response.Payload))
	}
	fmt.Println("*****************************")
}

func TestAcceptSchoolAdmissionApplication(t *testing.T) {
	response := STUB.MockInvoke(UUID(), GetBytesArrayFromStringArray([]string{
		"AcceptSchoolAdmissionApplication",
		ApplicationID,
		"Ok Accepted"}))
	fmt.Println("*****************************")
	if response.Status != 200 {
		fmt.Println("Error : " + string(response.Message))
		t.FailNow()
	} else {
		fmt.Println("Success : " + string(response.Payload))
	}
	fmt.Println("*****************************")
}

func GetBytesArrayFromStringArray(data []string) [][]byte {
	values := [][]byte{}
	for _, item := range data {
		values = append(values, []byte(item))
	}
	return values
}

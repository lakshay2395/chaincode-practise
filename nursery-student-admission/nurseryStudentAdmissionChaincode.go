package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type NurseryStudentAdmissionChaincode struct {
}

func (n *NurseryStudentAdmissionChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	schoolKey, err := GetSchoolKey(stub, "001")
	fmt.Println(schoolKey)
	if err != nil {
		return shim.Error("Error occurred : " + err.Error())
	}
	educationBoardKey, err := GetEducationBoardKey(stub, "001")
	fmt.Println(educationBoardKey)
	if err != nil {
		return shim.Error("Error occurred : " + err.Error())
	}
	stub.PutState(educationBoardKey, []byte(`{
		"educationBoardId" : "001",
		"name" : "Central Board Of Secondary Education",
		"headOfficer" : "Mr. Arun Joshi",
		"address" : "20 2nd Cross Greenlead Extension Kormangala Bengaluru"
	}`))
	stub.PutState(schoolKey, []byte(`{
		"schoolId" : "001",
		"name" : "Greenway Modern School",
		"principalName" : "Mr. Mohit Sachdeva",
		"address" : "1/550 G.T. Road MS Park Shahdara",
		"educationBoardId : "`+educationBoardKey+`"
	}`))
	return shim.Success(nil)
}

func (n *NurseryStudentAdmissionChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "apply":
		return n.apply(stub, args)
	case "reapply":
		return n.reapply(stub, args)
	case "acceptApplicationAsSchool":
		return n.acceptApplicationAsSchool(stub, args)
	case "acceptApplicationAsEducationBoard":
		return n.acceptApplicationAsEducationBoard(stub, args)
	case "rejectApplicationAsSchool":
		return n.rejectApplicationAsSchool(stub, args)
	case "rejectApplicationAsEducationBoard":
		return n.rejectApplicationAsEducationBoard(stub, args)
	case "getStudentDetails":
		return n.getStudentDetails(stub, args)
	case "getPendingListForSchool":
		return n.getPendingListForSchool(stub, args)
	case "getPendingListForEducationBoard":
		return n.getPendingListForEducationBoard(stub, args)
	}
	return shim.Error("Incorrect Function Name Provided = " + function)
}

func (n *NurseryStudentAdmissionChaincode) reapply(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments found. Expected 9, Found " + string(len(args)))
	}
	key, err := GetStudentKey(stub, args[0])
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	studentDetailsBytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	if len(studentDetailsBytes) == 0 {
		return shim.Error("No such existing application exists for student id = " + args[0])
	}
	studentDetails := StudentDetails{
		StudentId:  args[0],
		Name:       args[1],
		DOB:        args[2],
		Gender:     args[3],
		FatherName: args[4],
		MotherName: args[5],
		Address:    args[6],
		City:       args[7],
		SchoolID:   args[8],
	}
	studentDetails.Status = STUDENT_APPLIED
	studentDetailsBytes, err = json.Marshal(studentDetails)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	err = stub.PutState(key, studentDetailsBytes)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	return shim.Success([]byte("Application Successfully Re-Submitted. Student ID : " + key))
}

func (n *NurseryStudentAdmissionChaincode) apply(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 10 {
		return shim.Error("Incorrect number of arguments found. Expected 10, Found " + string(len(args)))
	}
	studentDetails := StudentDetails{
		StudentId:        args[0],
		Name:             args[1],
		DOB:              args[2],
		Gender:           args[3],
		FatherName:       args[4],
		MotherName:       args[5],
		Address:          args[6],
		City:             args[7],
		SchoolID:         args[8],
		EducationBoardID: args[9],
	}
	studentDetails.Status = STUDENT_APPLIED
	studentDetailsBytes, err := json.Marshal(studentDetails)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	key, err := GetStudentKey(stub, args[0])
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	err = stub.PutState(key, studentDetailsBytes)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	return shim.Success([]byte("Application Successfully Submitted. Student ID : " + key))
}

func (n *NurseryStudentAdmissionChaincode) acceptApplicationAsSchool(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments found. Expected 2, Found " + string(len(args)))
	}
	key, err := GetStudentKey(stub, args[0])
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	studentDetailsBytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	if len(studentDetailsBytes) == 0 {
		return shim.Error("No student details found with ID = " + args[0])
	}
	studentDetails := StudentDetails{}
	err = json.Unmarshal(studentDetailsBytes, &studentDetails)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	studentDetails.Status = APPLICATION_ACCEPTED_BY_SCHOOL
	studentDetails.SchoolRemark = args[1]
	studentDetailsBytes, err = json.Marshal(studentDetails)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	err = stub.PutState(key, studentDetailsBytes)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	return shim.Success([]byte("Student Successfully Accepted at Student Level"))
}

func (n *NurseryStudentAdmissionChaincode) acceptApplicationAsEducationBoard(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments found. Expected 2, Found " + string(len(args)))
	}
	key, err := GetStudentKey(stub, args[0])
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	studentDetailsBytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	if len(studentDetailsBytes) == 0 {
		return shim.Error("No student details found with ID = " + args[0])
	}
	studentDetails := StudentDetails{}
	err = json.Unmarshal(studentDetailsBytes, &studentDetails)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	studentDetails.Status = STUDENT_ADMITTED
	studentDetails.EducationBoardRemark = args[1]
	studentDetailsBytes, err = json.Marshal(studentDetails)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	err = stub.PutState(key, studentDetailsBytes)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	return shim.Success([]byte("Student Successfully Admitted"))
}

func (n *NurseryStudentAdmissionChaincode) rejectApplicationAsSchool(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments found. Expected 2, Found " + string(len(args)))
	}
	key, err := GetStudentKey(stub, args[0])
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	studentDetailsBytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	if len(studentDetailsBytes) == 0 {
		return shim.Error("No student details found with ID = " + args[0])
	}
	studentDetails := StudentDetails{}
	err = json.Unmarshal(studentDetailsBytes, &studentDetails)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	studentDetails.Status = APPLICATION_REJECTED_BY_SCHOOL
	studentDetails.SchoolRemark = args[1]
	studentDetailsBytes, err = json.Marshal(studentDetails)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	err = stub.PutState(key, studentDetailsBytes)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	return shim.Success([]byte("Student Rejected at Student Level"))
}

func (n *NurseryStudentAdmissionChaincode) rejectApplicationAsEducationBoard(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments found. Expected 2, Found " + string(len(args)))
	}
	key, err := GetStudentKey(stub, args[0])
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	studentDetailsBytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	if len(studentDetailsBytes) == 0 {
		return shim.Error("No student details found with ID = " + args[0])
	}
	studentDetails := StudentDetails{}
	err = json.Unmarshal(studentDetailsBytes, &studentDetails)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	studentDetails.Status = APPLICATION_REJECTED_BY_EDUCATION_BOARD
	studentDetails.SchoolRemark = args[1]
	studentDetailsBytes, err = json.Marshal(studentDetails)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	err = stub.PutState(key, studentDetailsBytes)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	return shim.Success([]byte("Student Rejected at Education Board Level"))
}

func (n *NurseryStudentAdmissionChaincode) getStudentDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments found. Expected 1, Found " + string(len(args)))
	}
	key, err := GetStudentKey(stub, args[0])
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	studentDetailsBytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	if len(studentDetailsBytes) == 0 {
		return shim.Error("No student details found with ID = " + args[0])
	}
	return shim.Success(studentDetailsBytes)
}

func (n *NurseryStudentAdmissionChaincode) getPendingListForSchool(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments found. Expected 1, Found " + string(len(args)))
	}
	queryString := `{ 
		"selector" : {
			"status" : 0,
			"schoolId" : "` + args[0] + `"
		}
	}`
	iterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	list := []StudentDetails{}
	for iterator.HasNext() {
		response, err := iterator.Next()
		if err != nil {
			return shim.Error("Error Occurred : " + err.Error())
		}
		studentDetails := StudentDetails{}
		err = json.Unmarshal(response.Value, &studentDetails)
		if err != nil {
			return shim.Error("Error Occurred : " + err.Error())
		}
		list = append(list, studentDetails)
	}
	listBytes, err := json.Marshal(list)
	if err != nil {
		return shim.Error("Error Occurred : " + err.Error())
	}
	return shim.Success(listBytes)
}

func (n *NurseryStudentAdmissionChaincode) getPendingListForEducationBoard(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments found. Expected 1, Found " + string(len(args)))
	}
	queryString := `{ 
		"selector" : {
			"status" : 2,
			"educationBoardId" : "` + args[0] + `"
		}
	}`
	iterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error("Error Occurred 1 : " + err.Error())
	}
	list := []StudentDetails{}
	for iterator.HasNext() {
		response, err := iterator.Next()
		if err != nil {
			return shim.Error("Error Occurred 2 : " + err.Error())
		}
		studentDetails := StudentDetails{}
		err = json.Unmarshal(response.Value, &studentDetails)
		if err != nil {
			return shim.Error("Error Occurred 3 : " + err.Error())
		}
		list = append(list, studentDetails)
	}
	listBytes, err := json.Marshal(list)
	if err != nil {
		return shim.Error("Error Occurred 4 : " + err.Error())
	}
	return shim.Success(listBytes)
}

func main() {
	nsacc := new(NurseryStudentAdmissionChaincode)
	err := shim.Start(nsacc)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error occurring in starting chaincode : %s", err.Error()))
	}
}

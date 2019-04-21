package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type HelloWorldChaincode struct {
}

func (h *HelloWorldChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Init Invoked")
	return shim.Success(nil)
}

func (h *HelloWorldChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, _ := stub.GetFunctionAndParameters()
	if function == "helloWorld" {
		err := stub.PutState("message", []byte("Hello World"))
		if err != nil {
			return shim.Error(err.Error())
		}
	}
	return shim.Success(nil)
}

func main() {
	hwcc := new(HelloWorldChaincode)
	err := shim.Start(hwcc)
	if err != nil {
		fmt.Println(fmt.Sprintf("Unable to start hello world chaincode : %s", err.Error()))
	}
}

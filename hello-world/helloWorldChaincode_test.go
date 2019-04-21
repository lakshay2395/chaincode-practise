package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestHelloWorldChaincodeInit(t *testing.T) {
	hwcc := new(HelloWorldChaincode)
	stub := shim.NewMockStub("MockStub", hwcc)
	stub.MockInit("1", nil)
}

func TestHelloWorldChaincodeInvoke(t *testing.T) {
	hwcc := new(HelloWorldChaincode)
	stub := shim.NewMockStub("MockStub", hwcc)
	stub.MockInit("1", nil)
	stub.MockInvoke("2", [][]byte{
		[]byte("helloWorld"),
	})
	bytes, err := stub.GetState("message")
	if err != nil {
		fmt.Println("Error : " + err.Error())
		t.FailNow()
	}
	if len(bytes) == 0 {
		fmt.Println("Error : Value for message is empty")
		t.FailNow()
	}
	if string(bytes) != "Hello World" {
		fmt.Println("Error : Value for message is not expected")
		t.FailNow()
	}
}

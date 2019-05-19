package main

import (
	"fmt"
	"os"
	"time"

	"github.com/schools-network-client/blockchain"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters
		OrdererID: "orderer.nic.in",

		// Channel parameters
		ChannelID:     "schooloneorg",
		ChannelConfig: "/home/lakshay/Desktop/hyperledger-fabric/chaincode-practise/schools-network/network/channel-artifacts/schooloneorg/schooloneorg.tx",

		// Chaincode parameters
		ChainCodeID:     "scc",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/schools-network/",
		OrgAdmin:        "Admin",
		OrgName:         "SchoolOneOrg",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "Admin",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	output, err := fSetup.Query()
	if err != nil {
		fmt.Printf("Unable to query: %v\n", err)
		// return
	} else {
		fmt.Println("output = ", output)
		time.Sleep(time.Duration(100) * time.Second)
	}
	// Close SDK
	// time.Sleep(time.Duration(100) * time.Second)
	defer fSetup.CloseSDK()
}

package blockchain

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/pkg/errors"
)

// FabricSetup implementation
type FabricSetup struct {
	ConfigFile      string
	OrgID           string
	OrdererID       string
	ChannelID       string
	ChainCodeID     string
	initialized     bool
	ChannelConfig   string
	ChaincodeGoPath string
	ChaincodePath   string
	OrgAdmin        string
	OrgName         string
	UserName        string
	admin           *resmgmt.Client
	sdk             *fabsdk.FabricSDK
	client          *channel.Client
	event           *event.Client
}

func (setup *FabricSetup) Initialize() error {

	if setup.initialized {
		return errors.New("sdk already initialized")
	}
	sdk, err := fabsdk.New(config.FromFile(setup.ConfigFile))
	if err != nil {
		return errors.WithMessage(err, "failed to create SDK")
	}
	setup.sdk = sdk
	resourceManagerClientContext := setup.sdk.Context(fabsdk.WithUser(setup.OrgAdmin), fabsdk.WithOrg(setup.OrgName))
	if err != nil {
		return errors.WithMessage(err, "failed to load Admin identity")
	}
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create channel management client from Admin identity")
	}
	setup.admin = resMgmtClient
	fmt.Println("Ressource management client created")
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(setup.UserName))
	setup.client, err = channel.New(clientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create new channel client")
	}
	setup.initialized = true
	return nil
}

func GetBytesArrayFromStringArray(data []string) [][]byte {
	values := [][]byte{}
	for _, item := range data {
		values = append(values, []byte(item))
	}
	return values
}

func (setup *FabricSetup) Query() (string, error) {

	data := GetBytesArrayFromStringArray([]string{
		"Srishty Bhambri",
		"23-06-1995",
		"MALE",
		"Vinay Bhambri",
		"Vimmi Bhambri"})

	// data := GetBytesArrayFromStringArray([]string{"350af37b-745d-4846-b18f-c6eb58970aca"})
	// setup.Client.
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "AddStudentDetails", Args: data})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	output := make(map[string]string)
	json.Unmarshal(response.Payload, &output)
	fmt.Println("output = ", output)
	data = GetBytesArrayFromStringArray([]string{output["id"]})
	response, err = setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "GetStudentDetails", Args: data})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil
}

func (setup *FabricSetup) CloseSDK() {
	setup.sdk.Close()
}

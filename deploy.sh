#!/bin/sh
echo "****************************************"
echo "DEPLOY CHAINCODE ON BASIC FABRIC NETWORK"
echo "****************************************"
read -p "Enter project folder name : " folderName
read -p "Enter version : " version
rm -rf $GOPATH/src/github.com/$folderName
mkdir $GOPATH/src/github.com/$folderName
cp -rf $folderName/* $GOPATH/src/github.com/$folderName
cd $GOPATH/src/github.com/$folderName
go test
go build .
cd ~/fabric-samples/basic-network
./start.sh -d true
LANGUAGE=golang
CC_SRC_PATH=github.com/$folderName
CC_VERSION=$version
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode install -n $folderName -v $CC_VERSION -p "$CC_SRC_PATH" -l "$LANGUAGE"
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n $folderName -l "$LANGUAGE" -v $CC_VERSION -c '{"Args":[""]}' -P "OR ('Org1MSP.member')"
echo "*******************************************************"
echo "CHAINCODE SUCCESSFULLY DEPLOYED ON BASIC FABRIC NETWORK"
echo "*******************************************************"
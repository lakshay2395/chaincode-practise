package main

import (
	"errors"
	"log"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func IsSchool(stub shim.ChaincodeStubInterface) bool {
	cert, err := cid.GetX509Certificate(stub)
	if err != nil {
		log.Println("Error :: ", err.Error())
		return false
	}
	if "ca."+stub.GetChannelID()+".school.nic.in" == cert.Issuer.CommonName {
		return true
	}
	return false
}

func IsCBSE(stub shim.ChaincodeStubInterface) bool {
	cert, err := cid.GetX509Certificate(stub)
	if err != nil {
		log.Println("Error :: ", err.Error())
		return false
	}
	if "ca.cbseorg.nic.in" == cert.Issuer.CommonName {
		return true
	}
	return false
}

func (sc *SchoolsChaincode) CheckCBSEAuthorization(stub shim.ChaincodeStubInterface) error {
	if sc.testing {
		return nil
	}
	mspid, err := cid.GetMSPID(stub)
	if err != nil {
		return err
	}
	cert, err := cid.GetX509Certificate(stub)
	if err != nil {
		return err
	}
	if mspid == "CBSEOrg" && cert.Issuer.CommonName == "ca.cbseorg.nic.in" {
		return errors.New("Unauthorized Access :: " + cert.Issuer.CommonName + "is not authorized to perform this operation")
	}
	return nil
}

func (sc *SchoolsChaincode) CheckSchoolAuthorization(stub shim.ChaincodeStubInterface) error {
	if sc.testing {
		return nil
	}
	cert, err := cid.GetX509Certificate(stub)
	if err != nil {
		return err
	}
	if "ca."+stub.GetChannelID()+".school.nic.in" != cert.Issuer.CommonName {
		return errors.New("Unauthorized Access :: " + cert.Issuer.CommonName + "is not authorized to perform this operation")
	}
	return nil
}

func (sc *SchoolsChaincode) CheckMutualAuthorization(stub shim.ChaincodeStubInterface) error {
	if sc.testing {
		return nil
	}
	cert, err := cid.GetX509Certificate(stub)
	if err != nil {
		return err
	}
	if strings.HasSuffix(cert.Issuer.CommonName, ".school.nic.in") {
		if "ca."+stub.GetChannelID()+".school.nic.in" != cert.Issuer.CommonName {
			return errors.New("Unauthorized Access :: " + cert.Issuer.CommonName + "is not authorized to perform this operation")
		} else {
			return nil
		}
	} else {
		if cert.Issuer.CommonName == "ca.cbseorg.nic.in" {
			return nil
		}
	}
	return errors.New("Unauthorized Access :: " + cert.Issuer.CommonName + "is not authorized to perform this operation")
}

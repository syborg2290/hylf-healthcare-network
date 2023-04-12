package main

import (
	healthcare "healthcare/Go"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	healthcareChaincode, err := contractapi.NewChaincode(&healthcare.SmartContract{})
	if err != nil {
		log.Panicf("Error creating healthcare chaincode: %v", err)
	}

	if err := healthcareChaincode.Start(); err != nil {
		log.Panicf("Error starting healthcare chaincode: %v", err)
	}
}

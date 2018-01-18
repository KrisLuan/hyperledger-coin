package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/hyperledger-coin/pocket"
)

func main() {
	//shim.SetLoggingLevel(shim.LoggingLevel(logging.ERROR))
	if err := shim.Start(&pocket.PocketChaincode{}); err != nil {
		panic(err)
	}
}
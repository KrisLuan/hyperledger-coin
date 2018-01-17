package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func main() {
	//shim.SetLoggingLevel(shim.LoggingLevel(logging.ERROR))
	if err := shim.Start(&); err != nil {
		panic(err)
	}
}
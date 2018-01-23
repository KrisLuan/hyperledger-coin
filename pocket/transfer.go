package pocket

import (
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *PocketChaincode)transfer(store Store, args []string) pb.Response {



	return shim.Success(nil)
}
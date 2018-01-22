package pocket

import (
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *PocketChaincode)registerAccount(store Store, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(ErrInvalidArgs.Error())
	}

	if tmpPocket, err := store.GetPocket {
		
	}
}
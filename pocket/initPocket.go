package pocket

import (
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

func (t *PocketChaincode)newPocket(store Store, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error(ErrInvalidArgs.Error())
	}

	kind := args[0]
	addr := args[1]
	pubkey := args[2]
	totalPoint, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		return shim.Error(err.Error())
	}
	if err := store.ModifyPointKind(kind); err != nil {
		return shim.Error(err.Error())
	}

	if err := store.InitPocketStatistics(addr, pubkey, totalPoint); err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

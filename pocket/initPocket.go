package pocket

import (
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

func (t *PocketChaincode)newPocket(store Store, args []string) pb.Response {
	if len(args) != 4 {
		logger.Error(ErrInvalidArgs.Error())
		return shim.Error(ErrInvalidArgs.Error())
	}

	kind := args[0]
	addr := args[1]
	pubkey := args[2]
	logger.Debugf("init new pocket, kind [%s], addr [%s], pubkey [%s], total point [%s]", kind, addr, pubkey, args[3])
	totalPoint, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	logger.Debugf("modify point kind")
	if err := store.ModifyPointKind(kind); err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	logger.Debugf("init [%s] pocket statistics", kind)
	if err := store.InitPocketStatistics(addr, pubkey, totalPoint); err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

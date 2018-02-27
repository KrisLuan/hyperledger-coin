package pocket

import (
pb "github.com/hyperledger/fabric/protos/peer"
"github.com/hyperledger/fabric/core/chaincode/shim"
"strconv"
)

func (t *PocketChaincode)modifyTxFee(store Store, args []string) pb.Response {
	if len(args) != 2 {
		logger.Error(ErrInvalidArgs.Error())
		return shim.Error(ErrInvalidArgs.Error())
	}

	ratio, err := strconv.ParseInt(args[1], 10, 32)
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	txFeeInfo, err := store.GetTxFeeInfo()
	if err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	txFeeInfo.Ratio = ratio
	if err := store.PutTxFeeInfo(txFeeInfo); err != nil {
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

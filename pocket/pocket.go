package pocket

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//需要检查，整个系统中的数据不能有重复读取的问题，否则会导致数据不一致
type PocketChaincode struct {
}

func (t *PocketChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Debugf("pocket chaincode Init")
	function, _ := stub.GetFunctionAndParameters()
	if function != "init" {
		return shim.Error(ErrInvalidArgs.Error())
	}

	store := MakeChaincodeStore(stub, DefaultPocketKind)
	err := store.InitPocket(InitAddr,InitPubkey, InitTotalPoint)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//Invoke function
const (
	IF_INITPOCKET string	= "invoke_initpocket"
	IF_REGISTER string		= "invoke_register"
	IF_TRANSFER string		= "invoke_transfer"
	IF_MODIFYTXFEE string	= "invoke_modifytxfee"

	QF_ADDRS string			= "query_addrs"
	QF_BALANCE string		= "query_balance"
	QF_STATISTICS string	= "query_statistics"
	QF_POINTKIND string		= "query_pointkind"
)

func (t *PocketChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Debugf("pocket chaincode invoke")
	function, args := stub.GetFunctionAndParameters()
	if len(args) == 0 {
		logger.Error(ErrInvalidArgs.Error());
		return shim.Error(ErrInvalidArgs.Error())
	}
	logger.Debugf("function [%v] args [%v]", function, args)
	store := MakeChaincodeStore(stub, args[0])

	switch function {
	case IF_INITPOCKET:
		return t.newPocket(store, args)
	case IF_REGISTER:
		return t.registerAccount(store, args)
	case IF_TRANSFER:
		return t.transfer(store, args)
	case IF_MODIFYTXFEE:
		return t.modifyTxFee(store, args)
	case QF_ADDRS:
		return t.queryAddrs(store, args)
	case QF_BALANCE:
		return t.queryBalance(store, args)
	case QF_STATISTICS:
		return t.queryStatistics(store, args)
	case QF_POINTKIND:
		return t.queryPointkind(store, args)
	default:
		logger.Error(ErrInvalidFunction.Error())
		return shim.Error(ErrInvalidFunction.Error())
	}

	return shim.Error(ErrInvalidFunction.Error())
}

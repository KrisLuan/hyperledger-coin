package pocket

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


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

	QF_ADDRS string			= "query_addrs"
	QF_STATISTICS string	= "query_statistics"
	QF_POINTKIND string		= "query_pointkind"
)

func (t *PocketChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Debugf("pocket chaincode invoke")
	function, args := stub.GetFunctionAndParameters()
	store := MakeChaincodeStore(stub, args[0])

	switch function {
	case IF_INITPOCKET:
		return t.newPocket(store, args)
	case IF_REGISTER:
		return t.registerAccount(store, args)
	case IF_TRANSFER:
		return shim.Error("")
	case QF_ADDRS:
		return shim.Error("")
	case QF_STATISTICS:
		return shim.Error("")
	case QF_POINTKIND:
		return shim.Error("")
	default:
		return shim.Error(ErrInvalidFunction.Error())
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

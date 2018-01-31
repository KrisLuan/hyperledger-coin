package pocket

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"encoding/base64"
	"github.com/golang/protobuf/proto"
)

func (t *PocketChaincode)queryAddrs(store Store, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error(ErrInvalidArgs.Error())
	}
	logger.Debugf("kind [%v] addr [%v]", args[0], args[1])

	addr := args[1]
	assets, nounce, err := store.GetAllAssets(addr)
	if err != nil {
		return shim.Error(err.Error())
	}

	queryResult := new(QueryResult)
	queryResult.Balance = assets
	//balance 作为nounce，保证不能重放
	queryResult.Nounce = nounce

	logger.Debugf("query result, balance [%v] nounce [%v]", queryResult.Balance, queryResult.Nounce)
	data, err := proto.Marshal(queryResult)
	if err != nil {
		return shim.Error(err.Error())
	}
	protobytebase64 := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(protobytebase64, []byte(data))
	return shim.Success(protobytebase64)
}

func (t *PocketChaincode)queryStatistics(store Store, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(ErrInvalidArgs.Error())
	}
	logger.Debugf("kind [%v]", args[0])

	pointInfo, err := store.GetPointInfo()
	if err != nil {
		return shim.Error(err.Error())
	}

	logger.Debugf("point info [%v]", pointInfo)
	data, err := proto.Marshal(pointInfo)
	if err != nil {
		return shim.Error(err.Error())
	}
	protobytebase64 := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(protobytebase64, []byte(data))
	return shim.Success(protobytebase64)
}

func (t *PocketChaincode)queryPointkind(store Store, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(ErrInvalidArgs.Error())
	}

	pointKind, err := store.GetPointKind()
	if err != nil {
		return shim.Error(err.Error())
	}

	logger.Debugf("point kind [%v]", pointKind)
	data, err := proto.Marshal(pointKind)
	if err != nil {
		return shim.Error(err.Error())
	}
	protobytebase64 := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(protobytebase64, []byte(data))
	return shim.Success(protobytebase64)
}

func (t *PocketChaincode)queryBalance(store Store, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error(ErrInvalidArgs.Error())
	}

	addr := args[1]
	assets, _, err := store.GetAllAssets(addr)
	if err != nil {
		return shim.Error(err.Error())
	}

	logger.Debugf("balance [%v]", assets)
	s := strconv.FormatInt(assets, 10)
	protobytebase64 := make([]byte, base64.StdEncoding.EncodedLen(len(s)))
	base64.StdEncoding.Encode(protobytebase64, []byte(s))
	return shim.Success(protobytebase64)
}
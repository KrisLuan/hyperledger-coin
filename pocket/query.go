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

	addr := args[1]
	pocket, err := store.GetPocket(addr)
	if err != nil {
		return shim.Error(err.Error())
	}
	s := strconv.FormatInt(pocket.Balance, 10)
	protobytebase64 := make([]byte, base64.StdEncoding.EncodedLen(len(s)))
	base64.StdEncoding.Encode(protobytebase64, []byte(s))
	return shim.Success(protobytebase64)
}

func (t *PocketChaincode)queryStatistics(store Store, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(ErrInvalidArgs.Error())
	}

	pointInfo, err := store.GetPointInfo()
	if err != nil {
		return shim.Error(err.Error())
	}
	data, err := proto.Marshal(pointInfo)
	protobytebase64 := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(protobytebase64, []byte(data))
	return shim.Success(protobytebase64)
}

func (t *PocketChaincode)queryPointkind(store Store, args []string) pb.Response {
	if len(args) != 0 {
		return shim.Error(ErrInvalidArgs.Error())
	}

	pointKind, err := store.GetPointKind()
	if err != nil {
		return shim.Error(err.Error())
	}

	data, err := proto.Marshal(pointKind)
	protobytebase64 := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(protobytebase64, []byte(data))
	return shim.Success(protobytebase64)
}
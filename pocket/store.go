package pocket

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/golang/protobuf/proto"
	"fmt"
)


type Store interface {
	InitPocket() error
	generateKey(string) string

	GetPointInfo() (*PointInfo, error)
	PutPointInfo(*PointInfo) error

	GetPocket(string) (*Pocket, error)
	PutPocket(*Pocket) error
}

// Store struct uses a chaincode stub for state access
type ChaincodeStore struct {
	stub shim.ChaincodeStubInterface
	kind string
}

// MakeChaincodeStore returns a store for storing keys in the state
func MakeChaincodeStore(stub shim.ChaincodeStubInterface, kind string) Store {
	store := &ChaincodeStore{}
	store.stub = stub
	store.kind = kind
	return store
}

func (s *ChaincodeStore)InitPocket() error {
	coinInfo := &PointInfo{
		AccountTotal:	0,
		TxTotal:	0,
		PointTotal:	0,
		Holder:		"foam",
	}

	return s.PutPointInfo(coinInfo)
}

//地址不允许包含‘_’，积分种类也不允许包含‘_’
func (s *ChaincodeStore) generateKey(addr string) string {
	return fmt.Sprintf("%s_%s", s.kind, addr)
}

func (s *ChaincodeStore) GetPointInfo() (*PointInfo, error) {
	data, err := s.stub.GetState(s.generateKey(pointInfoKey))
	if err != nil {
		return nil, err
	}

	return ParsePointInfo(data)
}

func (s *ChaincodeStore) PutPointInfo(pointInfo *PointInfo) error {
	coinBytes, err := proto.Marshal(pointInfo)
	if err != nil {
		return err
	}

	if err := s.stub.PutState(s.generateKey(pointInfoKey), coinBytes); err != nil {
		return err
	}

	return nil
}

func (s *ChaincodeStore) GetPocket(addr string) (*Pocket, error) {
	if addr == "" {
		return nil, ErrEmptyAddr
	}
	key := s.generateKey(addr)
	data, err := s.stub.GetState(key)
	if err != nil {
		return nil, err
	}

	return ParsePocket(data)
}

func (s *ChaincodeStore) PutPocket(pocket *Pocket) error {
	key := s.generateKey(pocket.Addr)

	data, err := proto.Marshal(pocket)
	if err != nil {
		return err
	}

	return s.stub.PutState(key, data)
}


















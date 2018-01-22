package pocket

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/golang/protobuf/proto"
	"fmt"
)


type Store interface {
	InitPocket(string, int64) error
	generateKey(string) string

	GetPointInfo() (*PointInfo, error)
	PutPointInfo(*PointInfo) error

	GetPocket(string) (*Pocket, error)
	PutPocket(*Pocket) error

	GetPointKind() (*PointKind, error)
	PutPointKind(*PointKind) error
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

//可能需要加入公钥字段，同时验证私钥的合法性
func (s *ChaincodeStore)InitPocket(addr string, totalPoint int64) error {
	kind := []string{DefaultPocketKind}
	pointKind := &PointKind{
		Kind: kind,
	}

	logger.Debugf("put point kind into fabric")
	if err := s.PutPointKind(pointKind); err != nil {
		return err
	}

	logger.Debugf("init pocket")
	return initPocket(addr, totalPoint, s)
}

func initPocket(addr string, totalPoint int64, s *ChaincodeStore) error {
	if IsValidAddr(InitAddr) {
		return ErrInvalidAddr
	}
	//TODO verfiy pubkey and addr

	pointInfo := &PointInfo{
		AccountTotal:	0,
		TxTotal:	0,
		PointTotal:	0,
		Holder:		"foam",
	}

	pocket := &Pocket{
		Addr: 		addr,
		Balance:	totalPoint,
	}
	logger.Debugf("init [%s] [%s]", addr, totalPoint)
	if err := s.PutPocket(pocket); err != nil {
		return err
	}

	logger.Debugf("init point info")
	return s.PutPointInfo(pointInfo)
}

//地址不允许包含‘_’，积分种类也不允许包含‘_’
func (s *ChaincodeStore) generateKey(addr string) string {
	return fmt.Sprintf("%s_%s", s.kind, addr)
}

func (s *ChaincodeStore) GetPointInfo() (*PointInfo, error) {
	logger.Debugf("get point info")
	data, err := s.stub.GetState(s.generateKey(pointInfoKey))
	if err != nil {
		return nil, err
	}

	return ParsePointInfo(data)
}

func (s *ChaincodeStore) PutPointInfo(pointInfo *PointInfo) error {
	logger.Debugf("put point info [%v]", pointInfo)
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
	logger.Debugf("get pocket with [%s]", addr)
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
	logger.Debugf("put pocket [%v]", pocket)
	key := s.generateKey(pocket.Addr)

	data, err := proto.Marshal(pocket)
	if err != nil {
		return err
	}

	return s.stub.PutState(key, data)
}

func (s *ChaincodeStore) PutPointKind(pointKind *PointKind) error {
	logger.Debugf("put point kind [%v]", pointKind)
	data, err := proto.Marshal(pointKind)
	if err != nil {
		return err
	}
	return s.stub.PutState(kindKey, data)
}

func (s *ChaincodeStore) GetPointKind() (*PointKind, error) {
	logger.Debugf("get point kind")
	data, err := s.stub.GetState(kindKey)
	if err != nil {
		return nil, err
	}

	return ParsePointKind(data)
}















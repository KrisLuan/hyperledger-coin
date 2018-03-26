package pocket

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger-coin/proto"
	"fmt"
	"strconv"
)

//TODO: fomat log print

type Store interface {
	InitPocket(addr string, pubkey string, totalPoint int64) error
	InitPocketStatistics(addr string, pubkey string, totalPoint int64) error
	generateKey(addr string) string

	GetPointInfo() (*PointInfo, error)
	PutPointInfo(pointInfo *PointInfo) error

	GetPocket(addr string) (*Pocket, error)
	PutPocket(pocket *Pocket) error
	GetAllAssets(addr string) (int64, int64, error)
	MergeStateByPartialCompositeKey(objectType string, keys []string) (int64, error)
	AddCompositeOutput(objectType string, attributes []string, value int64) error

	GetPointKind() (*PointKind, error)
	PutPointKind(pointKind *PointKind) error

	ModifyPointInfo(increaseAccount int64, increaseTx int64, increasePoint int64) error
	ModifyPointKind(kind string) error

	PutTxFeeInfo(txFeeInfo *TxFeeInfo) error
	GetTxFeeInfo() (*TxFeeInfo, error)

	GetTxID() string
}

// Store struct uses a chaincode stub for state access
type ChaincodeStore struct {
	stub shim.ChaincodeStubInterface
	kind string
}

//可能需要加入公钥字段，同时验证私钥的合法性
/*
	初始化设置
*/
func (s *ChaincodeStore)InitPocket(addr string, pubkey string, totalPoint int64) error {
	kind := []string{DefaultPocketKind}
	pointKind := &PointKind{
		Kind: kind,
	}

	logger.Debugf("put point kind into fabric")
	if err := s.PutPointKind(pointKind); err != nil {
		return err
	}

	logger.Debugf("init pocket")
	return s.InitPocketStatistics(addr, pubkey, totalPoint)
}

//初始化积分统计信息和初始积分
/*
	这个函数主要包括以下操作
	1.设置交易费以及交易费收取地址
	2.设置初始账户
	3.在区块链上设置账户信息
*/
func (s *ChaincodeStore)InitPocketStatistics(addr string, pubkey string, totalPoint int64) error {
	if !IsValidAddr(InitAddr, pubkey) {
		logger.Debugf("test")
		return ErrInvalidAddr
	}
	//TODO verfiy pubkey and addr

	logger.Debugf("put tx fee info")
	txFeeInfo := &TxFeeInfo{
		TxFeeAddr:	addr,
		Ratio:		2,
	}
	if err := s.PutTxFeeInfo(txFeeInfo); err != nil {
		return err
	}

	pointInfo := &PointInfo{
		AccountTotal:	0,
		TxTotal:	0,
		PointTotal:	0,
		Holder:		"foam",
	}

	pocket := &Pocket{
		Addr: 		addr,
		Balance:	totalPoint,
		Pubkey:		pubkey,
	}
	logger.Debugf("init [%s] [%s]", addr, totalPoint)
	if err := s.PutPocket(pocket); err != nil {
		return err
	}

	logger.Debugf("init point info")
	return s.PutPointInfo(pointInfo)
}

//地址不允许包含‘_’，积分种类也不允许包含‘_’
//生成key, kind_addr 	--ly
func (s *ChaincodeStore) generateKey(addr string) string {
	return fmt.Sprintf("%s_%s", s.kind, addr)
}

//  pointInfoKey是定义的一个常量，需要考虑之前存的是什么   --ly
func (s *ChaincodeStore) GetPointInfo() (*PointInfo, error) {
	data, err := s.stub.GetState(s.generateKey(pointInfoKey))
	if err != nil {
		return nil, err
	}

	return ParsePointInfo(data)
}

//这里是设置PointInfo的地方，
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


/*
	获取用户存在区块链上的信息，并用 ParsePocket()解析，
	其中涉及一个proto的数据结构，如下：
	message Pocket {
    		string addr = 1;    //用户地址
    		int64 balance = 2;  //余额
    		string pubkey = 3;  //用户公钥
	}
*/
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


//与上面那个函数相互对应，在区块链上记录数据信息
func (s *ChaincodeStore) PutPocket(pocket *Pocket) error {
	logger.Debugf("put pocket [%v]", pocket)
	key := s.generateKey(pocket.Addr)

	data, err := proto.Marshal(pocket)
	if err != nil {
		return err
	}

	return s.stub.PutState(key, data)
}

//在区块链上记录积分种类，pointKind是proto的数据结构，kindKey 在constant.go中被定义 --ly
func (s *ChaincodeStore) PutPointKind(pointKind *PointKind) error {
	logger.Debugf("put point kind [%v]", pointKind)
	data, err := proto.Marshal(pointKind)
	if err != nil {
		return err
	}
	return s.stub.PutState(kindKey, data)
}

//获取积分数据结构Pointkind信息，ParsePointKind在utils.go中被定义	--ly
func (s *ChaincodeStore) GetPointKind() (*PointKind, error) {
	data, err := s.stub.GetState(kindKey)
	if err != nil {
		return nil, err
	}

	return ParsePointKind(data)
}

/*
	积分属性信息的修改
*/
func (s *ChaincodeStore) ModifyPointInfo(increaseAccount int64, increaseTx int64, increasePoint int64) error {
	pointInfo, err := s.GetPointInfo()
	if err != nil {
		return err
	}

	logger.Debugf("modify point before [%v]", pointInfo)
	pointInfo.AccountTotal += increaseAccount
	pointInfo.TxTotal += increaseTx
	pointInfo.PointTotal += increasePoint

	return s.PutPointInfo(pointInfo)
}

//修改积分种类信息，加入其他的积分种类 --ly
func (s * ChaincodeStore) ModifyPointKind(kind string) error {
	pointKind, err := s.GetPointKind()
	if err != nil {
		return err
	}

	pointKind.Kind = append(pointKind.Kind, kind)
	logger.Debugf("put point kind [%v]", pointKind)
	return s.PutPointKind(pointKind)
}

//GetStateByPartialCompositeKey（）函数根据给的部分字首，返回匹配的一个迭代器，不太理解
//是不是将一堆地址的上的积分合并的函数？
func (s *ChaincodeStore)MergeStateByPartialCompositeKey(objectType string, keys []string) (int64, error) {
	attributesTemp := []string{s.kind}
	keys = append(attributesTemp, keys...)

	logger.Debug("composite key ", objectType, keys)
	resultIterator, err := s.stub.GetStateByPartialCompositeKey(objectType, keys)
	if err != nil {
		return 0, err
	}
	defer resultIterator.Close()

	var assets int64
	assets = 0
	for i := 0; resultIterator.HasNext(); i++ {
		responseRange, err := resultIterator.Next()
		if err != nil {
			return 0, err
		}
		value, err := strconv.ParseInt(string(responseRange.GetValue()), 10, 64)
		if err != nil {
			return 0, err
		}
		assets += value

		err = s.stub.DelState(responseRange.Key)
		if err != nil {
			return 0, err
		}
	}

	return assets, nil
}



//注意，这个函数和上面的getPocket有重复读取的问题，混用要注意了,重复读取可能会有不一致的情况
func (s *ChaincodeStore) GetAllAssets(addr string) (int64, int64, error) {
	resultIterator, err := s.stub.GetStateByPartialCompositeKey(CompositeIndexName, []string{s.kind, addr})
	if err != nil {
		return 0, 0, err
	}
	defer resultIterator.Close()

	var assets int64
	assets = 0
	for i := 0; resultIterator.HasNext(); i++ {
		responseRange, err := resultIterator.Next()
		if err != nil {
			return 0, 0, err
		}
		value, err := strconv.ParseInt(string(responseRange.GetValue()), 10, 64)
		if err != nil {
			return 0, 0, err
		}
		assets += value
	}

	pocket, err := s.GetPocket(addr)
	if err != nil {
		return 0, 0, err
	}
	assets += pocket.GetBalance()

	return assets, pocket.GetBalance(), nil
}

//获取交易id，需要具体了解 stub.GetTxID()
func (s *ChaincodeStore)GetTxID() string {
	return s.stub.GetTxID()
}



func (s *ChaincodeStore)AddCompositeOutput(objectType string, attributes []string, value int64) error {
	attributesTemp := []string{s.kind}
	attributes = append(attributesTemp, attributes...)

	logger.Debug("composite key ", objectType, attributes)
	compositeKey, compositeErr := s.stub.CreateCompositeKey(objectType, attributes)
	if compositeErr != nil {
		return compositeErr
	}
	compositePutErr := s.stub.PutState(compositeKey, []byte(strconv.FormatInt(value, 10)))
	if compositePutErr != nil {
		return compositePutErr
	}
	return nil
}

//将交易费信息记录到区块链中	 	--ly
func (s *ChaincodeStore) PutTxFeeInfo(txFeeInfo *TxFeeInfo) error {
	logger.Debugf("put point kind [%v]", txFeeInfo)
	data, err := proto.Marshal(txFeeInfo)
	if err != nil {
		return err
	}
	return s.stub.PutState(s.kind + "_" + txFeeKey, data)
}

//获取交易费信息	--ly
func (s *ChaincodeStore) GetTxFeeInfo() (*TxFeeInfo, error) {
	data, err := s.stub.GetState(s.kind + "_" + txFeeKey)
	if err != nil {
		return nil, err
	}

	return ParseTxFeeInfo(data)
}








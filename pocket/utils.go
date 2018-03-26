package pocket

import (
	"github.com/hyperledger-coin/proto"
	"strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"crypto/sha256"
	"encoding/hex"
	"encoding/base64"
	"strconv"
)

// MakeChaincodeStore returns a store for storing keys in the state
func MakeChaincodeStore(stub shim.ChaincodeStubInterface, kind string) Store {
	store := &ChaincodeStore{}
	store.stub = stub
	store.kind = kind
	return store
}

func ParsePocket(data []byte) (*Pocket, error) {
	if data == nil || len(data) == 0 {
		return nil, ErrNoAccount
	}

	pocket := new(Pocket)
	if err := proto.Unmarshal(data, pocket); err != nil {
		return nil, err
	}

	return pocket, nil
}

func ParsePointInfo(data []byte) (*PointInfo, error) {
	if data == nil || len(data) == 0 {
		return nil, ErrKeyNoData
	}
	pointInfo := new(PointInfo)
	if err := proto.Unmarshal(data, pointInfo); err != nil {
		return nil, err
	}

	return pointInfo, nil
}

func ParsePointKind(data []byte) (*PointKind, error) {
	if data == nil || len(data) == 0 {
		return nil, ErrKeyNoData
	}

	pointKind := new(PointKind)
	if err := proto.Unmarshal(data, pointKind); err != nil {
		return nil, err
	}

	return pointKind, nil
}

func ParseTxMaps(data []byte) (*TXMap, error) {
	if data == nil || len(data) == 0 {
		return nil, ErrKeyNoData
	}

	txMaps := new(TXMap)
	if err := proto.Unmarshal(data, txMaps); err != nil {
		return nil, err
	}
	return txMaps, nil
}

func IsValidAddr(addr string, pubkey string) bool {
	logger.Debugf("addr is [%v], pubkey is [%v]", addr, pubkey)
	if strings.Contains(addr, "_") {
		return false
	}

	pubbyte, err := base64.StdEncoding.DecodeString(pubkey)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	public, err := hex.DecodeString(byteToHexString(pubbyte))
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	if !(addr == NewAddrFromPubkey([]byte(public), byte(version)).String()) {
		logger.Errorf("addr [%v] is not equal to addrFromPublicKey [%v]", addr, NewAddrFromPubkey([]byte(public), byte(version)).String())
		return false
	}
	return true
}

// TxHash generates the Hash for the transaction.
func TxHash(tx *TXMap_TX) string {
	txBytes, err := proto.Marshal(tx)
	if err != nil {
		return ""
	}

	fHash := sha256.Sum256(txBytes)
	lHash := sha256.Sum256(fHash[:])
	return hex.EncodeToString(lHash[:])
}
func ParseTxFeeInfo(data []byte) (*TxFeeInfo, error) {
	if data == nil || len(data) == 0 {
		return nil, ErrKeyNoData
	}

	txFeeInfo := new(TxFeeInfo)
	if err := proto.Unmarshal(data, txFeeInfo); err != nil {
		return nil, err
	}

	return txFeeInfo, nil
}

func byteToHexString(byteArray []byte) string {
	result := ""
	for i := 0; i < len(byteArray); i++ {
		hex := strconv.FormatInt(int64(byteArray[i]&0xFF), 16)
		if len(hex) == 1 {
			hex = "0" + hex
		}
		result += hex
	}
	return strings.ToUpper(result)
}
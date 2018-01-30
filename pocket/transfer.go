package pocket

import (
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/base64"
	"time"
)

func (t *PocketChaincode)transfer(store Store, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(ErrInvalidArgs.Error())
	}

	txDataBase64 := args[1]
	txData, err := base64.StdEncoding.DecodeString(txDataBase64)
	if err != nil {
		return shim.Error(err.Error())
	}
	txMaps, err := ParseTxMaps(txData)
	if err != nil {
		return shim.Error(err.Error())
	}

	//保证签名时间在两分钟以内
	if time.Now().UTC().Unix() - txMaps.Timestamp > 120 {
		return shim.Error(ErrTimeOut.Error())
	}

	txid := store.GetTxID()
	for i, tx := range txMaps.TxMap {
		if !VerifyTx(tx, store) {
			return shim.Error(ErrInvalidTX.Error())
		}

		//合并复合键
		var assets int64
		inputPocket, err := store.GetPocket(tx.GetInputAddr())
		if err != nil {
			return shim.Error(err.Error())
		}
		assets += inputPocket.GetBalance()

		//删除复合键
		mergeAssets, err := store.MergeStateByPartialCompositeKey(CompositeIndexName, []string{tx.GetInputAddr()})
		if err != nil {
			return shim.Error(err.Error())
		}
		assets += mergeAssets

		for j, output := range tx.GetOutput() {
			err := store.AddCompositeOutput(CompositeIndexName, []string{output.GetOutputAddr(), txid, string(i), string(j)}, output.GetOutputValue())
			if err != nil {
				return shim.Error(err.Error())
			}
			assets -= output.GetOutputValue()
		}

		inputPocket.Balance = assets
		err = store.PutPocket(inputPocket)
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	return shim.Success(nil)
}

func VerifyTx(tx *TXMap_TX, store Store) bool {
	assets, nounce, err := store.GetAllAssets(tx.GetInputAddr())
	if err != nil {
		return false
	}
	//nounce和总资产是否符合
	if assets <= tx.GetInputBalance() || nounce != tx.GetNounce() {
		return false
	}
	var outputValue int64
	for _, output := range tx.GetOutput() {
		outputValue += output.GetOutputValue()
	}
	//输入输出是否匹配
	if tx.GetInputBalance() < (outputValue + tx.GetFee()) {
		return false
	}

	//签名验证
	return VerifySign()
}

//TODO now the verfiy sign is null
func VerifySign() bool{
	return true
}
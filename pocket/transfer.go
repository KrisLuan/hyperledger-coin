package pocket

import (
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/base64"
	"time"
)

func (t *PocketChaincode)transfer(store Store, args []string) pb.Response {
	if len(args) != 2 {
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
	logger.Debugf("tx maps is [%v]", txMaps)
	//保证签名时间在两分钟以内
	if time.Now().UTC().Unix() - txMaps.Timestamp > 120 {
		return shim.Error(ErrTimeOut.Error())
	}

	txid := store.GetTxID()
	for i, tx := range txMaps.TxMap {
		logger.Debugf("tx [%v] is [%v]", i, tx)
		var assets int64
		inputPocket, err := store.GetPocket(tx.GetInputAddr())
		if err != nil {
			return shim.Error(err.Error())
		}
		assets += inputPocket.GetBalance()

		txFeeInfo, err := store.GetTxFeeInfo()
		if err != nil {
			return shim.Error(err.Error())
		}
		if !VerifyTx(tx, inputPocket.GetPubkey(), txFeeInfo, store) {
			return shim.Error(ErrInvalidTX.Error())
		}

		//合并复合键
		//删除复合键，不需要考虑在共识过程增加的那些复合键，只合并并删除当前的复合键
		logger.Debugf("merge state which addr is [%v]", tx.GetInputAddr())
		mergeAssets, err := store.MergeStateByPartialCompositeKey(CompositeIndexName, []string{tx.GetInputAddr()})
		if err != nil {
			return shim.Error(err.Error())
		}
		assets += mergeAssets

		for j, output := range tx.GetOutput() {
			logger.Debugf("add composite output [%v] [%v] [%v] [%v]", output.GetOutputAddr(), txid, string(i), string(j))
			err := store.AddCompositeOutput(CompositeIndexName, []string{output.GetOutputAddr(), txid, string(i), string(j)}, output.GetOutputValue())
			if err != nil {
				return shim.Error(err.Error())
			}
			assets -= output.GetOutputValue()
		}

		//fee
		logger.Debugf("fee: add composite output [%v] [%v] [%v] [0]",InitAddr, txid, string(i))
		err = store.AddCompositeOutput(CompositeIndexName, []string{txFeeInfo.GetTxFeeAddr(), txid, string(i), "0"}, tx.GetFee())
		if err != nil {
			return shim.Error(err.Error())
		}
		assets -= tx.GetFee()
		if assets < 0 {
			return shim.Error(ErrAccountNotEnoughBalance.Error())
		}

		//change
		logger.Debugf("change pocket assets is [%v]", assets)
		inputPocket.Balance = assets
		err = store.PutPocket(inputPocket)
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	return shim.Success(nil)
}

func VerifyTx(tx *TXMap_TX, publicKey string, txFeeInfo *TxFeeInfo, store Store) bool {
	assets, nounce, err := store.GetAllAssets(tx.GetInputAddr())
	if err != nil {
		return false
	}
	//nounce和总资产是否符合
	logger.Debugf("assets [%v] input balance [%v] nounce [%v] input nounce [%v]", assets, tx.GetInputBalance(), nounce, tx.GetNounce())
	if assets <= tx.GetInputBalance() || nounce != tx.GetNounce() {
		return false
	}
	var outputValue int64
	//验证是否有输出
	if len(tx.GetOutput()) < 1 {
		return false
	}
	for _, output := range tx.GetOutput() {
		outputValue += output.GetOutputValue()
	}
	//输入输出是否匹配
	logger.Debugf("input balance [%v] output and fee [%v]", tx.GetInputBalance(), outputValue + tx.GetFee())
	if tx.GetInputBalance() < (outputValue + tx.GetFee()) {
		return false
	}

	//验证交易费是否大于0.2%
	if !(tx.GetFee()*1000/outputValue >= txFeeInfo.GetRatio()) {
		logger.Errorf(ErrNotEnoughFee.Error())
		return false
	}

	//签名验证
	return VerifySign(tx, publicKey)
}

package pocket

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"os/exec"
	"strings"
	"github.com/hyperledger/fabric/hyperledger-coin/btcd/btcec"
)

func signedmessage(tx *TXMap_TX) ([]byte, error) {
	txsigned := new(TXMap_TX)
	txsigned = tx

	script := tx.Script

	//	fmt.Println("txsigned txin", txsigned.Txin)
	//	fmt.Println("txsigned txout", txsigned.Txout)

	tx.Script = ""

	txhash := TxHash(txsigned)
	signmessage, err := hex.DecodeString(txhash)
	if err != nil {
		logger.Errorf("hex.DecodeString error : %v", err)
		return signmessage, err
	}

	tx.Script = script

	return signmessage, nil
}

//TODO now the verfiy sign is null
func VerifySign(tx *TXMap_TX, publicKey string) bool{
	message, err := signedmessage(tx)
	if err != nil {
		return false
	}

	out := verfiyEcdsa(publicKey, message, tx.GetScript())
	return out
}


func verfiyEcdsa(pubScr string, message []byte, sigScript string) bool {

	sigB, _ := base64.StdEncoding.DecodeString(sigScript)
	signature0, _ := btcec.ParseSignature(sigB, btcec.S256())

	//fmt.Printf("r:%s, s:%s\n",signature0.R.String(), signature0.S.String())

	//message0 := "test message"

	pubKeyB, _ := base64.StdEncoding.DecodeString(pubScr)

	pubkey0, _ := btcec.ParsePubKey(pubKeyB, btcec.S256())
	verified0 := signature0.Verify(message, pubkey0)

	return verified0

	//fmt.Printf("--Signature Verified? %v\n", verified0)
}


package pocket

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"

	//"github.com/hyperledger/fabric/coinbase/secp256k1"
	"strings"
	strconv "strconv"
)

func TestNewAddrFromPubkey(t *testing.T) {

	privatekey, err := hex.DecodeString(InitPrikey)
	pubstring := InitPubkey

	pubbyte, err := base64.StdEncoding.DecodeString(pubstring)
	fmt.Println("hex", byteToHexString(pubbyte))
	pubkey, err := hex.DecodeString(byteToHexString(pubbyte))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pubkey, privatekey)

	var version int
	version = 1

	addr := NewAddrFromPubkey(pubkey, byte(version))
	//fmt.Println(addr)
	fmt.Println(addr.String())
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
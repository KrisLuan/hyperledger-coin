package pocket

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestNewAddrFromPubkey(t *testing.T) {

	privatekey, err := hex.DecodeString(InitPrikey)
	pubstring := "03d05c8dd917f038c383788e0f7d4963aa98dcdb6fd7ace659259896f5fedc818c"

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

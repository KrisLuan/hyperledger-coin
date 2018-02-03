package pocket

import (
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *PocketChaincode)registerAccount(store Store, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error(ErrInvalidArgs.Error())
	}
	logger.Debugf("register account which kind [%v], addr [%v], public key[%v]", args[0], args[1], args[2])

	addr := args[1]
	pubkey := args[2]
	if IsValidAddr(addr, pubkey) {
		return shim.Error(ErrAddrWithPubkey.Error())
	}
	if tmpPocket, err := store.GetPocket(addr); err == nil && tmpPocket != nil {
		return shim.Error(ErrAlreadyRegistered.Error())
	}

	pocket := &Pocket{
		Addr:		addr,
		Balance:	0,
		Pubkey:		pubkey,
	}
	if err := store.PutPocket(pocket); err != nil {
		return shim.Error(err.Error())
	}

	if err := store.ModifyPointInfo(1, 0, 0); err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
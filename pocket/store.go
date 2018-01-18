package pocket

import "github.com/hyperledger/fabric/core/chaincode/shim"

type Store interface {
	InitPocket() error
}

// Store struct uses a chaincode stub for state access
type ChaincodeStore struct {
	stub shim.ChaincodeStubInterface
}

// MakeChaincodeStore returns a store for storing keys in the state
func MakeChaincodeStore(stub shim.ChaincodeStubInterface) Store {
	store := &ChaincodeStore{}
	store.stub = stub
	return store
}

func (s *ChaincodeStore)InitPocket()  {
	
}
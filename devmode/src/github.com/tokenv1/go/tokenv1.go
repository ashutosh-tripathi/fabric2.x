package main

import (
	"fmt"

	"github.com/tokenv1/go/logger"

	"strings"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

type TokenChaincode struct {
}

var log = logger.NewLogger("DEBUG")

func (token *TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	log.DEBUG("This is init")

	return shim.Success(nil)
}

func (token *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	log.DEBUG("This is invoke")
	signedProposal, err := stub.GetSignedProposal()
	if err != nil {
		fmt.Printf("transaction proposal %#v", signedProposal)
	}
	log.DEBUG("txid" + stub.GetTxID())
	t, err := stub.GetTxTimestamp()
	if err != nil {
		fmt.Printf("tx timestamp %v", &t)
	}

	// log.DEBUG("meta" + stub.GetMetadata())
	log.DEBUG("cha" + stub.GetChannelID())

	function, args := stub.GetFunctionAndParameters()
	log.DEBUG("function invoked is " + function + " with args" + strings.Join(args, " "))
	s := "function invoked is " + function + " with args" + strings.Join(args, " ")
	return shim.Success([]byte(s))
}
func main() {
	log.DEBUG("inside main func")
	err := shim.Start(new(TokenChaincode))

	if err != nil {
		log.DEBUG("error starting chaincode" + err.Error())
	}

}

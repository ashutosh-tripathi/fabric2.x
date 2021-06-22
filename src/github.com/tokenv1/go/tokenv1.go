package main

import (
	"fmt"
	"strconv"

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
	stub.PutState("myToken", []byte("200"))
	return shim.Success([]byte("true"))
}

func (token *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	log.DEBUG("This is invoke")

	// V3
	// Extract the information from proposal
	PrintSignedProposalInfo(stub)

	// V3
	// Will receieve empty map unless client set the transient data in Tx Proposal
	// transientData, _ := stub.GetTransient()
	// fmt.Println("GetTransient() =>", transientData)

	PrintCreatorInfo(stub)

	function, args := stub.GetFunctionAndParameters()
	log.DEBUG("function invoked is " + function + " with args" + strings.Join(args, " "))
	s := "function invoked is " + function + " with args" + strings.Join(args, " ")
	if function == "" {
		log.DEBUG("inside function nil")
		function = args[0]
	}
	switch function {
	case "set":
		return setToken(stub)
		break
	case "get":
		return getToken(stub)
		break
	case "del":
		return deleteToken(stub)
		break
	}
	return shim.Success([]byte(s))
}

func setToken(stub shim.ChaincodeStubInterface) peer.Response {

	fmt.Println("in setToken")
	value, err := stub.GetState("myToken")
	if err != nil {
		return shim.Error(err.Error())
	}

	intvalue, err := strconv.Atoi(string(value))
	if err != nil {
		// May also return sh.Error
		return shim.Success([]byte("false"))
	}
	intvalue += 10
	stub.PutState("myToken", []byte(strconv.Itoa(intvalue)))
	return shim.Success([]byte(strconv.Itoa(intvalue)))

}

func getToken(stub shim.ChaincodeStubInterface) peer.Response {

	fmt.Println("in getToken")
	value, err := stub.GetState("myToken")
	if err != nil {
		fmt.Println("Get Failed!!! ", err.Error())

		return shim.Error("Get Failed!! " + err.Error() + "!!!")
	}
	var myToken string
	if value == nil {
		// Return value -1 is to indicate to caller that MyToken
		// Does NOT exist in state data
		myToken = "-1"

	} else {

		myToken = "MyToken=" + string(value)

	}

	return shim.Success([]byte(myToken))

}

func deleteToken(stub shim.ChaincodeStubInterface) peer.Response {

	_, args := stub.GetFunctionAndParameters()

	val, err := stub.GetState(args[0])

	if err != nil {
		fmt.Println("error encountered in get" + err.Error())

		return shim.Error("error encountered in get" + err.Error())
	}

	if val == nil {
		fmt.Println("value is nil" + err.Error())

		return shim.Error("value is nil" + err.Error())
	} else {
		err = stub.DelState(args[0])
		if err == nil {
			return shim.Success([]byte("deleted value"))
		} else {
			return shim.Error("error while deleting" + err.Error())
		}
	}

}

func main() {
	log.DEBUG("inside main func")
	err := shim.Start(new(TokenChaincode))

	if err != nil {
		log.DEBUG("error starting chaincode" + err.Error())
	}

}

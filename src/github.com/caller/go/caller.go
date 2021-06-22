package main

import (
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-chaincode-go/shim"

	peer "github.com/hyperledger/fabric-protos-go/peer"
)

type CallerChaincode struct {
}

func (caller *CallerChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Init executed.")
	// Return success
	return shim.Success([]byte("Init Done."))
}

func (caller *CallerChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	function, args := stub.GetFunctionAndParameters()
	fmt.Println("function invoked is " + function + " with args" + strings.Join(args, " "))

	if function == "setCaller" {
		response := stub.InvokeChaincode(args[1], stub.GetArgs(), args[2])
		fmt.Println("Receieved SET response from 'token' : " + response.String())

		return response

	} else if function == "getCaller" {
		response := stub.InvokeChaincode(args[1], stub.GetArgs(), args[2])
		fmt.Println("Receieved SET response from 'token' : " + response.String())

		return response

	} else {
		return shim.Error("pass correct function")
	}

}

func main() {
	fmt.Println("inside main")
	err := shim.Start(new(CallerChaincode))
	if err != nil {
		fmt.Println("Error starting chaincode" + err.Error())
	}

}

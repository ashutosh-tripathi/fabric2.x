package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	peer "github.com/hyperledger/fabric-protos-go/peer"
	// KeyEndorsementPolicy interface for create key EP
	// https://godoc.org/github.com/hyperledger/fabric/core/chaincode/shim/ext/statebased#KeyEndorsementPolicy
	// "github.com/hyperledger/fabric/core/chaincode/shim/ext/statebased"
)

type TokenChaincode struct {
}

func (t *TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	fmt.Println("Init called")
	stub.PutState("UnProtectedToken", []byte("Inited-Unprotected"))
	stub.PutState("ProtectedToken", []byte("Inited-Protected"))
	ep := "AND('Org1MSP.peer','Org2MSP.peer')"
	err := stub.SetStateValidationParameter("ProtectedToken", []byte(ep))
	if err != nil {
		shim.Error("error" + err.Error())
	}
	return shim.Success([]byte("Iniated successfully"))
}

func (t *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	function, _ := stub.GetFunctionAndParameters()

	switch function {
	case "set":
		return setToken(stub)
	case "get":
		return getToken(stub)
	case "setEP":
		return setEPProtected(stub)
	case "getEP":
		return getEPProtected(stub)

	}
	return shim.Error("Pass valid function call")
}
func setToken(stub shim.ChaincodeStubInterface) peer.Response {
	_, args := stub.GetFunctionAndParameters()
	tokenName := args[0]
	tokenValue := args[1]
	err := stub.PutState(tokenName, []byte(tokenValue))
	if err != nil {
		return shim.Error("error in put token" + err.Error())
	}
	jsonString := "{ \"Token\":\"" + tokenName + "\","
	jsonString += "   \"Value\":\":" + tokenValue + "\"}"
	return shim.Success([]byte(jsonString))
}
func getToken(stub shim.ChaincodeStubInterface) peer.Response {
	_, args := stub.GetFunctionAndParameters()
	tokenName := args[0]
	valueProtected, err := stub.GetState(tokenName)
	if err != nil {
		shim.Error("error in getToken" + err.Error())
	}
	jsonString := "{\"Protected\":" + string(valueProtected) + "}"
	return shim.Success([]byte(jsonString))
}
func setEPProtected(stub shim.ChaincodeStubInterface) peer.Response {

	_, args := stub.GetFunctionAndParameters()
	ep := "AND("
	for i := 0; i < len(args)-1; i++ {
		if i > 0 {
			ep += ","
		}
		ep += "'" + args[i] + "'"
	}
	ep += ")"
	fmt.Println("constructed EP" + ep)
	err := stub.SetStateValidationParameter(args[len(args)-1], []byte(ep))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(ep))

}

func getEPProtected(stub shim.ChaincodeStubInterface) peer.Response {
	_, args := stub.GetFunctionAndParameters()
	ep, err := stub.GetStateValidationParameter(args[0])
	if err != nil {
		return shim.Error("error" + err.Error())
	}
	fmt.Println("fetched EP" + string(ep))
	return shim.Success(ep)
}

func main() {
	fmt.Println("in main")
	err := shim.Start(new(TokenChaincode))
	if err != nil {
		fmt.Println("in error")
	}
}

/**
 * SetEPProtected_1 Shows how to use
 **/
// func SetEPProtected_1(stub shim.ChaincodeStubInterface,args []string) peer.Response {

// 	ep, err  := statebased.NewStateEP(nil)
// 	for i:=0; i < len(args); i++ {
// 		ep.AddOrgs(statebased.RoleTypeMember , args[i])
// 	}

// 	epBytes,_ := ep.Policy()
// 	err = stub.SetStateValidationParameter("ProtectedToken", epBytes)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	return  shim.Success(epBytes)
// }

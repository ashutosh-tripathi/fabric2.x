package main

import (
	"fmt"

	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

type calccSmartContract struct {
}

func (c *calccSmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	value := 100
	stub.PutState("value", []byte(strconv.Itoa(value)))
	fmt.Printf("successfully initiated %s", strconv.Itoa(value))
	return shim.Success([]byte("successfully initiated" + strconv.Itoa(value)))
}
func (c *calccSmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	funct, args := stub.GetFunctionAndParameters()
	// funct := args[0]

	fmt.Println("passed func,args" + funct + " " + args[0])
	switch funct {
	case "calculate":
		return calculatevalue(stub)
		break
	}
	return shim.Success([]byte("pass valid function name"))
}
func calculatevalue(stub shim.ChaincodeStubInterface) peer.Response {
	_, args := stub.GetFunctionAndParameters()
	valuebytes, _ := stub.GetState("value")
	value, _ := strconv.Atoi(string(valuebytes))
	add, _ := strconv.Atoi(args[0])
	//fmt.Printf(" values %d %d %s", value, add, total)
	// fmt.Println("values" + string(valuebytes) + " " + args[0] + " " + string(add+value))
	total := strconv.Itoa(add + value)
	fmt.Printf(" values %d %d %s", value, add, total)
	return shim.Success([]byte("calculated value" + total))

}
func main() {
	fmt.Println("inside main")
	err := shim.Start(new(calccSmartContract))
	if err != nil {
		fmt.Println("error occured")
	}

}

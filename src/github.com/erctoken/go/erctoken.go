package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

type ERCTokenChaincode struct {
}
type ERCToken struct {
	Symbol      string `json:"symbol"`
	TotalSupply int    `json:"totalSupply"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
}

func (e *ERCTokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	function, args := stub.GetFunctionAndParameters()
	fmt.Println("function invoked is " + function + " with args" + strings.Join(args, " "))
	if len(args) != 4 {
		return shim.Error("incorrect args")
	}
	symbol := args[0]
	totalSupply, err := strconv.Atoi(args[1])
	description := args[2]
	creator := args[3]
	erctoken := ERCToken{Symbol: symbol, TotalSupply: totalSupply, Description: description, Creator: creator}
	fmt.Printf("erc token in %#v", erctoken)
	jsonERC20, _ := json.Marshal(erctoken)
	stub.PutState("token", []byte(jsonERC20))

	key := "owner." + creator
	err = stub.PutState(key, []byte(args[1]))
	if err != nil {
		return shim.Error("error in putting chaincode")
	}
	return shim.Success([]byte("successfully placed chaincode"))

}

func (e *ERCTokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	function, args := stub.GetFunctionAndParameters()
	fmt.Println("function invoked is " + function + " with args" + strings.Join(args, " "))
	switch function {
	case "totalSupply":
		return totalSupply(stub)
		break
	case "balanceOf":
		return balanceOf(stub)
		break
	case "transfer":
		return transfer(stub)
		break
	}
	return shim.Success([]byte("Invoke called"))
}

func totalSupply(stub shim.ChaincodeStubInterface) peer.Response {
	bytes, err := stub.GetState("token")
	if err != nil {
		return shim.Error(err.Error())
	}
	var erctoken ERCToken
	err = json.Unmarshal(bytes, &erctoken)
	fmt.Printf("fetched token %#v", erctoken)
	fmt.Printf("fetched balance %d", erctoken.TotalSupply)
	fmt.Printf("fetched token" + strconv.Itoa(erctoken.TotalSupply))
	return shim.Success([]byte("fetched token" + strconv.Itoa(erctoken.TotalSupply)))
}
func balanceOf(stub shim.ChaincodeStubInterface) peer.Response {
	_, args := stub.GetFunctionAndParameters()
	owner := args[0]
	bytes, err := stub.GetState("owner." + owner)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("fetched Balance%s", string(bytes))

	return shim.Success([]byte("fetched response" + string(bytes)))
}

func transfer(stub shim.ChaincodeStubInterface) peer.Response {

	_, args := stub.GetFunctionAndParameters()
	from := args[0]
	to := args[1]
	amount, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error(err.Error())
	}
	if amount <= 0 {
		return shim.Error("Amount MUST be > 0!!!")
	}

	bytes, err := stub.GetState("owner." + from)
	fmt.Printf("fetched from balance %s", string(bytes))
	frombalance, _ := strconv.Atoi(string(bytes))
	if frombalance < amount {
		return shim.Error("Insufficient balance to cover transfer!!!")
	}
	// Reduce the tokens in from account
	frombalance = frombalance - amount
	bytes, err = stub.GetState("owner." + to)
	tobalance, _ := strconv.Atoi(string(bytes))
	tobalance = tobalance + amount

	err = stub.PutState("owner."+from, []byte(strconv.Itoa(frombalance)))
	err = stub.PutState("owner."+to, []byte(strconv.Itoa(tobalance)))

	return shim.Success([]byte("from owner." + from + " to owner." + to + " value " + strconv.Itoa(frombalance) + " to value " + strconv.Itoa(tobalance)))

}

func main() {
	fmt.Println("in main")
	err := shim.Start(new(ERCTokenChaincode))
	if err != nil {
		fmt.Println("error in starting chaincode")
	}
}

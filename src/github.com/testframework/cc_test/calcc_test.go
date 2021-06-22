package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
)

func TestCalcc(t *testing.T) {
	mstub := InitTest(t)
	if mstub != nil {
		t.Log("stub creation successfully")
	}
	ccArgs := SetupArgsArray("calculate", "10")
	response := mstub.MockInvoke("mockTxID", ccArgs)
	t.Logf("received response %s", response.Payload)
	// result, _ := strconv.ParseInt(string(response.Payload), 10, 64)

	// Log the received value
	// t.Logf("Add Received Result = %d", result)
}

func InitTest(t *testing.T) *shimtest.MockStub {
	// Create an instance of the MockStub
	// 2.0 changed shim to shimtest
	mockstub := shimtest.NewMockStub("CalcTestStub", new(calccSmartContract))
	response := mockstub.MockInit("mockTxID", nil)
	status := response.GetStatus()
	t.Logf("Received status = %d", status)
	if status != shim.OK {
		t.Log("status not okay")
	}
	return mockstub
}

func SetupArgsArray(funcname string, args ...string) [][]byte {

	fmt.Printf("args %s %d\n", funcname, len(args))
	ccargs := make([][]byte, len(args)+1)
	ccargs[0] = []byte(funcname)
	for i, arg := range args {
		ccargs[i+1] = []byte(arg)
	}
	for j, ai := range ccargs {
		fmt.Printf("%d %s", j, string(ai))
	}
	return ccargs
}

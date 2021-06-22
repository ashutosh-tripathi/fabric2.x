package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shimtest"
)

func TestERCToken(t *testing.T) {
	t.Log("testing init")
	mockstub := testInit(t)
	if mockstub == nil {
		t.Fail()
	}
	ccArgs := SetupArgsArray("totalSupply")
	response := mockstub.MockInvoke("mockTxID", ccArgs)
	t.Logf("received response %s", response.Payload)
	ccArgs = SetupArgsArray("balanceOf", "creator1")
	response = mockstub.MockInvoke("mockTxID", ccArgs)
	t.Logf("received response %s", response.Payload)
	ccArgs = SetupArgsArray("transfer", "creator1", "creator2", "10")
	response = mockstub.MockInvoke("mockTxID", ccArgs)
	t.Logf("received response %s", response.Payload)
	ccArgs = SetupArgsArray("balanceOf", "creator1")
	response = mockstub.MockInvoke("mockTxID", ccArgs)
	t.Logf("received response %s", response.Payload)
	ccArgs = SetupArgsArray("balanceOf", "creator2")
	response = mockstub.MockInvoke("mockTxID", ccArgs)
	t.Logf("received response %s", response.Payload)
}

func testInit(t *testing.T) *shimtest.MockStub {
	mockst := shimtest.NewMockStub("erctokenstub", new(ERCTokenChaincode))

	arg := SetupArgsArray("Init", "token1", "100", "test token", "creator1")

	response := mockst.MockInit("txId", arg)
	fmt.Printf("response %s", response.Payload)
	return mockst
}
func SetupArgsArray(funcname string, args ...string) [][]byte {

	fmt.Printf("args %s %d\n", funcname, len(args))
	ccargs := make([][]byte, len(args)+1)
	ccargs[0] = []byte(funcname)
	for i, arg := range args {
		ccargs[i+1] = []byte(arg)
	}
	// for j, ai := range ccargs {
	// 	fmt.Printf("%d %s", j, string(ai))
	// }
	return ccargs
}

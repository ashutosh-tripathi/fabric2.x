package main

import (
	"fmt"

	"time"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/common"
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

// var log = logger.NewLogger("DEBUG")

func PrintSignedProposalInfo(stub shim.ChaincodeStubInterface) {

	fmt.Println("SignedProposal Function called!!!")

	signedProposal, err := stub.GetSignedProposal()
	if err == nil {
		fmt.Printf("transaction proposal %#v", signedProposal)
	}
	log.DEBUG("txid" + stub.GetTxID())
	timest, err := stub.GetTxTimestamp()
	if err == nil {
		timestr := time.Unix(timest.GetSeconds(), 0)
		fmt.Printf("tx timestamp %s ", timestr)
	}

	// log.DEBUG("meta" + stub.GetMetadata())
	log.DEBUG("cha" + stub.GetChannelID())
	// Get the SignedProposal
	// SignedProposal has 2 parts
	// 1. ProposalBytes
	// 2. Signature
	// signedProposal, _ := stub.GetSignedProposal()
	data := signedProposal.GetProposalBytes()
	signBytes := signedProposal.GetSignature()
	fmt.Println("sign bytes" + string(signBytes))
	proposal := &peer.Proposal{}
	proto.Unmarshal(data, proposal)
	fmt.Printf("proposal %v", &proposal)

	// Proposal has 2 parts
	// 1. Header
	// 2. Payload - the structure for this depends on the type in the ChannelHeader
	header := &common.Header{}
	proto.Unmarshal(proposal.GetHeader(), header)
	fmt.Printf("header %v", &header)
	// Header has 2 parts
	// 1. ChannelHeader
	// 2. SignatureHeader
	channelHeader := &common.ChannelHeader{}
	proto.Unmarshal(header.GetChannelHeader(), channelHeader)
	fmt.Printf("channel Header %v", &channelHeader)
	fmt.Println("channelHeader.GetType() => ", common.HeaderType(channelHeader.GetType()))
	fmt.Println("channelHeader.GetChannelId() => ", channelHeader.GetChannelId())

}

func PrintCreatorInfo(stub shim.ChaincodeStubInterface) {
	fmt.Println("PrintCreatorInfo() executed ")

	byteData, _ := stub.GetCreator()

	fmt.Println("PrintCreatorInfo => ", string(byteData))
}

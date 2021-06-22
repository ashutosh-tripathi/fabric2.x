export CORE_PEER_TLS_ENABLED=true
export ORDERER_CA=${PWD}/crypto-config/ordererOrganizations/orderer.com/orderers/orderer.orderer.com/msp/tlscacerts/tlsca.orderer.com-cert.pem
export PEER0_ORG1_CA=${PWD}/crypto-config/peerOrganizations/org1.com/peers/peer0.org1.com/tls/ca.crt
export PEER0_ORG2_CA=${PWD}/crypto-config/peerOrganizations/org2.com/peers/peer0.org2.com/tls/ca.crt
export FABRIC_CFG_PATH=${PWD}/crypto-config

export CHANNEL_NAME=testchannel

 setGlobalsForOrderer(){
     export CORE_PEER_LOCALMSPID="OrdererMSP"
     export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/crypto-config/ordererOrganizations/orderer.com/orderers/orderer.orderer.com/msp/tlscacerts/tlsca.orderer.com-cert.pem
     export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/ordererOrganizations/orderer.com/users/Admin@orderer.com/msp
    
 }

setGlobalsForPeer0Org1(){
    export CORE_PEER_LOCALMSPID="Org1MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/peerOrganizations/org1.com/users/Admin@org1.com/msp
    export CORE_PEER_ADDRESS=localhost:7051
}

setGlobalsForPeer1Org1(){
    export CORE_PEER_LOCALMSPID="Org1MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/peerOrganizations/org1.com/users/Admin@org1.com/msp
    export CORE_PEER_ADDRESS=localhost:8051
    
}

setGlobalsForPeer0Org2(){
    export CORE_PEER_LOCALMSPID="Org2MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/peerOrganizations/org2.com/users/Admin@org2.com/msp
    export CORE_PEER_ADDRESS=localhost:9051
    
}

setGlobalsForPeer1Org2(){
    export CORE_PEER_LOCALMSPID="Org2MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto-config/peerOrganizations/org2.com/users/Admin@org2.com/msp
    export CORE_PEER_ADDRESS=localhost:10051
    
}


#setGlobalsForOrderer
#setGlobalsForPeer0Org1
#setGlobalsForPeer0Org2
#setGlobalsForPeer1Org1
#setGlobalsForPeer1Org2

createChannelTransaction(){
    ./tools/configtxgen -outputCreateChannelTx=./${CHANNEL_NAME}.tx -profile=BasicChannel -channelID=$CHANNEL_NAME
}

createChannel(){
    
    setGlobalsForPeer0Org1
    
    ./tools/peer channel create -o localhost:7050 -c $CHANNEL_NAME \
    --ordererTLSHostnameOverride orderer \
    -f ./${CHANNEL_NAME}.tx --outputBlock ./${CHANNEL_NAME}.block \
    --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
}

joinChannel(){
    setGlobalsForPeer0Org1
    ./tools/peer channel join -b ./$CHANNEL_NAME.block
    
    setGlobalsForPeer1Org1
    ./tools/peer channel join -b ./$CHANNEL_NAME.block
    
    setGlobalsForPeer0Org2
    ./tools/peer channel join -b ./$CHANNEL_NAME.block
    
    setGlobalsForPeer1Org2
    ./tools/peer channel join -b ./$CHANNEL_NAME.block
    
}
createAnchorPeers(){
    ./tools/configtxgen -profile BasicChannel -configPath ./crypto-config -outputAnchorPeersUpdate ./Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
    ./tools/configtxgen -profile BasicChannel -configPath ./crypto-config -outputAnchorPeersUpdate ./Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP

}

updateAnchorPeers(){
    setGlobalsForPeer0Org1
    ./tools/peer channel update -o localhost:7050 --ordererTLSHostnameOverride orderer -c $CHANNEL_NAME -f ./${CORE_PEER_LOCALMSPID}anchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA

    setGlobalsForPeer0Org2
    ./tools/peer channel update -o localhost:7050 --ordererTLSHostnameOverride orderer -c $CHANNEL_NAME -f ./${CORE_PEER_LOCALMSPID}anchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA

}
#sleep 5s
#createAnchorPeers
#sleep 5s
#createChannelTransaction
#sleep 5s
#createChannel
#sleep 5s
#joinChannel
sleep 5s
updateAnchorPeers


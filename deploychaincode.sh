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

presetup() {
    echo Vendoring Go dependencies ...
    pushd ./src/github.com/fabcar/go
    GO111MODULE=on go mod vendor
    popd
    echo Finished vendoring Go dependencies
}
#presetup
CHANNEL_NAME="testchannel"
CC_RUNTIME_LANGUAGE=golang
VERSION="1"
CC_SRC_PATH="src/github.com/fabcar/go"
CC_NAME="fabcar"

packageChaincode() {
    rm -rf ${CC_NAME}.tar.gz
    setGlobalsForPeer0Org1
    ./tools/peer lifecycle chaincode package ${CC_NAME}.tar.gz \
        --path ${CC_SRC_PATH} --lang ${CC_RUNTIME_LANGUAGE} \
        --label ${CC_NAME}_${VERSION}
    echo "===================== Chaincode is packaged on peer0.org1 ===================== "
}
#packageChaincode

installChaincode() {
    setGlobalsForPeer0Org1
    ./tools/peer lifecycle chaincode install ${CC_NAME}.tar.gz
    echo "===================== Chaincode is installed on peer0.org1 ===================== "

     setGlobalsForPeer1Org1
     ./tools/peer lifecycle chaincode install ${CC_NAME}.tar.gz
     echo "===================== Chaincode is installed on peer1.org1 ===================== "

    setGlobalsForPeer0Org2
    ./tools/peer lifecycle chaincode install ${CC_NAME}.tar.gz
    echo "===================== Chaincode is installed on peer0.org2 ===================== "

    setGlobalsForPeer1Org2
     ./tools/peer lifecycle chaincode install ${CC_NAME}.tar.gz
     echo "===================== Chaincode is installed on peer1.org2 ===================== "
}

#installChaincode

queryInstalled() {
    setGlobalsForPeer1Org1
    ./tools/peer lifecycle chaincode queryinstalled >&log.txt
    cat log.txt
    PACKAGE_ID=$(sed -n "/${CC_NAME}_${VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    echo PackageID is ${PACKAGE_ID}
    echo "===================== Query installed successful on peer0.org1 on channel ===================== "
}

queryInstalled
approveForMyOrg1() {
    setGlobalsForPeer0Org1
     #set -x
    ./tools/peer lifecycle chaincode approveformyorg -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.orderer.com --tls \
         --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name ${CC_NAME} --version ${VERSION} \
        --init-required --package-id ${PACKAGE_ID} \
        --sequence ${VERSION}
     #set +x

    echo "===================== chaincode approved from org 1 ===================== "

}
#approveForMyOrg1

checkCommitReadyness() {
    setGlobalsForPeer0Org1
    ./tools/peer lifecycle chaincode checkcommitreadiness \
        --channelID $CHANNEL_NAME --name ${CC_NAME} --version ${VERSION} \
        --sequence ${VERSION} --output json --init-required
    echo "===================== checking commit readyness from org 1 ===================== "
}
#checkCommitReadyness
approveForMyOrg2() {
    setGlobalsForPeer0Org2
    # set -x
    ./tools/peer lifecycle chaincode approveformyorg -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.orderer.com --tls \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name ${CC_NAME} --version ${VERSION} \
        --init-required --package-id ${PACKAGE_ID} \
        --sequence ${VERSION}
     #set +x

    echo "===================== chaincode approved from org 1 ===================== "

}
#approveForMyOrg2

checkCommitReadyness() {
    setGlobalsForPeer0Org1
    ./tools/peer lifecycle chaincode checkcommitreadiness \
        --channelID $CHANNEL_NAME --name ${CC_NAME} --version ${VERSION} \
        --sequence ${VERSION} --output json --init-required
    echo "===================== checking commit readyness from org 1 ===================== "
}
#checkCommitReadyness

commitChaincodeDefinition() {
    setGlobalsForPeer0Org1
    ./tools/peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.orderer.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        --channelID $CHANNEL_NAME --name ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --version ${VERSION} --sequence ${VERSION} --init-required

}
#commitChaincodeDefinition
queryCommitted() {
    setGlobalsForPeer0Org1
    ./tools/peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name ${CC_NAME}

}

#queryCommitted

chaincodeInvokeInit() {
    setGlobalsForPeer0Org1
    ./tools/peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.orderer.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --isInit -c '{"Args":[]}'

}

#chaincodeInvokeInit

chaincodeInvoke() {

    setGlobalsForPeer0Org1


    ./tools/peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.orderer.com \
        --tls $CORE_PEER_TLS_ENABLED \
        --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
          -c '{"function": "initLedger","Args":[]}'
}
#chaincodeInvoke

chaincodeQuery() {
    setGlobalsForPeer0Org2
    ./tools/peer chaincode query -C $CHANNEL_NAME -n ${CC_NAME} -c '{"function": "queryCar","Args":["CAR0"]}'

}
chaincodeQuery











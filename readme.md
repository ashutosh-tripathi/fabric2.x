1. **Generating crypto material**: 

    ```
         cryptogen generate --config=<CONFIG-TEMPLATE-FILE> --output=<OUTPUT-FILE-LOCATION>
    ```
    **Extending crypto material**

     ``` shell 
         cryptogen extend  --input=<Crypto-Location>
         --config=<CONFIG-TEMPLATE-FILE-WIT H-ONLY-NEW-ORGS>
    ```
2. **Generating genesis block**
   ``` shell
       configtxgen -profile=<Profile- as-in-configtx.yml> -outputBlock=<address-of-genesis-block-file> -channelID=<channel-name>
    ```

    **Inspecting genesis block**
      ``` shell
      configtxgen -inspectBlock <block-file-name>   
      ```    
3. **Bring up the dockers**
     ```shell
     sudo docker compose up -d
     ```
4. **Create channel transaction and join peer to channel**
    
    *execute ./createChannel.sh*

5. **Deploy and instantiate chaincode**
   
   chaincode label format: ccname-version org-specific-version

    *execute ./deploychaincode.sh*

**TIPS: To debug chaincode we can launch peer in dev mode using command:**
     *peer node start --peer-chaincodedev=true* 
**-- Difference between address/listenaddress in peer configuration**
     *if the address parameter is not set, the peer will listen for all connection in listen address, which accepts all connections on open grpc connection. It will be a security risk for incoming connection from unreliable source.*
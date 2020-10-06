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
3. **Create channel transaction and join peer to channel**
    
    *execute ./createChannel.sh*

4. **Deploy and instantiate chaincode**
   
    *execute ./deploychaincode.sh*
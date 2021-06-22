**1. Fetch the config from peer using command**
```shell
peer channel fetch config <protobuf file location || generally:./config/cfg_block.pb>  -c <Channel Name> -o <Orderer Address>
```
**2. Convert the pb file to json using configtxlator**
```shell
configtxlator proto_decode --input <protobuf file> --type common.Block  > <config json>
```
**3. Make changes to the fetched json and save it as modified json**

**4. Convert original JSON with no envelope to PB**
```shell
configtxlator proto_encode --input <original json> --type common.Config --output <original pb file>
```

**5. Convert modified JSON with no envelope to PB**
```shell
configtxlator proto_encode --input <modified json> --type common.Config --output <modified pb>
```

**6. Compute the delta ie. difference between original and modified pb using given command**

```shell
configtxlator compute_update --channel_id <channel_name> --original <original pb file> --updated <updated pb file> --output <output file>
```

**7. Convert the Update block PB to JSON**
```shell
configtxlator proto_decode --input <output pb file from above>--type common.ConfigUpdate  > <output json file>
```

**8. Now wrap the Update config JSON with the envelope**
```shell
ENVELOPE='{"payload":{"header":{"channel_header":{"channel_id":"'$CHANNEL_ID'", "type":2}},"data":{"config_update":'$(cat $CFG_UPDATE_BLOCK_NO_ENVELOPE_JSON)'}}}'
echo $ENVELOPE | jq . > $CFG_UPDATE_BLOCK_WITH_ENVELOPE_JSON
```

**9. Convert the update block from JSON to PB**
```shell
configtxlator proto_encode --input $CFG_UPDATE_BLOCK_WITH_ENVELOPE_JSON --type common.Envelope --output $CFG_UPDATE_BLOCK_WITH_ENVELOPE_PB
```

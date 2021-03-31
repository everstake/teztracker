- [General WSS information](#general-wss-information)
  - [Live Subscribing/Unsubscribing to streams](#live-subscribingunsubscribing-to-streams)
    - [Subscribe to a stream](#subscribe-to-a-stream)
    - [Unsubscribe to a stream](#unsubscribe-to-a-stream)
- [Detailed Stream information](#detailed-stream-information)
  - [Operations Streams](#operations-streams)
  - [Blocks Streams](#blocks-streams)
  - [Info Streams](#info-streams)
  - [Account Streams](#account-streams)
  - [Assets Streams](#assets-streams)
  - [Mempool Streams](#mempool-streams)
  

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Web Socket Streams for TezTracker
# General WSS information
* The base endpoint is: **wss://api.teztracker.com/v2/{network}/ws**
* All symbols for streams are **lowercase**
* The websocket server will send a `ping frame` every 1 minute. If the websocket server does not receive a `pong frame` back from the connection within a 10 minute period, the connection will be disconnected. Unsolicited `pong frames` are allowed.
* In the response, if the result received is `{"event":"sys","data":{"msg":"hello","description":""}}` this means that connection is success.

## Live Subscribing/Unsubscribing to streams

* The following data can be sent through the websocket instance in order to subscribe/unsubscribe from streams. Examples can be seen below.

### Subscribe to a stream
* Request
  ```javascript
  {
     "type":"subscribe",
     "payload":{
        "channels":[
           "blocks"
        ]
    }
  }
  ```

* Response
  ```javascript
    {
       "event":"sys",
       "data":{
          "msg":"subscribed",
          "description":"blocks"
       }
    }
  ```

### Unsubscribe to a stream
* Request
  ```javascript
    {
       "type":"unsubscribe",
       "payload":{
          "channels":[
             "blocks"
          ]
       }
    }
  ```

* Response
  ```javascript
    {
       "event":"sys",
       "data":{
          "msg":"unsubscribed",
          "description":"blocks"
       }
    }
  ```

### Error Messages

`{"event":"sys","data":{"msg":"unknown_command","description":""}}`

# Detailed Stream information

## Operations Streams
The Operations Streams push raw operation information.

**Stream Name:** `operations, transactions, endorsements, delegations, originations, activations`

**Update Speed:** Real-time

**Payload:**
```javascript
{
   "event":"transactions",
   "data":{
      "event":"transactions",
      "data":{
         "amount":150713146,
         "blockHash":"BLKsPkYCJGRSubjznYH4SLe1qjohGzhR8PvZwf6dLrtmXW5qjRB",
         "blockLevel":1408576,
         "confirmations":0,
         "consumedGas":1427,
         "counter":3974366,
         "cycle":343,
         "destination":"tz1NzutYywx96MMGxieVVFm36ieyjpkXTqcF",
         "fee":1400,
         "gasLimit":10600,
         "kind":"transaction",
         "operationGroupHash":"onrzxfQjjhso1CaAk7FbzqMwue3BG2CzfyKWNs8bkWQzPaJmdRb",
         "operationId":42969152,
         "source":"tz1c2yBQNrofzKvKLouC9TqHELSnh64fm8dn",
         "status":"applied",
         "storageLimit":300,
         "timestamp":1617190023
      }
   }
}

```

## Blocks Streams
The Blocks Streams push raw block information.

**Stream Name:** `blocks`

**Update Speed:** Real-time

**Payload:**
```javascript
{
   "event":"blocks",
   "data":{
      "event":"blocks",
      "data":{
         "activeProposal":"PsFLorenaUUuikDWvMDr6fGBRG8kt3e3D3fHoXK1j1BFRxeSH4i",
         "baker":"tz3RB4aoyjov4KEVRbuhvQ1CKJgBJMWhaeB8",
         "bakerName":"Foundation Baker 3",
         "chainId":"NetXdQprcVkpaWU",
         "consumedGas":414916279,
         "context":"CoWNYLQsNsM6PPU6i8kGu3D8vC86jb26oPExHsnxuwgToxsMUj2u",
         "currentExpectedQuorum":5412,
         "delegations":1,
         "endorsements":23,
         "fees":69060,
         "fitness":"01,00000000000b7e83",
         "hash":"BMeh13bmxTuBMKBZjAWGwaGoVZR7BtQ7zyNejMKWshgfx7WByJJ",
         "level":1408643,
         "metaCycle":343,
         "metaCyclePosition":3714,
         "metaLevel":1408643,
         "metaLevelPosition":1408642,
         "metaVotingPeriod":44,
         "metaVotingPeriodPosition":3715,
         "number_of_operations":17,
         "operationsHash":"LLoaY7MfVrLfBHii1w8WzzT2WoxACxoo8eRKrWHUGo38Jx2ajnaDV",
         "periodKind":"testing",
         "predecessor":"BLiXSQhDhQZb84bvhb87QhX4fzYxJqcAcHnC1dCSWV9Ck6ed7Da",
         "priority":0,
         "proto":8,
         "protocol":"PtEdo2ZkT9oKpimTah6x2embF25oss54njMuPzkJTEi5RqfdZFA",
         "reveals":1,
         "reward":40000000,
         "signature":"sigqNhk2ZqrjPXEEaeXPgMpGSWWA4DgqPWgkbTJZieLNgzhzNegTtSdV8z3PD7B5jcBihCDbrsEthiRnD7hv9GgDZ8UsXsng",
         "timestamp":1617194043,
         "transactions":15,
         "validationPass":null,
         "volume":205188854
      }
   }
}
```

## Info Streams
The Info Streams push tezos info for 24 hours  .

**Stream Name:** `info_{currency}`

**Available currencies:** usd, eur, gbp, cny

**Update Speed:** Real-time

**Payload:**
```javascript
{
   "event":"info_usd",
   "data":{
      "event":"info_usd",
      "data":{
         "annual_yield":7.12,
         "blocks_in_cycle":4096,
         "circulating_supply":764873788.86517,
         "currency":"usd",
         "market_cap":3394247033,
         "price":4.45,
         "price_24h_change":-0.21165033,
         "staking_ratio":77.6306960147685,
         "volume_24h":276491386
      }
   }
}
```
## Account Streams
The Account Streams push created accounts.

**Stream Name:** `account_created_at, accounts, contracts`

**Update Speed:** Real-time
```javascript
{
   "event":"accounts",
   "data":{
      "event":"accounts",
      "data":{
         "accountId":"tz1dMNEYuD5J1FmfdVY11wsbvtubVWxpP7mQ",
         "balance":9998580,
         "blockId":"BLEmBYGrgPn5XMpWfJD95wZvWEEL8kTB1FpJ1wdomMxtGtrZoKa",
         "blockLevel":1408672,
         "counter":12237337,
         "createdAt":1617195797,
         "delegateSetable":null,
         "is_baker":false,
         "lastActive":1617195783,
         "manager":null,
         "revealed":false,
         "spendable":null
      }
   }
}
```

## Assets Streams
The Assets Streams push assets operations.

**Stream Name:** `asset_operations`

**Update Speed:** Real-time
```javascript
{
   "event":"asset_operations",
   "data":{
      "event":"asset_operations",
      "data":{
         "amount":0,
         "fee":14572,
         "from":"tz1gmajxhmKt22CbcsSA7WWXSATgbuVEmzTC",
         "gas_limit":141263,
         "operation_group_hash":"onitV2kqHU3oN8eLM9WJErogia98RMjfykNGb8EVrcsMHuhyvBE",
         "storage_limit":0,
         "timestamp":1617195963,
         "to":"KT1PWx2mnDueood7fEmfbBDKx1D9BAnnXitn",
         "type":"removeLiquidity"
      }
   }
}
```

## Mempool Streams
The Mempool Streams push mempool operations.

**Stream Name:** `mempool`

**Update Speed:** Real-time
```javascript
{
   "event":"mempool",
   "data":{
      "event":"mempool",
      "data":{
         "branch":"BLXiyEanSwm1HVT4tBgSr2usDxkfEGF7jzp8u71vy5ahSLFcpch",
         "contents":[
            {
               "kind":"transaction",
               "source":"tz1bDXD6nNSrebqmAnnKKwnX1QdePSMCj4MX",
               "fee":"20700",
               "counter":"267390",
               "gas_limit":"100000",
               "storage_limit":"300",
               "amount":"400000",
               "destination":"tz1dbRQEGe6Xa69UJp2fm8JUBSYmfJrYvvkL"
            }
         ],
         "protocol":"PtEdo2ZkT9oKpimTah6x2embF25oss54njMuPzkJTEi5RqfdZFA",
         "signature":"sigZEUX46VctHUGJKcNgJo6FX9YV9EnFyaw1jMdWoL9dQKEYDoYcZAAMH7Lm99r86G2g9PaApgeEePV1FgtDkWmAB87ne78G"
      }
   }
}
```
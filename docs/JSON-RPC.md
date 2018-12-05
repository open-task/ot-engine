# JSON-PRC 2.0

## GetPublished
列出指定`address`发布的所有的任务。

### Parameters

### Returns


### Example
```
// Request
curl -s -X POST --data '{"jsonrpc":"2.0","method":"GetPublished","params":["0xF562a7c51a158ae6E6170Ef7905af5d1cE43d24A", 5],"id":"11"}' 'http://localhost:8080/v1/' | jq .

// Result
{
  "id":"11",
  "jsonrpc": "2.0",
  "result": "0x0234c8a3397aab58" // 158972490234375000
}
```

## GetSolved


### Parameters

### Returns


### Example


## GetAccepted


### Parameters

### Returns


### Example


## GetRejected


### Parameters

### Returns


### Example
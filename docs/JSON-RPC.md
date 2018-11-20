# JSON-PRC 2.0

## GetPublished
列出指定`address`发布的所有的任务。

### Parameters

### Returns


### Example
```
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"GetPublished","params":["0xc94770007dda54cF92009BFF0dE90c06F603a09f", "latest"],"id":"11"}'

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
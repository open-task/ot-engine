# OpenTask Engine API

## Return Structure 返回值结构

### Mission 任务

Mission是任务结构，也是本API的顶层结构，对应着一个任务，含有以下字段：

- `mission`：即`MissionId`，任务ID
- `block`：区块高度
- `tx`：交易`Hash`
- `data`：任务描述
- `publisher`：发布者的地址
- `time`：发布的时间（链上时间，下同）
- `solutions`：与该任务关联的解决方案集合

### Solution 解决方案

Solution是解决方案的数据结构，对应着某次任务的一次解决方案（答案）提交尝试，含有以下字段：

- `solution`：即`SolutionId`，解决方案ID
- `block`：区块高度
- `tx`：交易`Hash`
- `data`：方案描述
- `solver`：方案提交者的地址
- `time`：发布的时间

## GetAllPublished

列出发布的所有任务。

### Parameters

- `offset`: 整数，分页起始位置
- `limit`: 整数，本页限定任务个数

### Returns

### Example
```
// Request
curl -s -X POST --data '{"jsonrpc":"2.0","method":"GetAllPublished","params":[0, 5],"id":"11"}' 'http://localhost:8080/v1/' | jq .

// Result
{
  "id": "11",
  "jsonrpc": "2.0",
  "result": [
    {
      "Block": 0,
      "Tx": "",
      "Mission": "1",
      "Reward": 100,
      "Publisher": "",
      "Solutions": null
    },
    {
      "Block": 0,
      "Tx": "",
      "Mission": "m1",
      "Reward": 101,
      "Publisher": "",
      "Solutions": [
        {
          "Block": 0,
          "Tx": "",
          "Solution": "m1s1",
          "Mission": "m1",
          "Data": "i solved m1",
          "Solver": "",
          "Status": "accept",
          "Process": {
            "Block": 0,
            "Tx": "",
            "Solution": "m1s1",
            "Time": "2018-11-16 15:27:23",
            "Status": "accept"
          }
        }
      ]
    },
    {
      "Block": 0,
      "Tx": "",
      "Mission": "m2",
      "Reward": 102,
      "Publisher": "",
      "Solutions": [
        {
          "Block": 0,
          "Tx": "",
          "Solution": "m2s1",
          "Mission": "m2",
          "Data": "i solved m2",
          "Solver": "",
          "Status": "reject",
          "Process": {
            "Block": 0,
            "Tx": "",
            "Solution": "m2s1",
            "Time": "2018-11-16 15:27:23",
            "Status": "reject"
          }
        },
        {
          "Block": 0,
          "Tx": "",
          "Solution": "m2s2",
          "Mission": "m2",
          "Data": "i solved m2 too",
          "Solver": "",
          "Status": "accept",
          "Process": {
            "Block": 0,
            "Tx": "",
            "Solution": "m2s2",
            "Time": "2018-11-16 15:27:23",
            "Status": "accept"
          }
        }
      ]
    },
    {
      "Block": 0,
      "Tx": "",
      "Mission": "m2",
      "Reward": 102,
      "Publisher": "",
      "Solutions": null
    },
    {
      "Block": 0,
      "Tx": "",
      "Mission": "m3",
      "Reward": 103,
      "Publisher": "",
      "Solutions": [
        {
          "Block": 0,
          "Tx": "",
          "Solution": "m3s1",
          "Mission": "m3",
          "Data": "i solved m3",
          "Solver": "",
          "Status": "reject",
          "Process": {
            "Block": 0,
            "Tx": "",
            "Solution": "m3s1",
            "Time": "2018-11-16 15:27:23",
            "Status": "reject"
          }
        },
        {
          "Block": 0,
          "Tx": "",
          "Solution": "m3s2",
          "Mission": "m3",
          "Data": "i solved m3 too",
          "Solver": "",
          "Status": "reject",
          "Process": {
            "Block": 0,
            "Tx": "",
            "Solution": "m3s2",
            "Time": "2018-11-16 15:27:23",
            "Status": "reject"
          }
        }
      ]
    }
  ]
}
```

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
  "id": "11",
  "jsonrpc": "2.0",
  "result": [
    {
      "Block": 0,
      "Tx": "",
      "Mission": "1",
      "Reward": 100,
      "Publisher": "",
      "Solutions": null
    },
    {
      "Block": 0,
      "Tx": "",
      "Mission": "m1",
      "Reward": 101,
      "Publisher": "",
      "Solutions": [
        {
          "Block": 0,
          "Tx": "",
          "Solution": "m1s1",
          "Mission": "m1",
          "Data": "i solved m1",
          "Solver": "",
          "Status": "accept",
          "Process": {
            "Block": 0,
            "Tx": "",
            "Solution": "m1s1",
            "Time": "2018-11-16 15:27:23",
            "Status": "accept"
          }
        }
      ]
    },
    {
      "Block": 0,
      "Tx": "",
      "Mission": "m2",
      "Reward": 102,
      "Publisher": "",
      "Solutions": [
        {
          "Block": 0,
          "Tx": "",
          "Solution": "m2s1",
          "Mission": "m2",
          "Data": "i solved m2",
          "Solver": "",
          "Status": "reject",
          "Process": {
            "Block": 0,
            "Tx": "",
            "Solution": "m2s1",
            "Time": "2018-11-16 15:27:23",
            "Status": "reject"
          }
        },
        {
          "Block": 0,
          "Tx": "",
          "Solution": "m2s2",
          "Mission": "m2",
          "Data": "i solved m2 too",
          "Solver": "",
          "Status": "accept",
          "Process": {
            "Block": 0,
            "Tx": "",
            "Solution": "m2s2",
            "Time": "2018-11-16 15:27:23",
            "Status": "accept"
          }
        }
      ]
    },
    {
      "Block": 0,
      "Tx": "",
      "Mission": "m2",
      "Reward": 102,
      "Publisher": "",
      "Solutions": null
    },
    {
      "Block": 0,
      "Tx": "",
      "Mission": "m3",
      "Reward": 103,
      "Publisher": "",
      "Solutions": [
        {
          "Block": 0,
          "Tx": "",
          "Solution": "m3s1",
          "Mission": "m3",
          "Data": "i solved m3",
          "Solver": "",
          "Status": "reject",
          "Process": {
            "Block": 0,
            "Tx": "",
            "Solution": "m3s1",
            "Time": "2018-11-16 15:27:23",
            "Status": "reject"
          }
        },
        {
          "Block": 0,
          "Tx": "",
          "Solution": "m3s2",
          "Mission": "m3",
          "Data": "i solved m3 too",
          "Solver": "",
          "Status": "reject",
          "Process": {
            "Block": 0,
            "Tx": "",
            "Solution": "m3s2",
            "Time": "2018-11-16 15:27:23",
            "Status": "reject"
          }
        }
      ]
    }
  ]
}
```

## GetMissionInfo

### Parameters

- `id`: `MissionId`，任务ID

### Returns

- `Mission`

### Example

**Request**

```bash
curl -s -X POST --data '{"jsonrpc":"2.0","method":"GetMissionInfo","params":["m1"],"id":"11"}' 'http://localhost:8080/v1/' | jq .
```

**Result**

```json
{
  "id": "11",
  "jsonrpc": "2.0",
  "result": {
    "block": 10483278,
    "tx": "0x3ad58a0b36360f4b0116dc3cae894161dc70e51e4601cfa0a69fe633ae7f44d8",
    "mission": "m1",
    "reward": 100,
    "data": "This is mission 1 published by alex.",
    "publisher": "0x2707732B64b6b10bC1658AE5eD39788C9D2479C5",
    "time": "2019-03-05 10:27:52 +0800 CST",
    "solutions": null
  }
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
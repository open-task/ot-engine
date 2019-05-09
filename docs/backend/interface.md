# Backend 接口

说明，冒号开头的表示参数，比如`/backend/v1/user/:user/skill`中的`:user`表示具体的用户，即以太坊地址。

URL|请求方法|含义|参数
|:---|:---:|---|---|
/backend/v1/user/:user/skill | GET  | 获得某用户的全部技能列表 |`user`: 用户(地址)|
/backend/v1/user/:user/skill | POST | 添加技能               |`user`: 用户(地址)<br>`skill`: 技能
/backend/v1/user/:user/skill/:skill|GET|获得技能信息|`user`: 用户(地址)<br>`skill`: 技能
/backend/v1/user/:user/skill/:skill|PUT|更新技能|`user`: 用户(地址)<br>`skill`: 技能
/backend/v1/user/:user/skill/:skill|PATCH|更新技能|`user`: 用户(地址)<br>`skill`: 技能
/backend/v1/user/:user/skill/:skill|DELETE|删除技能|`user`: 用户(地址)<br>`skill`: 技能
/backend/v1/skill/top|GET|获得全系统的技能列表|`limit`: 结果条数

## GET /backend/v1/user/:user/skill 获得某用户的全部技能列表

### 请求参数

- `user`: 用户公钥地址

### 返回参数

- `list`: 技能列表

### 示例

请求
```bash
curl -s -X GET '47.92.64.129/backend/v1/user/u1/skill'| jq .
```
返回
```json
[
  {
    "Id": 1,
    "user": "u1",
    "skill": "s1",
    "status": 0,
    "update_time": "2019-05-08 16:00:07"
  },
  {
    "Id": 2,
    "user": "u1",
    "skill": "s2",
    "status": 0,
    "update_time": "2019-05-08 16:28:22"
  }
]
```
## POST /backend/v1/user/:user/skill/:skill 添加技能

### 请求参数

### 返回参数

### 示例
请求
```bash
curl -s -X POST -H 'application/x-www-form-urlencoded' -d 'skill=s1' '127.0.0.1:8080/backend/v1/user/u1/skill' | jq .
```
返回
```json
{
  "Id": 1,
  "user": "u1",
  "skill": "s1",
  "status": 0
}
```
## GET /backend/v1/user/:user/skill/:skill 获得技能信息

### 请求参数

- `user`: 用户公钥地址
- `skill`: 技能ID
### 返回参数

### 示例
```bash
curl -s -X GET http://127.0.0.1:8080/backend/v1/user/u1/skill/1 | jq .
```

```json
{
  "Id": 1,
  "user": "u1",
  "skill": "s1",
  "status": 0,
  "update_time": "2019-05-08 17:35:12"
}
```

## PUT /backend/v1/user/:user/skill/:skill 更新技能

### 请求参数

### 返回参数

### 示例
```
curl -s -X PUT -H 'application/x-www-form-urlencoded' -d 'skill=s1' '47.92.64.129/backend/v1/user/u1/skill/s2' | jq .
```

## PATCH /backend/v1/user/:user/skill/:skill 更新技能

### 请求参数

### 返回参数

### 示例

## DELETE /backend/v1/user/:user/skill/:skill 删除技能

### 请求参数

### 返回参数

### 示例

## GET /backend/v1/skill/top 获得全系统的技能列表

### 请求参数

### 返回参数

### 示例

请求
```bash
curl -s -X GET '47.92.64.129/backend/v1/skill/top'| jq .
curl -s -X GET '47.92.64.129/backend/v1/skill/top?limit=30'| jq .
```
返回
```json
[
  {
    "Id": 1,
    "user": "u1",
    "skill": "\"C  \"",
    "update_time": "2019-05-09 11:31:53"
  }
]
```
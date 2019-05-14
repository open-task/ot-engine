# Backend 接口

说明，冒号开头的表示参数，比如`/backend/v1/user/:user/skill`中的`:user`表示具体的用户，即以太坊地址。

URL|请求方法|含义|参数
|:---|:---:|---|---|
/backend/v1/user/:address/skill/:skill|GET   |获得技能信息|`user`: 用户(地址)<br>`skill`: 技能
/backend/v1/user/:address/skill/:skill|PUT   |更新技能|`user`: 用户(地址)<br>`skill`: 技能
/backend/v1/user/:address/skill/:skill|PATCH |更新技能|`user`: 用户(地址)<br>`skill`: 技能
/backend/v1/user/:address/skill/:skill|POST  |更新技能|`user`: 用户(地址)<br>`skill`: 技能
/backend/v1/user/:address/skill/:skill|DELETE|删除技能|`user`: 用户(地址)<br>`skill`: 技能
/backend/v1/skill/top|GET|获得全系统的技能列表|`limit`: 结果条数


## GET /backend/v1/user/:user/skill/:skill 获得技能信息

### 请求参数

- `user`: 用户公钥地址
- `skill`: 技能ID
### 返回参数

### 示例
```bash
curl -s -X GET http://127.0.0.1:8080/backend/v1/user/u1/skill/1 | jq .

curl -s -X GET http://47.92.64.129/backend/v1/user/u1/skill/1 | jq .
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
请求
```bash
curl -s -X PUT -H 'application/x-www-form-urlencoded' -d 'skill=s1' '127.0.0.1:8080/backend/v1/user/u1/skill/s2' | jq .

curl -s -X PUT -H 'application/x-www-form-urlencoded' -d 'skill=s1' '47.92.64.129/backend/v1/user/u1/skill/s2' | jq .
```
返回
```json

```
## POST /backend/v1/user/:user/skill/:skill 更新技能

### 请求参数

### 返回参数

### 示例
请求
```bash
curl -s -X POST -H 'application/x-www-form-urlencoded' -d 'skill=s1' '127.0.0.1:8080/backend/v1/user/u1/skill/s2' | jq .

curl -s -X POST -H 'application/x-www-form-urlencoded' -d 'skill=s1' '47.92.64.129/backend/v1/user/u1/skill/s2' | jq .
```
返回
```json

```
## DELETE /backend/v1/user/:user/skill/:skill 删除技能

### 请求参数

### 返回参数

### 示例
请求
```bash
curl -s -X DELETE http://127.0.0.1:8080/backend/v1/user/u1/skill/s1 | jq .

curl -s -X DELETE http://47.92.64.129/backend/v1/user/u1/skill/s1 | jq .
```

```json
{
  "Id": 0,
  "user": "u1"
}
```
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
    "skill": "s1",
    "providers": 3
  },
  {
    "skill": "s3",
    "providers": 1
  }
]
```
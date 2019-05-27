# Backend 接口

## 返回结果结构体

### 用户信息
- `user`: 用户`json`
  - `id`: 用户ID
  - `address`: 地址
  - `email`:
  - `skills`:
  - `missions`:
  - `update_time`: 更新时间（用户信息更新时间，非任务和技能更新时间）

- `skill`: 技能`json`
  - `id`: 技能ID
  - `tag`: 技能标签，多用于直接展示，如"金融"、"区块链"等
  - `users`: 提供者列表（`user`)
  - `claim`: 声明数（声明拥有此项技能的人数）
  - `submit`: 提交数（提交过此类任务的人数）
  - `confirm`: 确认数（任务提交被确认/接受的数目）

## 接口目录

URL|请求方法|含义|请求参数|返回参数
|:---|:---|:---|:---|:---
/backend/v1/user_info   | GET  | 获得用户信息 | 无 | `skills`: 技能列表
/backend/v1/list_skills | GET  | 获得技能列表 | `limit`:条数 | `skills`: 技能列表
/backend/v1/get_users   | GET  | 获得用户列表 | `id`:技能id  | `users`: 用户列表
/backend/v1/add_skill   | POST | 添加技能    | `address`: 用户地址<br>`tag`: 技能标签 | `skill`: 技能
/backend/v1/user/:address/info | GET  | 查询用户信息 | `user_id`: 用户ID | `user`：用户信息
/backend/v1/user/:address/info | POST | 更新用户信息 | `user_id`: 用户ID<br>`address`:公钥地址<br>`email`:邮件地址<br> | `user`：用户信息
/backend/v1/user/:address/skill| GET  | 获得技能列表 |`skill`: 技能 |`skills`: 技能列表
/backend/v1/user/:address/skill| POST | 添加技能    |`skill`: 技能 | `skill`: 技能
/backend/v1/skill   | GET | 查询技能信息 |`tag`:技能<br>`limit`:条数| `skills`:`skill`列表|
/backend/v1/skill/:skill_id/user | GET | 查询技能提供者 |`id`:技能ID<br>`limit`:条数| `users`:`user`列表|
*/backend/v1/skill/update_skill(not implemented)* |POST | 更新技能列表 |`user`: 用户(地址)<br>`email`:邮件地址<br>`skill`: 技能列表|`skill`: 技能列表
*/backend/v1/skill/get_skill(not implemented)*    |POST | 查询技能列表 |`user`: 用户(地址)<br>`email`:邮件地址<br>`skill`: 技能列表|`skill`: 技能列表
*/backend/v1/skill/del_skill(not implemented)*    |POST | 删除技能信息 |`user`: 用户(地址)<br>`email`:邮件地址<br>`skill`: 技能列表|`skill`: 技能列表

## 获得技能列表

```bash
curl -s -X GET http://127.0.0.1:8080/backend/v1/list_skills | jq .
```

```json
[
  {
    "id": 1,
    "tag": "s1",
    "user_number": 2
  },
  {
    "id": 2,
    "tag": "s3",
    "user_number": 1
  },
  {
    "id": 3,
    "tag": "区块链",
    "user_number": 2
  },
  {
    "id": 4,
    "tag": "Golang",
    "user_number": 1
  }
]
```

## 获得用户列表


## 添加技能

```bash
curl -s -X POST \
  http://127.0.0.1:8080/backend/v1/add_skill \
  -H 'content-type: application/x-www-form-urlencoded' \
  -d 'email=user2@bountinet.com&address=0x2707732B64b6b10bC1658AE5eD39788C9D2479C5&tag=Golang' | jq .

curl -s -X POST \
  http://127.0.0.1:8080/backend/v1/add_skill \
  -H 'content-type: application/x-www-form-urlencoded' \
  -d 'address=0x1c635f4756ED1dD9Ed615dD0A0Ff10E3015cFa7b&tag=Golang' | jq .
```

## 查询用户信息

### 请求参数

- `address`: 用户地址

### 返回参数

- `user`: 用户信息

### 示例

请求

```bash
curl -s -X GET '127.0.0.1:8080/backend/v1/user/0x1c635f4756ED1dD9Ed615dD0A0Ff10E3015cFa7b/info' | jq .
```

返回
```json
{
  "id": 9,
  "address": "u1",
  "skills": [
    {
      "id": 5,
      "skill": "s1"
    },
    {
      "id": 6,
      "skill": "s2"
    },
    {
      "id": 8,
      "skill": "s3"
    }
  ]
}
```

## 更新用户信息

### 请求参数

- `user_id`: 用户ID
- `user`: 用户信息

### 返回参数

- `user`: 用户信息

### 示例

请求

```bash
curl -s -X POST \
  http://127.0.0.1:8080/backend/v1/user/0x1c635f4756ED1dD9Ed615dD0A0Ff10E3015cFa7b/info \
  -H 'content-type: application/x-www-form-urlencoded' \
  -d 'email=user111@bountinet.com' | jq .

curl -s -X POST \
  http://127.0.0.1:8080/backend/v1/user/9/info \
  -H 'content-type: application/json' \
  -d '{ "email": "user9@bountinet.com",
        "address": "0x1c635f4756ED1dD9Ed615dD0A0Ff10E3015cFa7b"
      }'  | jq .
```

返回
```json
{
  "id": 9,
  "address": "0x1c635f4756ED1dD9Ed615dD0A0Ff10E3015cFa7b",
  "email": "user9@bountinet.com"
}
```


## 查询技能列表(`GET /backend/v1/user/:user/skill`)

### 请求参数

- `skill`: 技能查询条件

### 返回参数

- `skills`: 技能列表

### 示例

请求
```bash
curl -s -X GET 'http://127.0.0.1:8080/backend/v1/user/9/skill'| jq .
```
返回
```json
[
  {
    "id": 1,
    "tag": "s1"
  },
  {
    "id": 2,
    "tag": "s3"
  },
  {
    "id": 3,
    "tag": "区块链"
  }
]
```

## 添加技能(`POST /backend/v1/user/:user/skill/:skill`)

### 请求参数

- `skill`: 技能

### 返回参数

- `skill`: 技能

### 示例
请求
```bash
 curl -s -X POST \
  'http://127.0.0.1:8080/backend/v1/user/0x1c635f4756ED1dD9Ed615dD0A0Ff10E3015cFa7b/skill' \
  -H 'content-type: application/json' \
  -d '{"tag": "区块链"}' | jq .
  
```
返回
```json
{
  "id": 3,
  "tag": "区块链"
}
```

## 查询技能信息

### 请求参数

- `limit`: 条数
- `tag`: 技能标签
- `claim`: 声明数（声明拥有此项技能的人数）
- `submit`: 提交数（提交过此类任务的人数）
- `confirm`: 确认数（任务提交被确认/接受的数目）

### 返回参数

- `skill`：技能

### 示例

请求

```bash
curl -s -X GET '127.0.0.1:8080/backend/v1/skill?tag=s1' | jq .
```

返回

```json

```


## 查询技能提供者

### 请求参数

- `skill_id`: 技能ID

### 返回参数

- `users`: `user`列表

### 示例


请求

```bash
curl -s -X GET 'http://127.0.0.1:8080/backend/v1/skill/1/user' | jq .
```

返回

```json
[
  {
    "id": 9,
    "address": "0x1c635f4756ED1dD9Ed615dD0A0Ff10E3015cFa7b",
    "email": "user111@bountinet.com",
    "update_time": "2019-05-14T18:33:47+08:00"
  }
]
```

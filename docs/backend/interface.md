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

## 接口目录

URL|请求方法|含义|请求参数|返回参数
|:---|:---|:---|:---|:---
/backend/v1/user/:user_id/info | GET  | 查询用户信息 | `user_id`: 用户ID | `user`：用户信息
/backend/v1/user/:user_id/info | POST | 更新用户信息 | `user_id`: 用户ID<br>`address`:公钥地址<br>`email`:邮件地址<br> | `user`：用户信息
*/backend/v1/skill/update_skill(not implemented)* |POST | 更新技能列表 |`user`: 用户(地址)<br>`email`:邮件地址<br>`skill`: 技能列表|`skill`: 技能列表
*/backend/v1/skill/get_skill(not implemented)*    |POST | 查询技能列表 |`user`: 用户(地址)<br>`email`:邮件地址<br>`skill`: 技能列表|`skill`: 技能列表
*/backend/v1/skill/del_skill(not implemented)*    |POST | 删除技能信息 |`user`: 用户(地址)<br>`email`:邮件地址<br>`skill`: 技能列表|`skill`: 技能列表

  
## 查询用户信息

### 请求参数

- `user_id`: 用户ID

### 返回参数

- `user`: 用户信息

### 示例

请求

```bash
curl -s -X GET '127.0.0.1:8080/backend/v1/user/9/info' | jq .
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
  http://127.0.0.1:8080/backend/v1/user/9/info \
  -H 'content-type: application/x-www-form-urlencoded' \
  -d 'email=user111@bountinet.com&address=0x1c635f4756ED1dD9Ed615dD0A0Ff10E3015cFa7b' | jq .

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
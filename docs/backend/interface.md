# Backend 接口

说明，冒号开头的表示参数，比如`/backend/v1/user/:user/skill`中的`:user`表示具体的用户，即以太坊地址。

URL|请求方法|含义|参数
|:---|:---:|---|---|
|/backend/v1/user/:user/skill|GET|获得某用户的全部技能列表|`user`: 用户(地址)|
/backend/v1/user/:user/skill/:skill|GET|获得技能信息|`user`: 用户(地址)<br>`skill`: 技能
/backend/v1/user/:user/skill/:skill|POST|添加技能|`user`: 用户(地址)<br>`skill`: 技能
/backend/v1/user/:user/skill/:skill|PUT|更新技能|`user`: 用户(地址)<br>`skill`: 技能
/backend/v1/user/:user/skill/:skill|PATCH|更新技能|`user`: 用户(地址)<br>`skill`: 技能
/backend/v1/user/:user/skill/:skill|DELETE|删除技能|`user`: 用户(地址)<br>`skill`: 技能
/backend/v1/skill/top|GET|获得全系统的技能列表|`limit`: 结果条数

## GET /backend/v1/user/:user/skill 获得某用户的全部技能列表

## GET /backend/v1/user/:user/skill/:skill 获得技能信息
## POST /backend/v1/user/:user/skill/:skill 添加技能
## PUT /backend/v1/user/:user/skill/:skill 更新技能
## PATCH /backend/v1/user/:user/skill/:skill 更新技能
## DELETE /backend/v1/user/:user/skill/:skill 删除技能
## GET /backend/v1/skill/top 获得全系统的技能列表
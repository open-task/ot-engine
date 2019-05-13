# Backend 接口

请求方法均为POST

URL|请求方法|含义|请求参数|返回参数
|:---|:---:|---|---|---
/backend/v1/skill/update_skill |POST | 更新技能列表 |`user`: 用户(地址)<br>`email`:邮件地址<br>`skill`: 技能列表|`skill`: 技能列表
/backend/v1/skill/get_skill    |POST | 查询技能列表 |`user`: 用户(地址)<br>`email`:邮件地址<br>`skill`: 技能列表|`skill`: 技能列表
/backend/v1/skill/del_skill    |POST | 删除技能信息 |`user`: 用户(地址)<br>`email`:邮件地址<br>`skill`: 技能列表|`skill`: 技能列表
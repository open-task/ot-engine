# DEBUG 常用命令

## de

### 从头重新处理日志

```bash
go run cmd/de.go -f 3146781 -t 3381889 -c config/config.json
```

### 检查日志是否写入

```bash
go run cmd/db_review.go
```
输出类似如下内容：
```bash
Top 5 row:
id: 1, tx: 0x37590c85618ddc87c3f148b3991fe2811243a01150fad96f0e9478b8fb30cb81, mission: 1, reward: 100
id: 2, tx: 0xfc725257d36b878829f449e991c07113933b81900acb805637216e8c3895ffe9, mission: m1, reward: 101
id: 3, tx: 0xec9943455aee564cd2678550a32b4229a8ce3aed2fcb9fd29cce3349344c8558, mission: m2, reward: 102
id: 4, tx: 0x06de61435dd6d5ae0ffec21acec16bf5b36cc8c00f9795690cfd1026a00b9588, mission: m2, reward: 102
id: 5, tx: 0x79c69531e1cc97a0c778c55751b1fb7cb6f741a225fca8d3ff3bbf9081a30dee, mission: m3, reward: 103
The reward of 1st row is: 100
```
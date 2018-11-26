# ot-engine

ot-engine是 OpenTask Engine，亦即任务管理系统。

## 开发环境配置

### 下载源码

```
git clone https://github.com/xyths/ot-engine.git
```

### 安装依赖

```
go get github.com/gin-gonic/gin
go get github.com/kardianos/govendor
go get github.com/go-sql-driver/mysql
go get github.com/ethereum/go-ethereum
go get github.com/kanocz/goginjsonrpc
```

### 准备数据库

```sql
create database ot_local;
CREATE USER 'engine'@'localhost' IDENTIFIED BY 'decopentask';
GRANT ALL ON ot_local.* TO 'engine'@'localhost';
```

#### 创建测试表
（可选）此表没别的用途，仅用于[测试数据库连接](#测试数据库连接)。
```sql
CREATE TABLE squareNum (number int PRIMARY KEY, squareNumber int);
```

### 测试数据库连接

```
$ go run dbtest.go
The square number of 13 is: 169
The square number of 1 is: 1
```

## 部署

```bash
go install ./...
```

### 启动服务

```
ot-engine
```

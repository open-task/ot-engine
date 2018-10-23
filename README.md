# ot-engine

ot-engine是 OpenTask Engine，亦即任务管理系统。

## 开发环境配置

### 下载源码

```
git clone https://github.com/xyths/ot-engine.git
```

### 安装依赖

```
go get -u github.com/gin-gonic/gin
go get -u github.com/kardianos/govendor
go get -u github.com/go-sql-driver/mysql
go get -u github.com/ethereum/go-ethereum
```

### 准备数据库

```sql
create database ot_local;
CREATE USER 'engine'@'localhost' IDENTIFIED BY 'decopentask';
GRANT ALL ON ot_local.* TO 'engine'@'localhost';
CREATE TABLE squareNum (number int PRIMARY KEY, squareNumber int);
```

### 测试数据库

```
$ go run dbtest.go
The square number of 13 is: 169
The square number of 1 is: 1
```

### 启动服务

```
go run main.go
```
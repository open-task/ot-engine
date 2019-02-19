# OpenTask Engine

`ot-engine` 即 `OpenTask Engine`，亦即`OpenTask`生态中的任务管理系统。设计稿详见[这里](https://dececo.github.io/docs/mission_system/design.html)。

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
go get gopkg.in/urfave/cli.v2
```

### 准备数据库

运行服务需要对应的数据库和表，请提前准备，以下仅为示例。

```sql
create database ot_local;
CREATE USER 'engine'@'localhost' IDENTIFIED BY 'decopentask';
GRANT ALL ON ot_local.* TO 'engine'@'localhost';
```

#### (可选）创建测试表
（可选）此表没别的用途，仅用于[测试数据库连接](#测试数据库连接)。
```sql
CREATE TABLE squareNum (number int PRIMARY KEY, squareNumber int);
```

### (可选）测试数据库连接

数据库新建好以后，可以用以下测试工具测试。
```
$ cd demo/test_db
$ go run main.go
The square number of 13 is: 169
The square number of 1 is: 1
```
如重复测试，需要清空表格。

## 部署

```bash
$ cd cmd/ote
$ go install .
```

会在`$GOPATH/bin`下生成ote文件。

可使用`ote --help`查看帮助，`ote --version`查看版本。

### 启动服务

#### 1. 监听链上事件
```bash
$ ote -c config.json listen
```
#### 2. 提供API
```bash
$ ote -c config.json serve
```
#### 3. 下载指定区块的日志
```bash
$ ote -c config.json download
```
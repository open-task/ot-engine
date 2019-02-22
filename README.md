# OpenTask Engine

`ot-engine` 即 `OpenTask Engine`，亦即`OpenTask`生态中的任务管理系统。设计稿详见[这里](https://dececo.github.io/docs/mission_system/design.html)。

## 开发环境配置

### 下载源码

```
git clone https://github.com/open-task/ot-engine.git
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

### 安装执行文件
```bash
$ go install -v ./...
```

会在`$GOPATH/bin`下生成`ote`的可执行文件。

可使用`ote --help`查看帮助，`ote --version`查看版本。

### 配置参考

- 智能合约地址
  - `Rinkeby`
    - `DET`: `0x04B703784D3d82B5d5E4C103d0bDb80169653f48`
    - `OpenTask`: `0xF562a7c51a158ae6E6170Ef7905af5d1cE43d24A`
    - `Node`: `wss://rinkeby.infura.io/ws/v3/e17969db9bc94e75a474b3d3c5257a75`
  - `Kovan`
    - `DET`: `0x6ffF60A882CE1Cd793dC14261Eec0f0d6A470E21`
    - `OpenTask`: `0x7b37CDa8c4633634E6dED334Bc033Bd9b2783184`
    - `Node`: `wss://kovan.infura.io/ws/v3/e17969db9bc94e75a474b3d3c5257a75`

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

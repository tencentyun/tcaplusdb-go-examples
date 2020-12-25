# tcaplusdb-go-examples
Table of Contents
=================

   * [tcaplusdb-go-examples](#tcaplusdb-go-examples)
   * [PROTOBUF说明](#protobuf\xE8\xAF\xB4\xE6\x98\x8E)
   * [入门](#\xE5\x85\xA5\xE9\x97\xA8)
      * [Docker环境准备](#docker\xE7\x8E\xAF\xE5\xA2\x83\xE5\x87\x86\xE5\xA4\x87)
      * [Go环境准备](#go\xE7\x8E\xAF\xE5\xA2\x83\xE5\x87\x86\xE5\xA4\x87)
      * [protoc工具准备](#protoc\xE5\xB7\xA5\xE5\x85\xB7\xE5\x87\x86\xE5\xA4\x87)
      * [TcaplusDB表准备](#tcaplusdb\xE8\xA1\xA8\xE5\x87\x86\xE5\xA4\x87)
         * [准备PROTO表示例文件](#\xE5\x87\x86\xE5\xA4\x87proto\xE8\xA1\xA8\xE7\xA4\xBA\xE4\xBE\x8B\xE6\x96\x87\xE4\xBB\xB6)
         * [TcaplusDB集群准备](#tcaplusdb\xE9\x9B\x86\xE7\xBE\xA4\xE5\x87\x86\xE5\xA4\x87)
         * [TcaplusDB表格组准备](#tcaplusdb\xE8\xA1\xA8\xE6\xA0\xBC\xE7\xBB\x84\xE5\x87\x86\xE5\xA4\x87)
         * [TcaplusDB表创建](#tcaplusdb\xE8\xA1\xA8\xE5\x88\x9B\xE5\xBB\xBA)
   * [示例代码](#\xE7\xA4\xBA\xE4\xBE\x8B\xE4\xBB\xA3\xE7\xA0\x81)
      * [示例代码下载](#\xE7\xA4\xBA\xE4\xBE\x8B\xE4\xBB\xA3\xE7\xA0\x81\xE4\xB8\x8B\xE8\xBD\xBD)
         * [示例代码目录结构](#\xE7\xA4\xBA\xE4\xBE\x8B\xE4\xBB\xA3\xE7\xA0\x81\xE7\x9B\xAE\xE5\xBD\x95\xE7\xBB\x93\xE6\x9E\x84)
            * [示例代码目录结构说明](#\xE7\xA4\xBA\xE4\xBE\x8B\xE4\xBB\xA3\xE7\xA0\x81\xE7\x9B\xAE\xE5\xBD\x95\xE7\xBB\x93\xE6\x9E\x84\xE8\xAF\xB4\xE6\x98\x8E)
            * [配置文件说明](#\xE9\x85\x8D\xE7\xBD\xAE\xE6\x96\x87\xE4\xBB\xB6\xE8\xAF\xB4\xE6\x98\x8E)
            * [公共代码说明](#\xE5\x85\xAC\xE5\x85\xB1\xE4\xBB\xA3\xE7\xA0\x81\xE8\xAF\xB4\xE6\x98\x8E)
            * [表定义说明](#\xE8\xA1\xA8\xE5\xAE\x9A\xE4\xB9\x89\xE8\xAF\xB4\xE6\x98\x8E)
            * [同步模式示例代码说明](#\xE5\x90\x8C\xE6\xAD\xA5\xE6\xA8\xA1\xE5\xBC\x8F\xE7\xA4\xBA\xE4\xBE\x8B\xE4\xBB\xA3\xE7\xA0\x81\xE8\xAF\xB4\xE6\x98\x8E)
            * [异步模式示例代码说明](#\xE5\xBC\x82\xE6\xAD\xA5\xE6\xA8\xA1\xE5\xBC\x8F\xE7\xA4\xBA\xE4\xBE\x8B\xE4\xBB\xA3\xE7\xA0\x81\xE8\xAF\xB4\xE6\x98\x8E)
         * [表接口代码生成](#\xE8\xA1\xA8\xE6\x8E\xA5\xE5\x8F\xA3\xE4\xBB\xA3\xE7\xA0\x81\xE7\x94\x9F\xE6\x88\x90)
         * [公共连接环境配置](#\xE5\x85\xAC\xE5\x85\xB1\xE8\xBF\x9E\xE6\x8E\xA5\xE7\x8E\xAF\xE5\xA2\x83\xE9\x85\x8D\xE7\xBD\xAE)
         * [编译代码](#\xE7\xBC\x96\xE8\xAF\x91\xE4\xBB\xA3\xE7\xA0\x81)
            * [同步模式编译](#\xE5\x90\x8C\xE6\xAD\xA5\xE6\xA8\xA1\xE5\xBC\x8F\xE7\xBC\x96\xE8\xAF\x91)
            * [异步模式编译](#\xE5\xBC\x82\xE6\xAD\xA5\xE6\xA8\xA1\xE5\xBC\x8F\xE7\xBC\x96\xE8\xAF\x91)
   * [接口示例](#\xE6\x8E\xA5\xE5\x8F\xA3\xE7\xA4\xBA\xE4\xBE\x8B)
      * [接口源代码](#\xE6\x8E\xA5\xE5\x8F\xA3\xE6\xBA\x90\xE4\xBB\xA3\xE7\xA0\x81)
      * [调用模式](#\xE8\xB0\x83\xE7\x94\xA8\xE6\xA8\xA1\xE5\xBC\x8F)
      * [接口源码](#\xE6\x8E\xA5\xE5\x8F\xA3\xE6\xBA\x90\xE7\xA0\x81)
      * [客户端接口](#\xE5\xAE\xA2\xE6\x88\xB7\xE7\xAB\xAF\xE6\x8E\xA5\xE5\x8F\xA3)
         * [初始化接口](#\xE5\x88\x9D\xE5\xA7\x8B\xE5\x8C\x96\xE6\x8E\xA5\xE5\x8F\xA3)
         * [连接接口](#\xE8\xBF\x9E\xE6\x8E\xA5\xE6\x8E\xA5\xE5\x8F\xA3)
      * [同步接口示例](#\xE5\x90\x8C\xE6\xAD\xA5\xE6\x8E\xA5\xE5\x8F\xA3\xE7\xA4\xBA\xE4\xBE\x8B)
         * [插入记录](#\xE6\x8F\x92\xE5\x85\xA5\xE8\xAE\xB0\xE5\xBD\x95)
         * [获取记录](#\xE8\x8E\xB7\xE5\x8F\x96\xE8\xAE\xB0\xE5\xBD\x95)
         * [替换记录](#\xE6\x9B\xBF\xE6\x8D\xA2\xE8\xAE\xB0\xE5\xBD\x95)
         * [更新记录](#\xE6\x9B\xB4\xE6\x96\xB0\xE8\xAE\xB0\xE5\xBD\x95)
         * [获取记录部分字段](#\xE8\x8E\xB7\xE5\x8F\x96\xE8\xAE\xB0\xE5\xBD\x95\xE9\x83\xA8\xE5\x88\x86\xE5\xAD\x97\xE6\xAE\xB5)
         * [更新记录部分字段](#\xE6\x9B\xB4\xE6\x96\xB0\xE8\xAE\xB0\xE5\xBD\x95\xE9\x83\xA8\xE5\x88\x86\xE5\xAD\x97\xE6\xAE\xB5)
         * [自增记录部分字段](#\xE8\x87\xAA\xE5\xA2\x9E\xE8\xAE\xB0\xE5\xBD\x95\xE9\x83\xA8\xE5\x88\x86\xE5\xAD\x97\xE6\xAE\xB5)
         * [删除记录](#\xE5\x88\xA0\xE9\x99\xA4\xE8\xAE\xB0\xE5\xBD\x95)
         * [批量获取记录](#\xE6\x89\xB9\xE9\x87\x8F\xE8\x8E\xB7\xE5\x8F\x96\xE8\xAE\xB0\xE5\xBD\x95)
         * [主键索引获取记录](#\xE4\xB8\xBB\xE9\x94\xAE\xE7\xB4\xA2\xE5\xBC\x95\xE8\x8E\xB7\xE5\x8F\x96\xE8\xAE\xB0\xE5\xBD\x95)
         * [二级索引获取记录](#\xE4\xBA\x8C\xE7\xBA\xA7\xE7\xB4\xA2\xE5\xBC\x95\xE8\x8E\xB7\xE5\x8F\x96\xE8\xAE\xB0\xE5\xBD\x95)
      * [异步接口示例](#\xE5\xBC\x82\xE6\xAD\xA5\xE6\x8E\xA5\xE5\x8F\xA3\xE7\xA4\xBA\xE4\xBE\x8B)
         * [遍历表](#\xE9\x81\x8D\xE5\x8E\x86\xE8\xA1\xA8)
            * [遍历状态](#\xE9\x81\x8D\xE5\x8E\x86\xE7\x8A\xB6\xE6\x80\x81)
            * [遍历条件设置（非必须）](#\xE9\x81\x8D\xE5\x8E\x86\xE6\x9D\xA1\xE4\xBB\xB6\xE8\xAE\xBE\xE7\xBD\xAE\xE9\x9D\x9E\xE5\xBF\x85\xE9\xA1\xBB)
            * [遍历示例](#\xE9\x81\x8D\xE5\x8E\x86\xE7\xA4\xBA\xE4\xBE\x8B)
   * [错误码](#\xE9\x94\x99\xE8\xAF\xAF\xE7\xA0\x81)


# PROTOBUF说明
PROTO表是基于PROTOBUF协议设计的TcaplusDB表，PROTOBUF协议是Google开源的通用RPC通信协议，用于TcaplusDB存储数据的序列化、反序列化等操作，具体关于PROTO表的定义说明可参考章节：表定义语言（PB，TDR）。PROTO表定义以protobuf格式来定义表结构，支持丰富的数据类型，请参考章节：数据类型(PB, TDR)。
#  入门
快速入手PROTOBUF协议表的开发涉及几个步骤，下面介绍如何基于TcalusDB本地Docker版环境，快速上手基于Golang进行PROTO表的增删查改操作。所有操作均在申请的开发测试机或云主机进行。
## Docker环境准备
在开始示例代码演示之前，需要提前准备好TcaplusDB本地Docker环境及tcapluscli工具，具体请参考资料：TcaplusDB入门-Docker部署篇.md。
Docker部署好后，对于命令行工具需要授权所有IP访问Docker环境，授权方式:
```
#access-id指定业务id, 2: tdr业务，3: pb业务，这里是pb业务所以默认为3
./tcapluscli privilege --endpoint-url=http://localhost --access-id=3--allow-all-ip
```

## Go环境准备
GO SDK示例依赖GO环境的部署，对于Centos系统可以直接安装通过:
```
yum install golang
```
建议版本：`1.13`以上。

##  protoc工具准备
对于protobuf定义文件，需要借助`protoc`工具来生成接口代码，如果要生成Go的接口代码，还需要借助`protoc-gen-go`插件来辅助。这里需要准备下载对应OS平台的这两个工具：
|工具名|下载|
|---|---|
|protoc|[Download](https://github.com/protocolbuffers/protobuf/releases)|
|protoc-gen-go|[Download](https://github.com/golang/protobuf)|

备注:
* protoc下载后，直接拷贝到/usr/bin目录
* protoc-gen-go下载后，进入相应目录，直接`go build -o protoc-gen-go main.go`，可得到二进制文件，把其拷贝到/usr/bin系统目录

## TcaplusDB表准备
### 准备PROTO表示例文件
这里以示例中的game_players.proto举例，表名: `game_players`, 表类型: `GENERIC`。文件具体内容如下：
```
syntax = "proto3";                      // Specify the version of the protocol buffers language


import "tcaplusservice.optionv1.proto"; // Use the public definitions of TcaplusDB by importing them.

message game_players {  // Define a TcaplusDB table with message

	// Specify the primary keys with the option tcaplusservice.tcaplus_primary_key
	// The primary key of a TcaplusDB table has a limit of 4 fields
    option(tcaplusservice.tcaplus_primary_key) = "player_id, player_name, player_email";

    // Specify the primary key indexes with the option tcaplusservice.tcaplus_index
    option(tcaplusservice.tcaplus_index) = "index_1(player_id, player_name)";
    option(tcaplusservice.tcaplus_index) = "index_2(player_id, player_email)";


    // Value Types supported by TcaplusDB
    // int32, int64, uint32, uint64, sint32, sint64, bool, fixed64, sfixed64, double, fixed32, sfixed32, float, string, bytes
    // Nested Types Message

    // primary key fields
    int64 player_id = 1;
    string player_name = 2;
    string player_email = 3;


    // Ordinary fields
    int32 game_server_id = 4;
    repeated string login_timestamp = 5;
    repeated string logout_timestamp = 6;
    bool is_online = 7;

    payment pay = 8;
}


message payment {

	int64 pay_id = 1;
	uint64 amount = 2;
    int64 method = 3;

}

```
将上述文件内容保存为`game_players.proto`。

### TcaplusDB集群准备
对于TcaplusDB,在创建表之前需要创建对应的表集群。对于Docker本地版，集群已经默认创建好一个供大家使用，所以不用再创建集群。对于PROTO集群，已经默认创建一个`pb_app`的业务，集群接入ID (AccessID) 默认为`3`。集群密码(AccessPassword)查看可打开TcaplusDB运维平台，打开方式：｀直接浏览器输入部署docker的机器ip即可，默认端口80`。默认登录方式：
```
用户名：　tcaplus
密码：　tcaplus
```
登录后，进入`业务管理->业务维护->选择业务名称，这里默认选pb_app业务,在对应行的右侧点击查看密码即可`。
### TcaplusDB表格组准备
TcaplusDB表在集群的基础上还依赖于表格组，相当于游戏里的逻辑分区，使用工具创建表格组命令如下：
```
#查看表格组帮助命令
./tcapluscli tablegroup -h

#创建一个表格组，id指定为2，　endpoint-url为上面docker暴露的80端口，access-id为集群接入ID(业务ID，3表示PROTO集群), 用于docker环境连接使用, group name由字母、数字和下划线组成
./tcapluscli tablegroup create --endpoint-url=http://localhost --access-id=3 --group-id=2　--group-name=zone_2
```

### TcaplusDB表创建
现在正式进入表创建环节，在上述表格组基础上创建一个PROTO表，执行创建表命令，如下所示：
```
#查看表创建命令提示帮助
./tcapluscli table -h

#创建一个表, 指定endpoint-url, 表格组id: group-id, 表类型: PROTO, 表定义文件: game_players.proto, 放当前路径
./tcapluscli table create  create --endpoint-url=http://localhost --access-id=3 --group-id=2 --schema-type=PROTO --schema-file=game_players.proto
```

# 示例代码
以Golang示例代码为例，介绍如何使用PROTOBUF接口进行TcaplusDB表数据操作，这里主要介绍Generic类型表操作。GO示例代码以`go mod`方式进行编译，GO版本以`1.15`举例.

## 示例代码下载
目前示例代码放在github路径，直接下载即可。
```
git clone https://github.com/tencentyun/tcaplusdb-go-examples.git
```
### 示例代码目录结构
```
[root@VM-48-13-centos tcaplusdb-go-examples]# tree
.
|-- pb
|   |-- async
|   |   |-- batchget.go
|   |   |-- delete.go
|   |   |-- fieldget.go
|   |   |-- fieldincrease.go
|   |   |-- fieldupdate.go
|   |   |-- getbypartkey.go
|   |   |-- get.go
|   |   |-- indexquery.go
|   |   |-- insert.go
|   |   |-- replace.go
|   |   |-- traverse.go
|   |   `-- update.go
|   |-- cfg
|   |   |-- api_cfg.go
|   |   |-- api_cfg.xml
|   |   `-- logconf.xml
|   |-- go.mod
|   |-- go.sum
|   |-- sync
|   |   |-- example.go
|   |   |-- Makefile
|   |   `-- test.go
|   |-- table
|   |   |-- game_players.proto
|   |   |-- tcaplusservice
|   |   |   |-- game_players.pb.go
|   |   |   `-- tcaplusservice.optionv1.pb.go
|   |   `-- tcaplusservice.optionv1.proto
|   `-- tools
|       `-- tools.go
`-- README.md

```
#### 示例代码目录结构说明
* __async__: 异步调用模式示例代码，每个接口一个示例文件
* __cfg__: 配置目录，放置tcaplusdb连接配置信息文件和日志配置文件，在跑示例前需要提前配置api_cfg.xml文件
* __sync__: 同步调用模式示例代码，包含两个文件，一个example.go放置所有接口示例代码，一个test.go压测读写接口示例
* __table__: 放置tcaplusdb表相关定义文件及生成的接口定义代码(protoc,protoc-gen-go生成)
* __tools__: 公共目录，放置一些公共代码，如客户端初始化、接收响应、结构体转换等代码

#### 配置文件说明
在cfg目录下存放了两个配置文件：api_cfg.xml, logconf.xml。主要是异步调用模式代码需要使用，同步模式这些配置直接在代码中。
* __api_cfg.xml__: 放置TcaplusDB集群访问配置信息，在跑示例代码之前需要提前配置好云上环境信息或本地docker环境信息。异步模式使用
* __logconf.xml__: 日志配置文件，默认ERROR级别，可配置DEBUG | INFO | WARN等级别，如果是要压测，建议配置为ERROR级别，提高压测性能
* __api_cfg.go__: 解析配置文件代码

#### 公共代码说明
在tools目录下，存放一些公共代码，如初始化连接客户端、异步接收响应、JSON转换等代码

#### 表定义说明
在table目录下存放相关表的接口定义文件和生成的接口定义代码。
|文件|文件说明|
|---|---|
|game_players.proto|示例表定义文件|
|tcaplusservice.optionv1.proto|tcaplusdb pb协议定义文件|
|tcaplusservice|用protoc生成的表定义接口代码目录|
|game_players.pb.go|生成的表定义接口代码|
|tcaplusservice.optionv1.pb.go|生成的pb协议接口代码|

#### 同步模式示例代码说明
在sync目录下存放同步调用模式接口代码。
|文件|文件说明|
|---|---|
|example.go| 示例代码主文件,包含CRUD接口|
|test.go|示例压测文件，压测读写接口，开协程并发压测|
|Makefile|编译文件，直接执行make可编译得到示例二进制文件|
|logconf.xml| 日志配置文件，默认级别ERROR, 如果需要更详细的，可开DEBUG,　如果压测的话用ERROR既可，避免性能损耗|

#### 异步模式示例代码说明
在async目录下存放异步调用模式接口代码。
|文件|文件说明|
|---|---|
|insert.go|插入记录|
|get.go|获取记录|
|delete.go|删除记录|
|update.go|更新记录|
|replace.go|替换记录|
|batchget.go|批量获取记录|
|fieldget.go|获取部分字段|
|fieldupdate.go|更新部分字段|
|fieldincrease.go|自增部分字段|
|getbypartkey.go|根据主键索引字段获取记录|
|indexquery.go|二级索引获取记录|
|traverse.go|遍历记录|

### 表接口代码生成
如果不想用示例代码中的示例表，参照game_players.proto定义好后，可以用如下命令生成：
```
cd table
mkdir tcaplusservice
#生成pb协议接口代码
protoc --go_out=./tcaplusservice tcaplusservice.optionv1.proto
#生成表定义接口代码
protoc --go_out=./tcaplusservice mytable.proto
```
备注：
* 需要同时安装protoc-gen-go插件才行
* 需要在proto文件中指定package, 如默认的package tcaplusservice

### 公共连接环境配置
下载示例代码后，异步模式代码找到cfg目录的`api_cfg.xml`进行配置，同步模式代码直接到sync下example.go文件，并修改如下参数，改成业务方申请的tcaplusdb连接信息或本地部署的tcaplusdb信息;
```
　　//集群访问地址，本地docker版：配置docker部署机器IP, 端口默认:9999, 腾讯云线上环境配置为连接地址IP和端口
        DirUrl = "tcp://x.x.x.x:xxxx"
        //集群接入ID, 默认为3，本地docker版：直接填3，云上版本：根据实际集群接入ID填写
        AppId = 3
        //集群访问密码，本地docker版：参考集群准备阶段获取集群密码步骤，需要借助tcaplusdb web运维平台查看; 云上版本：根据实际申请集群详情页查看
        Signature = "xxxxx"
        //表格组ID，替换为自己创建的表格组ID
        ZoneId = 2
        //表名称
        TableName = "game_players"
```


### 编译代码
#### 同步模式编译
在示例代码目录，同步模式已经集成了`Makefile`, 方便用户用make方式进行编译。Makefile文件中把GO执行的一系列命令放在里面，用户无需再单独执行
```
#进示例代码目录
cd  tcaplusdb-go-examples/pb/sync

#直接执行make
make

#生成example, test可执行文件,直接执行即可进行相关操作
#体验CRUD接口
./example

#体验读写压测,　大批量压测建议在腾讯云线上环境进行，本地docker环境适合开发调试
./test -t 10 -n 10

```
#### 异步模式编译
异步模式直接进async目录，执行`go build xxx.go`编译相关接口示例代码并执行即可。

# 接口示例
## 接口源代码
目前TcaplusDB PB SDK支持12个接口，基本覆盖所有数据操作场景。SDK所有源代码放置于: `https://github.com/tencentyun/tcaplusdb-go-sdk/pb`。
## 调用模式
TcaplusDB Go PB SDK包括两类模式调用：
* __同步模式__: 接口调用方便，适用于并发量不高的场景。
* __异步模式__: 接口调用相对复杂，适用于高并发、高吞吐的场景。

## 接口调用步骤
* 1.初始化连接客户端
* 2.按需设置客户端参数
* 3.根据指定业务信息，建立客户端连接
* 4.选择调用模式：同步，或异步
* 5.处理响应请求

## 接口源码

导入源码包，软链成`tcaplus`
```
...
import  (
    tcaplus "github.com/tencentyun/tcaplusdb-go-sdk/pb"
)
```

## 客户端接口
### 初始化接口

接口示例
```
//初始化client指针
client := tcaplus.NewPBClient()
```

初始化client后，可灵活设置一些额外配置，可通过如下接口实现：
```
// （非必须）加载日志配置（默认在控制台打印，默认打印级别ERROR, 支持DEBUG|INFO|WARN）
err = client.SetLogCfg("./logconf.xml")
// （非必须）可以设置请求超时时间（默认5s）
err = client.SetDefaultTimeOut(5*time.Second)
// （非必须）可以设置默认表格组ID(ZoneId)(默认zoneList或传入map第一个zoneId)
err = client.SetDefaultZoneId(ZoneId)
```


### 连接接口

接口调用示例
```
//集群ID（业务ID)
appId :=3;
//表格组列表，指定集群下对应的表格组ID, 目的：通过建立一次连接，支持连接到多个游戏区业务表
zoneList := []uint32{3,4};
//集群访问地址，默认端口9999, 对于云环境替换为集群地址，对于docker环境替换为部署机器ip
dirUrl := "tcp://x.x.x.x:9999";
//集群访问密码
signature := "xxxxx";
//超时时间，单位: s
timeOut := 60
//表列表,支持同时指定对应表格组下多个表，key: 表格组id, value: 表列表
zoneTables := map[uint32][]string{3:[]string{"game_players"}};
//建立连接
err := client.Dial(appId, zoneList, dirUrl, signature, timeOut, zoneTables)
//错误处理
if err != nil {
    log.ERR("dial failed %s", err.Error())
    return
}
```

## 同步接口示例
### 插入记录
接口示例
```
// 插入记录, 记录已存在会报错，记录不存在则插入
func insertRecord() {
    //初始化记录
	record := &tcaplusservice.GamePlayers{
		PlayerId:        10805514,
		PlayerName:      "Calvin",
		PlayerEmail:     "calvin@test.com",
		GameServerId:    10,
		LoginTimestamp:  []string{"2019-12-12 15:00:00"},
		LogoutTimestamp: []string{"2019-12-12 16:00:00"},
		IsOnline:        false,
		Pay: &tcaplusservice.Payment{
			PayId:  10101,
			Amount: 1000,
			Method: 1,
		},
	}
    //调用同步模式Insert接口
	err := client.Insert(record)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Case Insert:")
	fmt.Printf("error:%s, message:%+v\n", err, record)
}

```
### 获取记录
```
// 获取记录，记录不存在会报错
func getRecord() {
    //指定获取记录主键，参考表定义主键
	record := &tcaplusservice.GamePlayers{
		PlayerId:    10805514,
		PlayerName:  "Calvin",
		PlayerEmail: "calvin@test.com",
	}
    //调用同步模式Get接口
	err := client.Get(record)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Case Get:")
	fmt.Printf("message:%+v\n", record)
}
```
### 替换记录
```
// 替换记录, 记录不存在则插入,存在则替换
func replaceRecord() {
　//初始化完整记录
	record := &tcaplusservice.GamePlayers{
		PlayerId:        10805514,
		PlayerName:      "Calvin",
		PlayerEmail:     "calvin@test.com",
		GameServerId:    12,
		LoginTimestamp:  []string{"2019-12-12 15:00:00"},
		LogoutTimestamp: []string{"2019-12-12 16:00:00"},
		IsOnline:        false,
		Pay: &tcaplusservice.Payment{
			PayId:  10102,
			Amount: 1002,
			Method: 2,
		},
	}
    //调用同步模式下Replace接口
	err := client.Replace(record)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Case Replace:")
	fmt.Printf("message:%+v\n", record)
}
```
### 更新记录
```
// 更新记录，记录存在则更新，记录不存在则报错
func updateRecord() {
    //初始化更新的记录
	record := &tcaplusservice.GamePlayers{
		PlayerId:        10805514,
		PlayerName:      "Calvin",
		PlayerEmail:     "calvin@test.com",
		GameServerId:    12,
		LoginTimestamp:  []string{"2019-12-12 15:00:00"},
		LogoutTimestamp: []string{"2019-12-12 16:00:00"},
		IsOnline:        false,
		Pay: &tcaplusservice.Payment{
			PayId:  10104,
			Amount: 1004,
			Method: 4,
		},
	}
    //调用同步模式Update接口
	err := client.Update(record)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Case Update:")
	fmt.Printf("message:%+v\n", record)
}

```

### 获取记录部分字段
此接口相比于Get整条记录，能大幅减少返回数据包大小, 节省资源损耗，提升数据获取效率
```
// 获取记录部分字段值,
func fieldGetRecord() {
    //指定记录主键
	record := &tcaplusservice.GamePlayers{
		PlayerId:    10805514,
		PlayerName:  "Calvin",
		PlayerEmail: "calvin@test.com",
	}
    //调用同步模式FieldGet接口获取部分字段值，指定要获取的字段名列表,支持通过点分方式指定嵌套的字段名如pay.pay_id
	err := client.FieldGet(record, []string{"pay", "pay.pay_id"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Case FieldGet:")
	fmt.Printf("message:%+v\n", record)
}
```
### 更新记录部分字段
类似FieldGet接口，避免更新整条记录
```
// 更新部分字段值，记录不存在则报错
func fieldUpdateRecord() {
   //初始化要更新的记录主键，和要更新的字段值
	record := &tcaplusservice.GamePlayers{
		PlayerId:    10805514,
		PlayerName:  "Calvin",
		PlayerEmail: "calvin@test.com",
		Pay: &tcaplusservice.Payment{
			PayId:  10102,
			Amount: 1002,
		},
	}
    //调用同步模式FieldUpdate接口，指定要更新的字段名
	err := client.FieldUpdate(record, []string{"pay.amount", "pay.pay_id"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Case FieldUpdate:")
	fmt.Printf("message:%+v\n", record)
}
```
### 自增记录部分字段
主要适用于数值类型字段自增，使用场景如：玩家等级、经验值。通过自增接口，避免先读值更新后再写回，减少50％的交互次数，并且可保证更新的原子性。
```
// 部分字段自增，仅适用于数值类型字段,记录不存在则报错
//举例：pay.amount 原始值为1000，通过接口指定要自增的量为1002，则自增后，最终pay.amount值为2002
func fieldIncreaseRecord() {
　//指定要更新的记录主键，及要自增的字段值，
	record := &tcaplusservice.GamePlayers{
		PlayerId:    10805514,
		PlayerName:  "Calvin",
		PlayerEmail: "calvin@test.com",
		Pay: &tcaplusservice.Payment{
			PayId:  10102,
			Amount: 1002,
		},
	}
    //调用同步模式接口FieldIncrease, 指定自增的字段名
	err := client.FieldIncrease(record, []string{"pay.amount", "pay.pay_id"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Case FieldIncrease:")
	fmt.Printf("message:%+v\n", record)
}
```
### 删除记录
```
// 删除记录,记录不存在则报错
func deleteRecord() {
    //指定要删除的记录主键
	record := &tcaplusservice.GamePlayers{
		PlayerId:    10805514,
		PlayerName:  "Calvin",
		PlayerEmail: "calvin@test.com",
	}
    //调用同步模式Delete接口
	err := client.Delete(record)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Case Delete:")
	fmt.Printf("message:%+v\n", record)
}
```
### 批量获取记录
```
// 批量获取记录,指定要获取记录的主键
func batchGetRecord() {

	key := &tcaplusservice.GamePlayers{
		PlayerId:    10805510,
		PlayerName:  "Calvin",
		PlayerEmail: "calvin@test.com",
	}
	key2 := &tcaplusservice.GamePlayers{
		PlayerId:    10805511,
		PlayerName:  "Calvin",
		PlayerEmail: "calvin@test.com",
	}

	msgs := []proto.Message{key, key2}
    //调用同步模式BatchGet接口
	err := client.BatchGet(msgs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Case BatchGet:")
	fmt.Printf("message:%+v\n", msgs)
}
```

### 主键索引获取记录
对于TcaplusDB表，在proto表定义文件中会基于主键字段定义主键索引，一个表最多支持4个主键索引，每个主键索引可以包含一个或多个主键字段。
此接口主要是基于主键索引字段进行记录获取，场景：如从公会表中，查询某个玩家所参与的所有公会记录。
```
// 部分key字段获取记录
func partkeyGetRecord() {
    //指定获取记录的主键索引字段，
	record := &tcaplusservice.GamePlayers{
		PlayerId:   10805514,
		PlayerName: "Calvin",
	}
	msgs, err := client.GetByPartKey(record, []string{"player_id", "player_name"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Case GetByPartKey:")
	fmt.Printf("message:%+v\n", msgs)
}
```
### 二级索引获取记录
分布式二级索引是TcaplusDB一个重要的特性，类似MySQL二级索引，支持将表一级字段设置成二级索引字段，通过索引字段可支持范围查询、模糊查询、等值查询等。

```
// 二级索引查询, 需设置全局索引才能使用
func indexQuery() {

	// 非聚合查询
	query := fmt.Sprintf("select pay.pay_id, pay.amount from game_players where player_id=10805514")
	msgs, _, err := client.IndexQuery(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Case IndexQuery:")
	fmt.Printf("message:%+v\n", msgs)

	// 聚合查询
	query = fmt.Sprintf("select count(pay) from game_players where player_id=10805514")
	_, res, err := client.IndexQuery(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Case IndexQuery:")
	fmt.Printf("message:%+v\n", res)
}
```

## 异步接口示例
异步接口调用步骤：
* 1.客户端建立，参考客户端接口部分，包括初始化(NewPBClient)和连接(Dial)两步骤, client := tcaplus.NewPBClient(), client.Dial(...)
* 2.初始化连接请求, client.NewRequest(...)
* 3. 为请求添加一条记录，record, err := request.AddRecord(...)
* 4. 为记录初始化值, record.SetPBData(...)
* 5.发送请求, client.SendRequest(...)
* 6.异步接收请求，resp, err := client.RecvResponse(...)
* 7. 判断响应结果返回码，resp.GetResult(), 如果为0表示SUCCESS，否则处理相应错误码
* 8.若响应SUCCESS, 获取响应记录条数：resp.GetRecordCount(), 并解析记录： resp.FetchRecord()

* __备注__: 异步模式除了遍历表traverse接口外，其它11个接口都是用统一的调用步骤，方便用户统一调用，同时每个接口通过命令字方式来区分接口的不同。
目前异步模式11个接口命令字列表如下：
|命令字别名|命令字编码|命令字说明|
|---|---|---|
|TcaplusApiInsertReq|0x0001|插入记录|
|TcaplusApiReplaceReq|0x0003|替换记录|
|TcaplusApiGetReq|0x0007|获取记录|
|TcaplusApiDeleteReq|0x0009|删除记录|
|TcaplusApiUpdateReq|0x001d|更新记录|
|TcaplusApiBatchGetReq|0x0017|批量获取记录|
|TcaplusApiGetByPartkeyReq|0x0019|根据主键索引获取记录|
|TcaplusApiPBFieldGetReq|0x0067|获取记录部分字段|
|TcaplusApiPBFieldUpdateReq|0x0069|更新记录部分字段|
|TcaplusApiPBFieldIncreaseReq|0x006b|记录部分字段自增|
|TcaplusApiSqlReq|0x0081|二级索引查询|

备注：命令字在`github.com/tencentyun/tcaplusdb-go-sdk/pb/protocol/cmd`中定义

对于异步接口示例代码直接参考`async`下对应接口的示例代码文件，这里不一一列出。下面只针对遍历接口示例进行说明。

### 遍历表
#### 遍历状态
遍历表会实时判断遍历状态，整个遍历过程会在如下状态中进行状态流转，以便合理处理遍历过程。遍历状态源码在：
```
github.com/tencentyun/tcaplusdb-go-sdk/pb/traverser/manager.go
```
如下所示：
```
TraverseStateIdle          = 1      // 结束状态（遍历完毕）
TraverseStateReady         = 2      // 准备状态（初始化成功，可以start）
TraverseStateNormal        = 4      // 遍历中
TraverseStateStop          = 8      // 停止状态（处于此状态会被回收）
TraverseStateRecoverable   = 16     // 可恢复状态（某个响应出问题，可以恢复继续遍历）
TraverseStateUnRecoverable = 32     // 不可恢复状态（获取shardlist出错，或者发生了主备切换）
```

#### 遍历条件设置（非必须）
```
//获取遍历器
// 指定操作表的ZoneId，表名
tra := client.GetTraverser(ZoneId, TableName)

// 设定本次遍历多少条记录，默认遍历所有
err = tra.SetLimit(1000)
// 设置异步id
err = tra.SetAsyncId(12345)
// 设置仅从slave获取记录，默认false
err = tra.SetOnlyReadFromSlave(true)
// 设置userbuf,　自定义数据，序列化为字段流传入
err = tra.SetUserBuff([]byte("test"))
```
备注：
* __userbuf说明__: 对于异步调用，响应包与请求包不在同一内存块，通过userbuf可将一些自定义数据传入遍历请求，再通过响应包返回，避免把这些自定义数据放入全局变量中。不影响业务数据本身，属于用户行为传参，纯粹是为了方便用户编程，少把一些公共的数据直接放全局变量一直保存着。
* __遍历器说明__: 遍历器上限最多 8 个，请在用完后调用 t.Stop() 来回收，否则可能导致
#### 遍历示例
客户端初始化代码参考上述相关部分介绍。
```
// 遍历记录
func traverse() {
   //获取遍历指针
	tra := client.GetTraverser(ZoneId, TableName)
    //遍历开始
	err := tra.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
    //结束遍历
    defer tra.Stop()

	for {
        //异步接收响应
		resp, err := client.RecvResponse()
		if err != nil {
			fmt.Println(err.Error())
			return
		} else if resp == nil {
            //遍历结束状态
			if tra.State() == traverser.TraverseStateIdle {
				break
			} else {
				time.Sleep(time.Microsecond * 10)
				continue
			}
		}
        // 操作response的GetResult获取响应结果
		if err := resp.GetResult(); err != 0 {
			fmt.Println(err)
			return
		}
        //GetRecordCount获取本次响应记录条数,FetchRecord获取响应消息中的记录record，
        // 通过record的GetPBData接口获取响应记录
		for i := 0; i < resp.GetRecordCount(); i++ {
			record, err := resp.FetchRecord()
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			msg := &tcaplusservice.GamePlayers{}
			err = record.GetPBData(msg)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("Case traverse:")
			fmt.Printf("message:%+v\n", msg)
		}
	}
}
```

# 错误码
目前错误码放置在SDK源码中，地址:
```
github.com/tencentyun/tcaplusdb-go-sdk/pb/terror/error.go
```
如果执行代码过程中报相关错误，可先通过源码进行初步查看，也可随时在拉通的业务群中沟通。





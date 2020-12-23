# tcaplusdb-go-examples
[TOC]
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

## 示例代码
以Golang示例代码为例，介绍如何使用PROTOBUF接口进行TcaplusDB表数据操作，这里主要介绍Generic类型表操作。GO示例代码以`go mod`方式进行编译，GO版本以`1.15`举例.



#### 示例代码下载
目前示例代码放在github路径，直接下载即可。
```
git clone https://github.com/tencentyun/tcaplusdb-go-examples.git
```
示例代码目录结构：
```
[root@VM-48-13-centos tcaplusdb-go-examples]# tree
.
|-- pb
|   |-- example.go
|   |-- go.mod
|   |-- go.sum
|   |-- logconf.xml
|   |-- Makefile
|   |-- table
|   |   |-- game_players.proto
|   |   |-- tcaplusservice
|   |   |   |-- game_players.pb.go
|   |   |   `-- tcaplusservice.optionv1.pb.go
|   |   `-- tcaplusservice.optionv1.proto
|   `-- test.go
`-- README.md

```

示例代码说明：
|文件|文件说明|
|---|---|
|example.go| 示例代码主文件,包含CRUD接口|
|test.go|示例压测文件，压测读写接口，开协程并发压测|
|logconf.xml| 日志配置文件，默认级别ERROR, 如果需要更详细的，可开DEBUG,　如果压测的话用ERROR既可，避免性能损耗|
|Makefile|编译文件，直接执行make可编译得到示例二进制文件|
|table|表接口定义目录，包含表相关定义文件，及生成的表接口定义代码文件|
|game_players.proto|示例表定义文件|
|tcaplusservice.optionv1.proto|tcaplusdb pb协议定义文件|
|tcaplusservice|用protoc生成的表定义接口代码目录|
|game_players.pb.go|生成的表定义接口代码|
|tcaplusservice.optionv1.pb.go|生成的pb协议接口代码|

#### 表接口代码生成
如果不想用示例代码中的示例表，参照game_players.proto定义好后，可以用如下命令生成：
```
protoc --go_out=./tcaplusservice mytable.proto
```
备注：
* 需要同时安装protoc-gen-go插件才行
* 需要在proto文件中指定package, 如默认的package tcaplusservice




#### 公共连接环境配置
下载示例代码后，找到example.go文件，并修改如下参数，改成业务方申请的tcaplusdb连接信息或本地部署的tcaplusdb信息;
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


#### 编译代码
在示例代码目录，已经集成了`Makefile`, 方便用户用make方式进行编译。Makefile文件中把GO执行的一系列命令放在里面，用户无需再单独执行
```
#进示例代码目录
cd  tcaplusdb-go-examples/pb

#直接执行make
make

#生成example, test可执行文件,直接执行即可进行相关操作
#体验CRUD接口
./example

#体验读写压测
./test -t 10 -n 10

```


package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"github.com/tencentyun/tcaplusdb-go-examples/pb/table/tcaplusservice"
	"runtime"
	"sync"
	"time"

	"github.com/tencentyun/tcaplusdb-go-sdk/pb"
	"google.golang.org/protobuf/proto"
)

/*******************************************************************************************************************************************
* author : Tcaplus
* created : 2020.05.21
* note :本例将演示TcaplusDB PB API的使用方法, 假定用户已经通过 game_players.proto 在自己的TcaplusDB应用中创建了名为 game_players 的表
创建表格、获取访问点信息的指引请参考 https://cloud.tencent.com/document/product/596/38807。
********************************************************************************************************************************************/

//TcaplusDB RESTful API的连接参数
const (
	//服务接入点,表所在集群Dir连接地址
	DirUrl = "tcp://x.x.x.x:xxxx"
	//应用接入id，表所在集群接入ID
	AppId = 1
	//应用密码,表所在集群访问密码
	Signature = "xxxxx"
	//表格组ID
	ZoneId = 2
	//表名称
	TableName = "game_players"
)

func example() {
	//1.通过指定接入ID(AppId), 表格组id表表(zoneList), 接入地址(DirUrl), 集群密码(Signature) 参数创建TcaplusClient的对象client
	//通过client对象可以访问集群下的所有大区和表
	//创建表格、获取访问点信息的指引请参考 https://cloud.tencent.com/document/product/596/38807
	client := tcaplus.NewPBClient()

	zoneList := []uint32{ZoneId}
	zoneTable := make(map[uint32][]string)
	//构造Map对象存储对应表格组下所有的表
	zoneTable[ZoneId] = []string{TableName}
	//建立到对应集群的连接客户端
	err := client.Dial(AppId, zoneList, DirUrl, Signature, 30, zoneTable)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//2.AddRecord插入记录
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
	err = client.Insert(record)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Case AddRecord:")
	fmt.Printf("error:%s, message:%+v\n", err, record)

	//3.GetRecord查询记录
	key := &tcaplusservice.GamePlayers{
		PlayerId:    10805514,
		PlayerName:  "Calvin",
		PlayerEmail: "calvin@test.com",
	}
	err = client.Get(key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Case GetRecord:")
	fmt.Printf("error:%s, message:%+v\n", err, key)

	//4.DeleteRecord查询记录
	key = &tcaplusservice.GamePlayers{
		PlayerId:    10805514,
		PlayerName:  "Calvin",
		PlayerEmail: "calvin@test.com",
	}
	err = client.Delete(key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Case GetRecord:")
	fmt.Printf("error:%s, message:%+v\n", err, key)

	//5.ReplaceRecord设置记录,存在则更新；不存在，则插入
	//record用户也可将其定义成结构体/map/slice，需可转成json
	record = &tcaplusservice.GamePlayers{
		PlayerId:        10805514,
		PlayerName:      "Calvin",
		PlayerEmail:     "calvin@test.com",
		GameServerId:    10,
		LoginTimestamp:  []string{"2019-12-12 15:00:03"},
		LogoutTimestamp: []string{"2019-12-12 16:00:03"},
		IsOnline:        false,
		Pay: &tcaplusservice.Payment{
			PayId:  10102,
			Amount: 10002,
			Method: 2,
		},
	}
	err = client.Replace(record)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Case SetRecord:")
	fmt.Printf("error:%s, message:%+v\n", err, record)

	//6.UpdateRecord设置记录,存在则更新；不存在，则插入
	record = &tcaplusservice.GamePlayers{
		PlayerId:        10805514,
		PlayerName:      "Calvin",
		PlayerEmail:     "calvin@test.com",
		GameServerId:    10,
		LoginTimestamp:  []string{"2019-12-12 15:00:03"},
		LogoutTimestamp: []string{"2019-12-12 16:00:03"},
		IsOnline:        false,
		Pay: &tcaplusservice.Payment{
			PayId:  10102,
			Amount: 10002,
			Method: 2,
		},
	}
	err = client.Update(record)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Case SetRecord:")
	fmt.Printf("error:%s, message:%+v\n", err, record)

	//7.GetRecord查询记录
	key2 := &tcaplusservice.GamePlayers{
		PlayerId:    10805514,
		PlayerName:  "Calvin",
		PlayerEmail: "calvin@test.com",
	}
	msgs := []proto.Message{key, key2}
	err = client.BatchGet(msgs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Case GetRecord:")
	fmt.Printf("error:%s, message:%+v\n", err, msgs)
}

var (
	ttt = flag.Int("t", 5, "route num")
	nnn = flag.Int("n", 2000, "num")
	f   = flag.String("f", "insert", "select func insert|get")
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	if *f == "get" {
		TcaplusQueryTest(*ttt, *nnn)
	}
	TcaplusInsertTest(*ttt, *nnn)
}

func TcaplusInsertTest(tCount int, num int) {
	c := tcaplus.NewPBClient()
	if err := c.SetLogCfg("./logconf.xml"); err != nil {
		fmt.Println(err.Error())
		return
	}
	zoneList := []uint32{ZoneId}
	zoneTable := make(map[uint32][]string)
	zoneTable[ZoneId] = []string{TableName}
	err := c.Dial(AppId, zoneList, DirUrl, Signature, 30, zoneTable)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	records := make([]*tcaplusservice.GamePlayers, tCount*num)
	for i := 0; i < tCount*num; i++ {
		randseed := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(int64(i)*12345))))
		records[i] = &tcaplusservice.GamePlayers{
			PlayerId:        int64(i) * 12345,
			PlayerName:      string(randseed[:]),
			PlayerEmail:     string(randseed[:]),
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
	}

	for {
		var wg sync.WaitGroup
		start := time.Now()
		maxCost := time.Duration(0)
		for i := 0; i < tCount; i++ {
			wg.Add(1)
			go func(count int, t int) {
				for n := 0; n < count; n++ {
					begin := time.Now()
					c.Insert(records[t*count+n])
					end := time.Since(begin)
					if end > maxCost {
						maxCost = end
					}
				}
				wg.Done()
			}(num, i)
			if i%2 == 0 {
				time.Sleep(time.Nanosecond * 1)
			}
		}
		fmt.Println("start cost", time.Since(start))
		wg.Wait()
		fmt.Println("insert cost: ", time.Since(start), "max ", maxCost)
		time.Sleep(time.Microsecond * 1000)
	}
}

func TcaplusQueryTest(tCount int, num int) {
	c := tcaplus.NewPBClient()
	if err := c.SetLogCfg("./logconf.xml"); err != nil {
		fmt.Println(err.Error())
		return
	}
	zoneList := []uint32{ZoneId}
	zoneTable := make(map[uint32][]string)
	zoneTable[ZoneId] = []string{TableName}
	err := c.Dial(AppId, zoneList, DirUrl, Signature, 30, zoneTable)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	records := make([]*tcaplusservice.GamePlayers, tCount*num)
	for i := 0; i < tCount*num; i++ {
		randseed := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(int64(i)*12345))))
		records[i] = &tcaplusservice.GamePlayers{
			PlayerId:    int64(i) * 12345,
			PlayerName:  string(randseed[:]),
			PlayerEmail: string(randseed[:]),
		}
	}

	for {
		var wg sync.WaitGroup
		start := time.Now()
		maxCost := time.Duration(0)
		for i := 0; i < tCount; i++ {
			wg.Add(1)
			go func(count int, t int) {
				for n := 0; n < count; n++ {
					begin := time.Now()
					c.Get(records[t*count+n])
					end := time.Since(begin)
					if end > maxCost {
						maxCost = end
					}
				}
				wg.Done()
			}(num, i)
			if i%2 == 0 {
				time.Sleep(time.Nanosecond * 1)
			}
		}
		fmt.Println("start cost", time.Since(start))
		wg.Wait()
		fmt.Println("insert cost: ", time.Since(start), "max ", maxCost)
		time.Sleep(time.Microsecond * 1000)
	}
}

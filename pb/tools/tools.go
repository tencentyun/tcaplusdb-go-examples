package tools

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/tencentyun/tcaplusdb-go-examples/pb/cfg"
	tcaplus "github.com/tencentyun/tcaplusdb-go-sdk/pb"
	"github.com/tencentyun/tcaplusdb-go-sdk/pb/response"
)

type clieninf interface {
	RecvResponse() (response.TcaplusResponse, error)
}

//同步接收
func RecvResponse(client clieninf) (response.TcaplusResponse, error) {
	//recv response
	timeOutChan := time.After(5 * time.Second)
	for {
		select {
		case <-timeOutChan:
			return nil, errors.New("5s timeout")
		default:
			resp, err := client.RecvResponse()
			if err != nil {
				return nil, err
			} else if resp == nil {
				time.Sleep(time.Microsecond * 1)
			} else {
				return resp, nil
			}
		}
	}
}

var pbclient *tcaplus.PBClient
var once sync.Once
var ZoneId uint32

func InitPBSyncClient() *tcaplus.PBClient {
	var err error
	once.Do(func() {
		err = cfg.ReadApiCfg("../cfg/api_cfg.xml")
		if err != nil {
			fmt.Printf("ReadApiCfg fail %s", err.Error())
			return
		}

		pbclient = tcaplus.NewPBClient()
		err = pbclient.SetLogCfg("../cfg/logconf.xml")
		if err != nil {
			fmt.Printf("excepted SetLogCfg success")
			return
		}
		//可选，作用于docker环境，用于解决本地开发环境无法连通docker容器问题
　　pbclient.SetPublicIP(cfg.ApiConfig.PublicIP)

		ZoneId = cfg.ApiConfig.ZoneId

		tables := strings.Split(cfg.ApiConfig.PBTable, ",")
		zoneTable := map[uint32][]string{cfg.ApiConfig.ZoneId: tables}
		err = pbclient.Dial(cfg.ApiConfig.AppId, []uint32{cfg.ApiConfig.ZoneId}, cfg.ApiConfig.DirUrl, cfg.ApiConfig.Signature, 30, zoneTable)
		if err != nil {
			fmt.Printf("excepted dial success, %s", err.Error())
			return
		}
	})
	if err != nil {
		return nil
	}
	return pbclient
}

func ConvertToJson(v interface{}) string {
	body, _ := json.Marshal(v)
	buf := &bytes.Buffer{}
	json.Indent(buf, body, "", "\t")
	return buf.String()
}

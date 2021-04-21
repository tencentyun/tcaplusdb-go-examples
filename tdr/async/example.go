package main

import (
	"errors"
	"fmt"
	"github.com/tencentyun/tcaplusdb-go-examples/tdr/async/service_info"
	tcaplus "github.com/tencentyun/tcaplusdb-go-sdk/tdr"
	"github.com/tencentyun/tcaplusdb-go-sdk/tdr/protocol/cmd"
	"github.com/tencentyun/tcaplusdb-go-sdk/tdr/response"
	"github.com/tencentyun/tcaplusdb-go-sdk/tdr/terror"
	"strconv"
	"time"
)

const (
	AppId     = uint64(2)
	ZoneId    = uint32(3)
	DirUrl    = "tcp://9.135.8.93:9999"
	Signature = "2220C41F8AA79100"
	TableName = "service_info"
)

var client *tcaplus.Client

//同步接收
func recvResponse(client *tcaplus.Client) (response.TcaplusResponse, error) {
	//5s超时
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

func preInsert(data *service_info.Service_Info) {
	req, err := client.NewRequest(ZoneId, TableName, cmd.TcaplusApiInsertReq)
	if err != nil {
		fmt.Printf("NewRequest TcaplusApiInsertReq failed %v\n", err.Error())
		return
	}
	rec, err := req.AddRecord(0)
	if err != nil {
		fmt.Printf("AddRecord failed %v\n", err.Error())
		return
	}
	if err := rec.SetData(data); err != nil {
		fmt.Printf("SetData failed %v\n", err.Error())
		return
	}
	//发送请求接收响应
	_, err = client.Do(req, 5*time.Second)
	if err != nil {
		fmt.Printf("recv err %s\n", err.Error())
		return
	}
}

func preDelete(data *service_info.Service_Info) {
	req, err := client.NewRequest(ZoneId, TableName, cmd.TcaplusApiDeleteReq)
	if err != nil {
		fmt.Printf("NewRequest TcaplusApiDeleteReq failed %v\n", err.Error())
		return
	}
	rec, err := req.AddRecord(0)
	if err != nil {
		fmt.Printf("AddRecord failed %v\n", err.Error())
		return
	}
	if err := rec.SetData(data); err != nil {
		fmt.Printf("SetData failed %v\n", err.Error())
		return
	}
	//发送请求接收响应
	_, err = client.Do(req, 5*time.Second)
	if err != nil {
		fmt.Printf("recv err %s\n", err.Error())
		return
	}
}

func main() {
	client = tcaplus.NewClient()
	//日志配置，不配置则debug打印到控制台
	if err := client.SetLogCfg("./logconf.xml"); err != nil {
		fmt.Println(err.Error())
		return
	}
	//client连接tcaplus
	err := client.Dial(AppId, []uint32{ZoneId}, DirUrl, Signature, 60)
	if err != nil {
		fmt.Printf("init failed %v\n", err.Error())
		return
	}
	fmt.Printf("Dial finish\n")
	getPartKeyExample()
	deleteByPartKeyExample()
	insertExample()
	getExample()
	updateExample()
	replaceExample()
	deleteExample()
	batchGetExample()
	indexQueryExample()
}

func insertExample() {
	//创建insert请求
	req, err := client.NewRequest(ZoneId, TableName, cmd.TcaplusApiInsertReq)
	if err != nil {
		fmt.Printf("NewRequest TcaplusApiInsertReq failed %v\n", err.Error())
		return
	}
	fmt.Printf("insertExample NewRequest TcaplusApiInsertReq finish\n")
	//设置异步请求ID，异步请求通过ID让响应和请求对应起来
	req.SetAsyncId(666)
	//为request添加一条记录，（index只有在list表中支持，generic不校验）
	rec, err := req.AddRecord(0)
	if err != nil {
		fmt.Printf("AddRecord failed %v\n", err.Error())
		return
	}
	fmt.Printf("insertExample AddRecord finish\n")
	//申请tdr结构体并赋值，最好调用tdr pkg的NewXXX函数，会将成员初始化为tdr定义的tdr默认值，
	// 不要自己new，自己new，某些结构体未初始化，存在panic的风险
	data := service_info.NewService_Info()
	data.Gameid = "dev"
	data.Envdata = "oa"
	data.Name = "com"
	data.Filterdata = time.Now().Format("2006-01-02T15:04:05.000000Z")
	data.Updatetime = uint64(time.Now().UnixNano())
	data.Inst_Max_Num = 2
	data.Inst_Min_Num = 3
	//数组类型为slice需要准确赋值长度，与refer保持一致
	route := "test"
	data.Routeinfo_Len = uint32(len(route))
	data.Routeinfo = []byte(route)

	// 防止记录已存在
	preDelete(data)

	//将tdr的数据设置到请求的记录中
	if err := rec.SetData(data); err != nil {
		fmt.Printf("SetData failed %v\n", err.Error())
		return
	}
	//发送请求
	if err := client.SendRequest(req); err != nil {
		fmt.Printf("SendRequest failed %v\n", err.Error())
		return
	}
	fmt.Printf("insertExample send finish\n")
	//接收响应
	resp, err := recvResponse(client)
	if err != nil {
		fmt.Printf("recv err %s\n", err.Error())
		return
	}
	//带回请求的异步ID
	fmt.Printf("insertExample resp success, AsyncId:%d\n", resp.GetAsyncId())
	tcapluserr := resp.GetResult()
	if tcapluserr != 0 {
		fmt.Printf("response ret errCode: %d, errMsg: %s", tcapluserr, terror.GetErrMsg(tcapluserr))
		return
	}
	fmt.Println("insertExample Success")
}

func getExample() {
	//创建Get请求
	req, err := client.NewRequest(ZoneId, TableName, cmd.TcaplusApiGetReq)
	if err != nil {
		fmt.Printf("getExample NewRequest TcaplusApiGetReq failed %v\n", err.Error())
		return
	}
	fmt.Printf("getExample NewRequest TcaplusApiGetReq finish\n")
	//设置异步请求ID，异步请求通过ID让响应和请求对应起来
	req.SetAsyncId(667)
	//为request添加一条记录，（index只有在list表中支持，generic不校验）
	rec, err := req.AddRecord(0)
	if err != nil {
		fmt.Printf("getExample AddRecord failed %v\n", err.Error())
		return
	}
	fmt.Printf("getExample AddRecord finish\n")
	//申请tdr结构体并赋值Key，最好调用tdr pkg的NewXXX函数，会将成员初始化为tdr定义的tdr默认值，
	// 不要自己new，自己new，某些结构体未初始化，存在panic的风险
	data := service_info.NewService_Info()
	data.Gameid = "dev"
	data.Envdata = "oa"
	data.Name = "com"
	//将tdr的数据设置到请求的记录中

	// 方剂记录不存在
	preInsert(data)

	if err := rec.SetData(data); err != nil {
		fmt.Printf("SetData failed %v\n", err.Error())
		return
	}
	fmt.Printf("getExample SetData finish\n")
	if err := client.SendRequest(req); err != nil {
		fmt.Printf("SendRequest failed %v\n", err.Error())
		return
	}
	fmt.Printf("getExample send finish\n")
	resp, err := recvResponse(client)
	if err != nil {
		fmt.Printf("recv err %s\n", err.Error())
		return
	}
	//带回请求的异步ID
	fmt.Printf("getExample resp success, AsyncId:%d\n", resp.GetAsyncId())
	tcapluserr := resp.GetResult()
	if tcapluserr != 0 {
		fmt.Printf("response ret %s\n",
			"errCode: "+strconv.Itoa(tcapluserr)+", errMsg: "+terror.ErrorCodes[tcapluserr])
		return
	}
	//response中带有获取的记录
	fmt.Printf("getExample response success record count %d\n", resp.GetRecordCount())
	for i := 0; i < resp.GetRecordCount(); i++ {
		record, err := resp.FetchRecord()
		if err != nil {
			fmt.Printf("FetchRecord failed %s\n", err.Error())
			return
		}
		//通过GetData获取记录
		data := service_info.NewService_Info()
		if err := record.GetData(data); err != nil {
			fmt.Printf("record.GetData failed %s\n", err.Error())
			return
		}
		fmt.Printf("getExample response record data %+v, route: %s\n",
			data, string(data.Routeinfo[0:data.Routeinfo_Len]))
	}
	fmt.Println("getExample Success")
}

func getPartKeyExample() {
	//创建Get请求
	req, err := client.NewRequest(ZoneId, TableName, cmd.TcaplusApiGetByPartkeyReq)
	if err != nil {
		fmt.Printf("getPartKeyExample NewRequest TcaplusApiGetReq failed %v\n", err.Error())
		return
	}
	fmt.Printf("getPartKeyExample NewRequest TcaplusApiGetReq finish\n")
	//设置异步请求ID，异步请求通过ID让响应和请求对应起来
	req.SetAsyncId(667)
	req.SetResultLimit(5000,1)
	//为request添加一条记录，（index只有在list表中支持，generic不校验）
	rec, err := req.AddRecord(0)
	if err != nil {
		fmt.Printf("getPartKeyExample AddRecord failed %v\n", err.Error())
		return
	}
	fmt.Printf("getPartKeyExample AddRecord finish\n")
	//申请tdr结构体并赋值Key，最好调用tdr pkg的NewXXX函数，会将成员初始化为tdr定义的tdr默认值，
	// 不要自己new，自己new，某些结构体未初始化，存在panic的风险
	data := service_info.NewService_Info()
	data.Gameid = "dev"
	//data.Envdata = "oaasqomk"
	data.Name = "com"

	// 防止记录不存在
	data.Envdata = "456"
	data.Expansion = "123"
	preInsert(data)
	data.Envdata = "123"
	data.Expansion = "456"
	preInsert(data)

	//将tdr的数据设置到请求的记录中
//	flist := []string {"updatetime"}
	var flist  []string = nil
	if err := rec.SetDataWithIndexAndField(data, flist,"Index_Gameid_Name"); err != nil {
		fmt.Printf("SetData failed %v\n", err.Error())
		return
	}
	fmt.Printf("value map : %d\n", len(rec.ValueMap))
	fmt.Printf("getPartKeyExample SetData finish\n")
	if err := client.SendRequest(req); err != nil {
		fmt.Printf("SendRequest failed %v\n", err.Error())
		return
	}
	var total int
	for {
		fmt.Printf("getPartKeyExample send finish\n")
		resp, err := recvResponse(client)
		if err != nil {
			fmt.Printf("recv err %s\n", err.Error())
			return
		}
		//带回请求的异步ID
		fmt.Printf("getPartKeyExample resp success, AsyncId:%d\n", resp.GetAsyncId())
		tcapluserr := resp.GetResult()
		if tcapluserr != 0 {
			fmt.Printf("response ret %s\n",
				"errCode: "+strconv.Itoa(tcapluserr)+", errMsg: "+terror.ErrorCodes[tcapluserr])
			return
		}
		haveMore := resp.HaveMoreResPkgs()
		//response中带有获取的记录
		total +=  resp.GetRecordCount()
		fmt.Printf("getPartKeyExample response success record count %d, total:%d\n",
			resp.GetRecordCount(),total)
		//idx_max := resp.GetRecordCount()
		//receive_flag := resp.(*response.GetByPartKeyResponse).IsRspReceiveFinish()
		//fmt.Printf("getPartKeyExample response success record count %d\n", resp.GetRecordCount())

		//for i := 0; i < idx_max; i++ {
		//	record, err := resp.FetchRecord()
		//	if err != nil {
		//		fmt.Printf("FetchRecord failed %s\n", err.Error())
		//		return
		//	}
		//	//通过GetData获取记录
		//	data := service_info.NewService_Info()
		//	if err := record.GetData(data); err != nil {
		//		fmt.Printf("record.GetData failed %s\n", err.Error())
		//		return
		//	}
		//	//fmt.Printf("")
		//	//fmt.Printf("getPartKeyExample response record data %+v, route: %s\n",
		//	data, string(data.Routeinfo[0:data.Routeinfo_Len]))
		//}
		if 0 == haveMore{
			break
		}
	}
	fmt.Println("getPartKeyExample Success")
}

func updateExample() {
	//创建Update请求
	req, err := client.NewRequest(ZoneId, TableName, cmd.TcaplusApiUpdateReq)
	if err != nil {
		fmt.Printf("NewRequest TcaplusApiUpdateReq failed %v\n", err.Error())
		return
	}
	fmt.Printf("updateExample NewRequest TcaplusApiUpdateReq finish\n")
	//设置异步请求ID，异步请求通过ID让响应和请求对应起来
	req.SetAsyncId(668)
	//设置结果标记位，更新成功后，返回tcaplus端的旧数据，默认为0
	if err := req.SetResultFlag(3); err != nil {
		fmt.Printf("SetResultFlag failed %v\n", err.Error())
		return
	}
	fmt.Printf("updateExample SetResultFlag finish\n")
	//为request添加一条记录，（index只有在list表中支持，generic不校验）
	rec, err := req.AddRecord(0)
	if err != nil {
		fmt.Printf("updateExample AddRecord failed %v\n", err.Error())
		return
	}
	fmt.Printf("updateExample AddRecord finish\n")
	//申请tdr结构体并赋值，最好调用tdr pkg的NewXXX函数，会将成员初始化为tdr定义的tdr默认值，
	// 不要自己new，自己new，某些结构体未初始化，存在panic的风险
	data := service_info.NewService_Info()
	data.Gameid = "dev"
	data.Envdata = "oa"
	data.Name = "com"
	data.Filterdata = time.Now().Format("2006-01-02T15:04:05.000000Z")
	data.Updatetime = uint64(time.Now().UnixNano())
	data.Inst_Max_Num = 2
	data.Inst_Min_Num = 3
	route := "test"
	data.Routeinfo_Len = uint32(len(route))
	data.Routeinfo = []byte(route)

	// 防止记录不存在
	preInsert(data)

	//将tdr的数据设置到请求的记录中
	if err := rec.SetData(data); err != nil {
		fmt.Printf("SetData failed %v\n", err.Error())
		return
	}
	if err := client.SendRequest(req); err != nil {
		fmt.Printf("SendRequest failed %v\n", err.Error())
		return
	}
	fmt.Printf("updateExample send finish\n")
	//recv response
	resp, err := recvResponse(client)
	if err != nil {
		fmt.Printf("recv err %s\n", err.Error())
		return
	}
	//带回请求的异步ID
	fmt.Printf("updateExample resp success, AsyncId:%d\n", resp.GetAsyncId())
	tcapluserr := resp.GetResult()
	if tcapluserr != 0 {
		fmt.Printf("response ret errCode: %d, errMsg: %s", tcapluserr, terror.GetErrMsg(tcapluserr))
		return
	}
	//response中带有获取的旧记录
	fmt.Printf("updateExample response success record count %d\n", resp.GetRecordCount())
	for i := 0; i < resp.GetRecordCount(); i++ {
		record, err := resp.FetchRecord()
		if err != nil {
			fmt.Printf("FetchRecord failed %s\n", err.Error())
			return
		}
		oldData := service_info.NewService_Info()
		if err := record.GetData(oldData); err != nil {
			fmt.Printf("record.GetData failed %s\n", err.Error())
			return
		}
		fmt.Printf("updateExample response record data %+v, route: %s\n",
			oldData, string(oldData.Routeinfo[0:oldData.Routeinfo_Len]))
		fmt.Printf("updateExample request  record data %+v, route: %s\n",
			data, string(data.Routeinfo[0:data.Routeinfo_Len]))
	}
	fmt.Println("updateExample Success")
}

func replaceExample() {
	//创建Replace请求
	req, err := client.NewRequest(ZoneId, TableName, cmd.TcaplusApiReplaceReq)
	if err != nil {
		fmt.Printf("NewRequest TcaplusApiReplaceReq failed %v\n", err.Error())
		return
	}
	fmt.Printf("replaceExample NewRequest TcaplusApiReplaceReq finish\n")
	//设置异步请求ID，异步请求通过ID让响应和请求对应起来
	req.SetAsyncId(669)
	//设置结果标记位，更新成功后，返回tcaplus端的旧数据，默认为0
	if err := req.SetResultFlag(3); err != nil {
		fmt.Printf("SetResultFlag failed %v\n", err.Error())
		return
	}
	fmt.Printf("replaceExample SetResultFlag finish\n")
	//为request添加一条记录，（index只有在list表中支持，generic不校验）
	rec, err := req.AddRecord(0)
	if err != nil {
		fmt.Printf("AddRecord failed %v\n", err.Error())
		return
	}
	fmt.Printf("replaceExample AddRecord finish\n")
	//申请tdr结构体并赋值，最好调用tdr pkg的NewXXX函数，会将成员初始化为tdr定义的tdr默认值，
	// 不要自己new，自己new，某些结构体未初始化，存在panic的风险
	data := service_info.NewService_Info()
	data.Gameid = "dev"
	data.Envdata = "oa"
	data.Name = "com"
	data.Filterdata = time.Now().Format("2006-01-02T15:04:05.000000Z")
	data.Updatetime = uint64(time.Now().UnixNano())
	data.Inst_Max_Num = 2
	data.Inst_Min_Num = 3
	route := "test"
	data.Routeinfo_Len = uint32(len(route))
	data.Routeinfo = []byte(route)
	//将tdr的数据设置到请求的记录中
	if err := rec.SetData(data); err != nil {
		fmt.Printf("SetData failed %v\n", err.Error())
		return
	}
	if err := client.SendRequest(req); err != nil {
		fmt.Printf("SendRequest failed %v\n", err.Error())
		return
	}
	fmt.Printf("replaceExample send finish\n")
	//recv response
	resp, err := recvResponse(client)
	if err != nil {
		fmt.Printf("recv err %s\n", err.Error())
		return
	}
	//带回请求的异步ID
	fmt.Printf("replaceExample resp success, AsyncId:%d\n", resp.GetAsyncId())
	tcapluserr := resp.GetResult()
	if tcapluserr != 0 {
		fmt.Printf("response ret errCode: %d, errMsg: %s", tcapluserr, terror.GetErrMsg(tcapluserr))
		return
	}
	//response中带有获取的旧记录
	fmt.Printf("replaceExample response success record count %d\n", resp.GetRecordCount())
	for i := 0; i < resp.GetRecordCount(); i++ {
		record, err := resp.FetchRecord()
		if err != nil {
			fmt.Printf("FetchRecord failed %s\n", err.Error())
			return
		}
		oldData := service_info.NewService_Info()
		if err := record.GetData(oldData); err != nil {
			fmt.Printf("record.GetData failed %s\n", err.Error())
			return
		}
		fmt.Printf("replaceExample response record data %+v, route: %s\n",
			oldData, string(oldData.Routeinfo[0:oldData.Routeinfo_Len]))
		fmt.Printf("replaceExample request  record data %+v, route: %s\n",
			data, string(data.Routeinfo[0:data.Routeinfo_Len]))
	}
	fmt.Println("replaceExample Success")
}

func deleteExample() {
	//创建Delete请求
	req, err := client.NewRequest(ZoneId, TableName, cmd.TcaplusApiDeleteReq)
	if err != nil {
		fmt.Printf("NewRequest TcaplusApiDeleteReq failed %v\n", err.Error())
		return
	}
	fmt.Printf("deleteExample NewRequest TcaplusApiDeleteReq finish\n")
	//设置异步请求ID，异步请求通过ID让响应和请求对应起来
	req.SetAsyncId(670)
	//设置结果标记位，删除成功后，返回tcaplus端的旧数据，默认为0
	if err := req.SetResultFlag(3); err != nil {
		fmt.Printf("SetResultFlag failed %v\n", err.Error())
		return
	}
	fmt.Printf("deleteExample SetResultFlag finish\n")
	//为request添加一条记录，（index只有在list表中支持，generic不校验）
	rec, err := req.AddRecord(0)
	if err != nil {
		fmt.Printf("AddRecord failed %v\n", err.Error())
		return
	}
	fmt.Printf("deleteExample AddRecord finish\n")
	//申请tdr结构体并赋值key，最好调用tdr pkg的NewXXX函数，会将成员初始化为tdr定义的tdr默认值，
	// 不要自己new，自己new，某些结构体未初始化，存在panic的风险
	data := service_info.NewService_Info()
	data.Gameid = "gid_1"
	data.Envdata = "oagfgadsf"
	data.Name = "com"
	data.Expansion = "fds"

	// 防止记录不存在
	preInsert(data)

	//将tdr的数据设置到请求的记录中
	if err := rec.SetData(data); err != nil {
		fmt.Printf("SetData failed %v\n", err.Error())
		return
	}
	if err := client.SendRequest(req); err != nil {
		fmt.Printf("SendRequest failed %v\n", err.Error())
		return
	}
	fmt.Printf("deleteExample send finish\n")
	//recv response
	resp, err := recvResponse(client)
	if err != nil {
		fmt.Printf("recv err %s\n", err.Error())
		return
	}
	//带回请求的异步ID
	fmt.Printf("deleteExample resp success, AsyncId:%d\n", resp.GetAsyncId())
	tcapluserr := resp.GetResult()
	if tcapluserr != 0 {
		fmt.Printf("response ret errCode: %d, errMsg: %s", tcapluserr, terror.GetErrMsg(tcapluserr))
		return
	}
	//response中带有获取的旧记录
	fmt.Printf("deleteExample response success record count %d\n", resp.GetRecordCount())
	for i := 0; i < resp.GetRecordCount(); i++ {
		record, err := resp.FetchRecord()
		if err != nil {
			fmt.Printf("FetchRecord failed %s\n", err.Error())
			return
		}
		oldData := service_info.NewService_Info()
		if err := record.GetData(oldData); err != nil {
			fmt.Printf("record.GetData failed %s\n", err.Error())
			return
		}
		fmt.Printf("deleteExample response record data %+v, route: %s\n",
			oldData, string(oldData.Routeinfo[0:oldData.Routeinfo_Len]))
		fmt.Printf("deleteExample request  record data %+v, route: %s\n",
			data, string(data.Routeinfo[0:data.Routeinfo_Len]))
	}

	fmt.Println("deleteExample Success")
}

func deleteByPartKeyExample() {
	//创建Delete请求
	req, err := client.NewRequest(ZoneId, TableName, cmd.TcaplusApiDeleteByPartkeyReq)
	if err != nil {
		fmt.Printf("NewRequest TcaplusApiDeleteByPartkeyReq failed %v\n", err.Error())
		return
	}
	fmt.Printf("deleteByPartKeyExample NewRequest TcaplusApiDeleteByPartkeyReq finish\n")
	//设置异步请求ID，异步请求通过ID让响应和请求对应起来
	req.SetAsyncId(770)
	//设置结果标记位，删除成功后，返回tcaplus端的旧数据，默认为0
	if err := req.SetVersionPolicy(3); err != nil {
		fmt.Printf("SetResultFlag failed %v\n", err.Error())
		return
	}
	fmt.Printf("deleteByPartKeyExample SetResultFlag finish\n")
	//为request添加一条记录，（index只有在list表中支持，generic不校验）
	rec, err := req.AddRecord(0)
	if err != nil {
		fmt.Printf("AddRecord failed %v\n", err.Error())
		return
	}
	req.SetResultLimit(20000,0)
	fmt.Printf("deleteByPartKeyExample AddRecord finish\n")
	//申请tdr结构体并赋值key，最好调用tdr pkg的NewXXX函数，会将成员初始化为tdr定义的tdr默认值，
	// 不要自己new，自己new，某些结构体未初始化，存在panic的风险
	data := service_info.NewService_Info()
	data.Gameid = "dev"
	//data.Envdata = "oaasqomk"
	data.Name = "com"

	// 防止记录不存在
	data.Envdata = "456"
	data.Expansion = "123"
	preInsert(data)
	data.Envdata = "123"
	data.Expansion = "456"
	preInsert(data)

	//将tdr的数据设置到请求的记录中
	if err := rec.SetDataWithIndexAndField(data, nil,"Index_Gameid_Name"); err != nil {
		fmt.Printf("SetData failed %v\n", err.Error())
		return
	}
	if err := client.SendRequest(req); err != nil {
		fmt.Printf("SendRequest failed %v\n", err.Error())
		return
	}
	fmt.Printf("deleteByPartKeyExample send finish\n")
	var total_num int = 0;
	//recv response
	for {
		resp, err := recvResponse(client)
		if err != nil {
			fmt.Printf("recv err %s\n", err.Error())
			return
		}
		//带回请求的异步ID
		fmt.Printf("deleteByPartKeyExample resp success, AsyncId:%d\n", resp.GetAsyncId())
		tcapluserr := resp.GetResult()
		if tcapluserr != 0 {
			fmt.Printf("response ret errCode: %d, errMsg: %s", tcapluserr, terror.GetErrMsg(tcapluserr))
			return
		}
		has_more := resp.HaveMoreResPkgs()
		total_num += resp.GetRecordCount()
		fmt.Printf("deleteByPartKeyExample resp success, total_num:%d\n", total_num)
		  //  response中带有获取的旧记录
		fmt.Printf("deleteByPartKeyExample response success record count %d\n", resp.GetRecordCount())
		for i := 0; i < resp.GetRecordCount(); i++ {
			record, err := resp.FetchRecord()
			if err != nil {
				fmt.Printf("FetchRecord failed %s\n", err.Error())
				return
			}
			oldData := service_info.NewService_Info()
			if err := record.GetData(oldData); err != nil {
				fmt.Printf("record.GetData failed %s\n", err.Error())
				return
			}
			fmt.Printf("gameid:%s, Envdata %s, name:%s, Expansion:%s\n",
				oldData.Gameid, oldData.Envdata,oldData.Name,oldData.Expansion)
			//fmt.Printf("\ndeleteByPartKeyExample response record data %+v, route: %s",
			// oldData, string(oldData.Routeinfo[0:oldData.Routeinfo_Len]))
			//fmt.Printf("\ndeleteByPartKeyExample request  record data %+v, route: %s",
			// data, string(data.Routeinfo[0:data.Routeinfo_Len]))
		}
		if 0 == has_more{
			break
		}
	}
	fmt.Println("deleteByPartKeyExample Success")
}

func batchGetExample() {
	//创建BatchGet请求
	req, err := client.NewRequest(ZoneId, TableName, cmd.TcaplusApiBatchGetReq)
	if err != nil {
		fmt.Printf("batchGetExample NewRequest TcaplusApiGetReq failed %v\n", err.Error())
		return
	}
	fmt.Printf("batchGetExample NewRequest TcaplusApiGetReq finish\n")
	//设置异步请求ID，异步请求通过ID让响应和请求对应起来
	req.SetAsyncId(775)
	//为request添加一条记录，（index只有在list表中支持，generic不校验）
	rec, err := req.AddRecord(0)
	if err != nil {
		fmt.Printf("batchGetExample AddRecord failed %v\n", err.Error())
		return
	}
	fmt.Printf("batchGetExample AddRecord finish\n")
	//申请tdr结构体并赋值Key，最好调用tdr pkg的NewXXX函数，会将成员初始化为tdr定义的tdr默认值，
	// 不要自己new，自己new，某些结构体未初始化，存在panic的风险
	data := service_info.NewService_Info()
	data.Gameid = "dev"
	//data.Envdata = "oaasqomk"
	data.Name = "com"

	// 防止记录不存在
	data.Envdata = "456"
	data.Expansion = "123"
	preInsert(data)
	if err := rec.SetData(data); err != nil {
		fmt.Printf("SetData failed %v\n", err.Error())
		return
	}

	data.Envdata = "123"
	data.Expansion = "456"
	preInsert(data)
	rec, err = req.AddRecord(0)
	if err != nil {
		fmt.Printf("batchGetExample AddRecord failed %v\n", err.Error())
		return
	}
	fmt.Printf("batchGetExample AddRecord finish\n")
	if err := rec.SetData(data); err != nil {
		fmt.Printf("SetData failed %v\n", err.Error())
		return
	}

	fmt.Printf("value map : %d\n", len(rec.ValueMap))
	fmt.Printf("batchGetExample SetData finish\n")
	if err := client.SendRequest(req); err != nil {
		fmt.Printf("SendRequest failed %v\n", err.Error())
		return
	}
	var total int
	for {
		fmt.Printf("batchGetExample send finish\n")
		resp, err := recvResponse(client)
		if err != nil {
			fmt.Printf("recv err %s\n", err.Error())
			return
		}
		//带回请求的异步ID
		fmt.Printf("batchGetExample resp success, AsyncId:%d\n", resp.GetAsyncId())
		tcapluserr := resp.GetResult()
		if tcapluserr != 0 {
			fmt.Printf("response ret %s\n",
				"errCode: "+strconv.Itoa(tcapluserr)+", errMsg: "+terror.ErrorCodes[tcapluserr])
			return
		}
		haveMore := resp.HaveMoreResPkgs()
		//response中带有获取的记录
		total +=  resp.GetRecordCount()
		fmt.Printf("batchGetExample response success record count %d, total:%d\n",
			resp.GetRecordCount(),total)

		if 0 == haveMore{
			break
		}
	}
	fmt.Println("batchGetExample Success")
}

func indexQueryExample()() {
	//创建BatchGet请求
	req, err := client.NewRequest(ZoneId, TableName, cmd.TcaplusApiSqlReq)
	if err != nil {
		fmt.Printf("indexQueryExample NewRequest TcaplusApiGetReq failed %v\n", err.Error())
		return
	}
	fmt.Printf("indexQueryExample NewRequest TcaplusApiGetReq finish\n")
	//设置异步请求ID，异步请求通过ID让响应和请求对应起来
	req.SetAsyncId(775)

	query := fmt.Sprintf("select * from service_info where gameid=dev and name=com")
	req.SetSql(query)
	fmt.Printf("indexQueryExample SetSql finish\n")
	//申请tdr结构体并赋值Key，最好调用tdr pkg的NewXXX函数，会将成员初始化为tdr定义的tdr默认值，
	// 不要自己new，自己new，某些结构体未初始化，存在panic的风险
	data := service_info.NewService_Info()
	data.Gameid = "dev"
	data.Name = "com"

	// 防止记录不存在
	data.Envdata = "456"
	data.Expansion = "123"
	preInsert(data)
	data.Envdata = "123"
	data.Expansion = "456"
	preInsert(data)

	if err := client.SendRequest(req); err != nil {
		fmt.Printf("SendRequest failed %v\n", err.Error())
		return
	}
	var total int
	for {
		fmt.Printf("indexQueryExample send finish\n")
		resp, err := recvResponse(client)
		if err != nil {
			fmt.Printf("recv err %s\n", err.Error())
			return
		}
		//带回请求的异步ID
		fmt.Printf("indexQueryExample resp success, AsyncId:%d\n", resp.GetAsyncId())
		tcapluserr := resp.GetResult()
		if tcapluserr != 0 {
			fmt.Printf("response ret %s\n",
				"errCode: "+strconv.Itoa(tcapluserr)+", errMsg: "+terror.ErrorCodes[tcapluserr])
			return
		}
		haveMore := resp.HaveMoreResPkgs()
		//response中带有获取的记录
		total +=  resp.GetRecordCount()
		fmt.Printf("indexQueryExample response success record count %d, total:%d\n",
			resp.GetRecordCount(),total)

		if 0 == haveMore{
			break
		}
	}
	fmt.Println("indexQueryExample Success")
}

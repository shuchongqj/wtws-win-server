package dto

import (
	wtws_mysql "wtws-server/models/wtws-mysql"
)

type WeighOrderList struct {
	List  []wtws_mysql.OrWeighOrder `json:"list"`
	Count int                       `json:"count"`
}

type ScanVehicleData struct {
	//Mid          int         `json:"mid"`
	//Mly          string      `json:"mly"`
	//Mno          string      `json:"mno"`
	//Mzdrid       int         `json:"mzdrid"`
	//Mzdrname     string      `json:"mzdrname"`
	//Msjid        int         `json:"msjid"`
	//Msjname      string      `json:"msjname"`
	//Msjcp        string      `json:"msjcp"`
	//Msjtel       string      `json:"msjtel"`
	//Msjzd        string      `json:"msjzd"`
	//Mshdwid      int         `json:"mshdwid"`
	//Mshdh        string      `json:"mshdh"`
	//Mshdz        string      `json:"mshdz"`
	//Mfhdwid      int         `json:"mfhdwid"`
	//Mfhdh        string      `json:"mfhdh"`
	//Mfhdz        string      `json:"mfhdz"`
	//Myssj        string      `json:"myssj"`
	//Mshrid       int         `json:"mshrid"`
	//Mshrname     string      `json:"mshrname"`
	//Mshzt        string      `json:"mshzt"`
	//Mshbz        interface{} `json:"mshbz"`
	//Mcpid        int         `json:"mcpid"`
	//Mcpname      string      `json:"mcpname"`
	//Mcpbh        string      `json:"mcpbh"`
	//Mcpsl        int         `json:"mcpsl"`
	//Mcpdw        string      `json:"mcpdw"`
	//Mcpzl        float32     `json:"mcpzl"`
	//Mcpgg        string      `json:"mcpgg"`
	//Mcpbz        string      `json:"mcpbz"`
	//Myh          string      `json:"myh"`
	//Myhkh        string      `json:"myhkh"`
	//Mysl         float32     `json:"mysl"`
	//Voidtime     int         `json:"voidtime"`
	//Over         int         `json:"over"`
	//Xianzhong    int         `json:"xianzhong"`
	//ID           int         `json:"id"`
	//Cateid       int         `json:"cateid"`
	//Prname       string      `json:"prname"`
	//Pnum         string      `json:"pnum"`
	//Spec         float64     `json:"spec"`
	//Goodsbianhao string      `json:"goodsbianhao"`

	Allow         float32 `json:"allow"`
	Wckname       string  `json:"Wckname"`
	Bank          string  `json:"bank"`
	Banknum       string  `json:"banknum"`
	Drivername    string  `json:"drivername"`
	Drivertel     string  `json:"drivertel"`
	Createtime    string  `json:"createtime"`
	Deduction     float64 `json:"deduction"`
	Goodsname     string  `json:"goodsname"`
	Goodsnum      int     `json:"goodsnum"`
	Goodsspec     string  `json:"goodsspec"`
	Goodsunit     string  `json:"goodsunit"`
	Goodsweight   float32 `json:"goodsweight"`
	Jid           int     `json:"jid"`
	Pmid          int     `json:"pmid"`
	Source        string  `json:"source"`
	Sourceno      string  `json:"sourceno"`
	Vehicle       string  `json:"vehicle"`
	Vehicleload   string  `json:"vehicleload"`
	Transporttime string  `json:"transporttime"`
	Makeid        string  `json:"makeid"`
	Weight        float64 `json:"weight"`
	Mgh           string  `json:"mgh"`
	Mdd           string  `json:"mdd"`
	Mjsfs         string  `json:"mjsfs"`
	Mshdw         string  `json:"mshdw"`
	Mfhdw         string  `json:"mfhdw"`
}

type WareHouseCheckGoods struct {
	Wid        int     `json:"wid"`        //wid 过磅单ID
	Wno        string  `json:"wno"`        //wno 过磅单号
	Wdriver    string  `json:"wdriver"`    //wdriver 司机姓名
	DriverTel  string  `json:"driver_tel"` //drivertel 电话
	Wvecgucle  string  `json:"wvecgucle"`  //wvecgucle 车牌号
	Wdruverid  int     `json:"wdruverid"`  //wdruverid 司机ID
	OrderNum   string  `json:"ordernum"`   //ordernum 订单号
	Upid       int     `json:"upid"`       //upid 是否可编辑 0-可编辑 1-不能编辑(已完成不能编辑)
	Pmid       int     `json:"pmid"`       //pmid	预约单id
	WGoods     string  `json:"wgoods"`     //wgoods 货品
	Pihao      string  `json:"pihao"`      //pihao 生产批号
	Kou        float32 `json:"kou"`        //kou 扣杂扣重
	Wnum       int     `json:"wnum"`       //wnum 货品数量
	Wspec      float32 `json:"wspec"`      //wspec 货品规格
	Wunit      string  `json:"wunit"`      //wunit 货品单位
	Wweight    float32 `json:"wweight"`    //wweight 货品重量
	Wdeduction float32 `json:"wdeduction"` //wdeduction 每吨扣kg数（吨）
	Wallow     float32 `json:"wallow"`     //wallow 允许额外可配发重量
	Bzdweight  float32 `json:"bzdweight"`  //bzdweight 编织袋重量
	Wtype      string  `json:"wtype"`      //wtype 预约单来源（销售单（XS）和采购单（CG））
	Wbianhao   string  `json:"wbianhao"`   //wbianhao 货品编号
	Wghdw      string  `json:"wghdw"`      //wghdw 供货单位
	Wshdw      string  `json:"wshdw"`      //wshdw 收货单位
	Wjzxgh     string  `json:"wjzxgh"`     //wjzxgh 集装箱柜号
	Wzxhdd     string  `json:"wzxhdd"`     //wzxhdd 装卸货地点
	Wckname    string  `json:"wckname"`    //wckname 仓库人员
	Wxianz     string  `json:"wxianz"`     //wxianz 限重字段
	Wsjzz      string  `json:"wsjzz"`      //wsjzz  车辆的载重字段
}

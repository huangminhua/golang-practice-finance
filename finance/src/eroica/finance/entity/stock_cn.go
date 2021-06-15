package entity

import "encoding/json"

type IsHs string //是否沪深港通标的

const (
	IsHs_N IsHs = "N" //否
	IsHs_H IsHs = "H" //沪股通
	IsHs_S IsHs = "S" //深股通
)

func AllIsHs() []IsHs {
	return []IsHs{IsHs_N, IsHs_H, IsHs_S}
}

type ListStatus string //上市状态

const (
	ListStatus_L ListStatus = "L" //上市
	ListStatus_D ListStatus = "D" //退市
	ListStatus_P ListStatus = "P" //暂停上市
)

func AllListStatus() []ListStatus {
	return []ListStatus{ListStatus_L, ListStatus_D, ListStatus_P}
}

type Exchange string //交易所
const (
	Exchange_SSE  Exchange = "SSE"  //上交所
	Exchange_SZSE Exchange = "SZSE" //深交所
)

func AllExchange() []Exchange {
	return []Exchange{Exchange_SSE, Exchange_SZSE}
}

type StockBasic struct { //基础信息
	TsCode     string     `tushare:"ts_code"`     //TS代码
	Symbol     string     `tushare:"symbol"`      //股票代码
	Name       string     `tushare:"name"`        //股票名称
	Area       string     `tushare:"area"`        //地域
	Industry   string     `tushare:"industry"`    //所属行业
	Fullname   string     `tushare:"fullname"`    //股票全称
	Enname     string     `tushare:"enname"`      //英文全称
	Cnspell    string     `tushare:"cnspell"`     //拼音缩写
	Market     string     `tushare:"market"`      //市场类型（主板/创业板/科创板/CDR）
	Exchange   Exchange   `tushare:"exchange"`    //交易所代码
	CurrType   string     `tushare:"curr_type"`   //交易货币
	ListStatus ListStatus `tushare:"list_status"` //上市状态 L上市 D退市 P暂停上市
	ListDate   string     `tushare:"list_date"`   //上市日期
	DelistDate string     `tushare:"delist_date"` //退市日期
	IsHs       IsHs       `tushare:"is_hs"`       //是否沪深港通标的，N否 H沪股通 S深股通
}

type IsOpen string //是否交易
const (
	IsOpen_0 IsOpen = "0" //休市
	IsOpen_1 IsOpen = "1" //交易
)

func AllIsOpen() []IsOpen {
	return []IsOpen{IsOpen_0, IsOpen_1}
}

type TradeCal struct { //交易日历
	Exchange     Exchange `tushare:"exchange"`      //交易所 SSE上交所 SZSE深交所
	CalDate      string   `tushare:"cal_date"`      //日历日期
	IsOpen       IsOpen   `tushare:"is_open"`       //是否交易 0休市 1交易
	PretradeDate string   `tushare:"pretrade_date"` //上一个交易日
}

type Namechange struct { //股票曾用名
	TsCode       string `tushare:"ts_code"`       //TS代码
	Name         string `tushare:"name"`          //证券名称
	StartDate    string `tushare:"start_date"`    //开始日期
	EndDate      string `tushare:"end_date"`      //结束日期
	AnnDate      string `tushare:"ann_date"`      //公告日期
	ChangeReason string `tushare:"change_reason"` //变更原因
}

type HsType string //沪深港通类型

const (
	HsType_SH HsType = "SH" //沪
	HsType_SZ HsType = "SZ" //深
)

func AllHsType() []HsType {
	return []HsType{HsType_SH, HsType_SZ}
}

type Is string //是否

const (
	Is_0 Is = "0" //否
	Is_1 Is = "1" //是
)

func AllIs() []Is {
	return []Is{Is_0, Is_1}
}

type HsConst struct { //沪深股通成份股
	TsCode  string `tushare:"ts_code"`  //TS代码
	HsType  HsType `tushare:"hs_type"`  //沪深港通类型SH沪SZ深
	InDate  string `tushare:"in_date"`  //纳入日期
	OutDate string `tushare:"out_date"` //剔除日期
	IsNew   Is     `tushare:"is_new"`   //是否最新 1是 0否
}

type StockCompany struct { //上市公司基本信息
	TsCode        string      `tushare:"ts_code"`        //TS代码
	Exchange      Exchange    `tushare:"exchange"`       //交易所代码
	Chairman      string      `tushare:"chairman"`       //法人代表
	Manager       string      `tushare:"manager"`        //总经理
	Secretary     string      `tushare:"secretary"`      //董秘
	RegCapital    json.Number `tushare:"reg_capital"`    //注册资本
	SetupDate     string      `tushare:"setup_date"`     //注册日期
	Province      string      `tushare:"province"`       //所在省份
	City          string      `tushare:"city"`           //所在城市
	Introduction  string      `tushare:"introduction"`   //公司介绍
	Website       string      `tushare:"website"`        //公司主页
	Email         string      `tushare:"email"`          //电子邮件
	Office        string      `tushare:"office"`         //办公室
	Employees     int         `tushare:"employees"`      //员工人数
	MainBusiness  string      `tushare:"main_business"`  //主要业务及产品
	BusinessScope string      `tushare:"business_scope"` //经营范围
}

type StkManagers struct { //上市公司管理层
	TsCode    string `tushare:"ts_code"`    //TS代码
	AnnDate   string `tushare:"ann_date"`   //公告日期
	Name      string `tushare:"name"`       //姓名
	Gender    string `tushare:"gender"`     //性别
	Lev       string `tushare:"lev"`        //岗位类别
	Title     string `tushare:"title"`      //岗位
	Edu       string `tushare:"edu"`        //学历
	National  string `tushare:"national"`   //国籍
	Birthday  string `tushare:"birthday"`   //出生年月
	BeginDate string `tushare:"begin_date"` //上任日期
	EndDate   string `tushare:"end_date"`   //离任日期
	Resume    string `tushare:"resume"`     //个人简历
}

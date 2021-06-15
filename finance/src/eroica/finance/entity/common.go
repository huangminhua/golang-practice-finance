package entity

import (
	"fmt"
	"reflect"
	"strings"
)

// 实体类型、tushare api接口和数据库表的对应关系 结构体
type EntityMapping struct {
	Type    reflect.Type //struct
	ApiName string       //tushare
	Table   string       //db
}

// 获取实体类型、tushare api接口和数据库表的对应关系。
// 输入:
// t reflect.Type 本包中定义的实体类型，如reflect.TypeOf(StockBasic{})
// 输出:
// EntityMapping 实体类型、tushare api接口和数据库表的对应关系
func GetEntityMapping(t reflect.Type) EntityMapping {
	switch t {
	case reflect.TypeOf(StockBasic{}): //基础信息
		return EntityMapping{Type: t, ApiName: "stock_basic", Table: "stock_cn.stock_basic"}
	case reflect.TypeOf(TradeCal{}): //交易日历
		return EntityMapping{Type: t, ApiName: "trade_cal", Table: "stock_cn.trade_cal"}
	case reflect.TypeOf(Namechange{}): //股票曾用名
		return EntityMapping{Type: t, ApiName: "namechange", Table: "stock_cn.namechange"}
	case reflect.TypeOf(HsConst{}): //沪深股通成份股
		return EntityMapping{Type: t, ApiName: "hs_const", Table: "stock_cn.hs_const"}
	case reflect.TypeOf(StockCompany{}): //上市公司基本信息
		return EntityMapping{Type: t, ApiName: "stock_company", Table: "stock_cn.stock_company"}
	case reflect.TypeOf(StkManagers{}): //上市公司管理层
		return EntityMapping{Type: t, ApiName: "stk_managers", Table: "stock_cn.stk_managers"}
	default:
		panic(fmt.Sprintf("Unsupported data mapping: %v\n", t))
	}
}

// 根据实体类型、tushare api接口和数据库表的对应关系，获取tushare某一接口字段总和的字符串。
// 输入:
// t *EntityMapping 实体类型、tushare api接口和数据库表的对应关系
// 输出:
// string t所对应tushare接口的字段总和的字符串，以,分隔
func (em *EntityMapping) RequestFields() string {
	var sb strings.Builder
	for i := 0; i < em.Type.NumField(); i++ {
		f := em.Type.Field(i)
		t := f.Tag
		ts, ok := t.Lookup("tushare")
		if !ok {
			panic(fmt.Sprintf("Tag tushare of %v.%v is undefined.\n", t, f))
		}
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(ts)
	}
	return sb.String()
}

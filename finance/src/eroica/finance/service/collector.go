package service

import (
	"reflect"
	"strconv"

	da "eroica.finance/src/eroica/finance/data_access"
	"eroica.finance/src/eroica/finance/entity"
	"eroica.finance/src/eroica/finance/tushare"
)

func commonCollect(em entity.EntityMapping, params map[string]interface{}) {
	offset := 0
	for {
		params["offset"] = offset
		trd := tushare.Fetch(em, params)
		da.SaveTxless(em, trd)
		if trd.HasMore {
			offset += len(trd.Items)
		} else {
			break
		}
	}
}

// 调用tushare stock_basic接口获取股票基本信息，并全量更新。
func CollectStockBasic() {
	em := entity.GetEntityMapping(reflect.TypeOf(entity.StockBasic{}))
	da.DeleteAllTxless(em)
	for _, ls := range entity.AllListStatus() {
		params := map[string]interface{}{"list_status": string(ls)}
		commonCollect(em, params)
	}
}

func CollectTradeCal() {
	em := entity.GetEntityMapping(reflect.TypeOf(entity.TradeCal{}))
	da.DeleteAllTxless(em)
	commonCollect(em, map[string]interface{}{"is_open": string(entity.IsOpen_1)})
}

func CollectNamechange() {
	em := entity.GetEntityMapping(reflect.TypeOf(entity.Namechange{}))
	da.DeleteAllTxless(em)
	commonCollect(em, map[string]interface{}{})
}

func CollectHsConst() {
	em := entity.GetEntityMapping(reflect.TypeOf(entity.HsConst{}))
	da.DeleteAllTxless(em)
	for _, ht := range entity.AllHsType() {
		for _, in := range entity.AllIs() {
			commonCollect(em, map[string]interface{}{"hs_type": ht, "is_new": in})
		}
	}
}

func CollectStockCompany() {
	em := entity.GetEntityMapping(reflect.TypeOf(entity.StockCompany{}))
	da.DeleteAllTxless(em)
	for _, e := range entity.AllExchange() {
		commonCollect(em, map[string]interface{}{"exchange": e})
	}
}

func CollectStkManagers() {
	em := entity.GetEntityMapping(reflect.TypeOf(entity.StkManagers{}))
	_, d := da.SelectTxless("select max(ann_date) date from stock_cn.stk_managers")
	if d[0][0] == nil {
		commonCollect(em, map[string]interface{}{})
	} else {
		lastDate, err := strconv.Atoi(*d[0][0])
		if err != nil {
			panic(err)
		}
		dateThreshold := strconv.Itoa(lastDate - 10000)
		da.DeleteTxless(em, "where ann_date >= ?", dateThreshold)
		commonCollect(em, map[string]interface{}{"start_date": dateThreshold})
	}
}

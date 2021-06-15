package tushare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"eroica.finance/src/eroica/finance/conf"
	"eroica.finance/src/eroica/finance/entity"
)

// tushare请求结构体
type tushareRequest struct {
	ApiName string                 `json:"api_name"`
	Fields  string                 `json:"fields"`
	Params  map[string]interface{} `json:"params"`
	Token   string                 `json:"token"`
}

// tushare返回结构体
type tushareResponse struct {
	RequestId string     `json:"request_id"`
	Code      int        `json:"code"`
	Msg       string     `json:"msg"`
	Data      TsRespData `json:"data"`
}

// tushare数据结构体
type TsRespData struct {
	Fields  []string        `json:"fields"`
	Items   [][]interface{} `json:"items"`
	HasMore bool            `json:"has_more"`
}

var config = conf.GetConf()

// 请求tushare接口，获取数据。
// 输入:
// em EntityMapping 实体类型、tushare api接口和数据库表的对应关系，由entity.GetEntityMapping(t reflect.Type)获取
// params map[string]interface{} 搜索条件
// 输出:
// TsRespData tushare数据结构体
func Fetch(em entity.EntityMapping, params map[string]interface{}) TsRespData {
	reqBody := tushareRequest{
		ApiName: em.ApiName,
		Fields:  em.RequestFields(),
		Params:  params,
		Token:   config.Tushare.Token,
	}
	reqBodyJson, err := json.Marshal(&reqBody)
	if err != nil {
		panic(err)
	}
	urlObj, err := url.Parse(config.Tushare.Host)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(urlObj.String(), "application/json", bytes.NewReader(reqBodyJson))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	respBodyJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var respBody tushareResponse
	decoder := json.NewDecoder(bytes.NewReader(respBodyJson))
	decoder.UseNumber()
	err = decoder.Decode(&respBody)
	if err != nil {
		panic(err)
	}
	if respBody.Code != 0 {
		panic(fmt.Sprintf("Fetch %v with condition %v failed. Response is %v.\n", em.ApiName, params, respBodyJson))
	}
	log.Printf("[INFO] Fetch %v with condition %v success.\n", em.ApiName, params)
	return respBody.Data
}

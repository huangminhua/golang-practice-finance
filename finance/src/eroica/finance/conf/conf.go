package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v2"
)

// 本项目配置结构体
type Conf struct {
	Tushare struct {
		Host  string `yaml:"host"`
		Token string `yaml:"token"`
	} `yaml:"tushare"`
	Db struct {
		Driver       string `yaml:"driver"`
		DataSource   string `yaml:"data-source"`
		MaxOpenConns int    `yaml:"max-open-conns"`
		MaxIdleConns int    `yaml:"max-idle-conns"`
	} `yaml:"db"`
}

var conf Conf
var loadOnce sync.Once

// 获取本项目配置
func GetConf() Conf {
	loadOnce.Do(func() {
		confYml, err := ioutil.ReadFile("conf.yml")
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(confYml, &conf)
		if err != nil {
			panic(err)
		}
		confJson, err := json.Marshal(&conf)
		if err != nil {
			panic(err)
		}
		log.Println("[INFO] conf.yml loaded. The configuration is", string(confJson))
	})
	return conf
}

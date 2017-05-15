package conf

import (
    "strings"
    "github.com/Unknwon/goconfig"
)

type Db struct {
   Host string
   Port string
   User string
   Password string
   Dbname string
}

var (
    DbMap map[string]Db
    CompareMap map[string]string
    DingDingUrl string
)

func LoadConfigFile(confPath string) {
    cfg, err := goconfig.LoadConfigFile(confPath)
    if err != nil {
        panic("load config failed")
    }

    mysqls, err := cfg.GetValue("mysqls", "host")
    if err != nil {
        panic("get mysqls failed")
    }

    /* 处理mysqls分别获取各自配置 */
    mysqlList := strings.Split(mysqls, "|")
    if len(mysqlList) < 2 {
        panic("mysql less then 2")
    }

    DbMap = make(map[string]Db)
    CompareMap = make(map[string]string)
    for _, _v := range mysqlList {
       var tmpDb Db
       tmpDb.Host, err = cfg.GetValue(_v, "host")
       tmpDb.Port, err = cfg.GetValue(_v, "port")
       tmpDb.User, err = cfg.GetValue(_v, "user")
       tmpDb.Password, err = cfg.GetValue(_v, "password")
       tmpDb.Dbname, err = cfg.GetValue(_v, "dbname")
       DbMap[_v] = tmpDb

       target, err := cfg.GetValue("diff", _v)
       if err == nil  {
           CompareMap[_v] = target
       }
    }

    DingDingUrl, _ = cfg.GetValue("dingding", "url")
}

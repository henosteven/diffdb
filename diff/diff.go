package main

import (
    "fmt"
    "conf"
    "notify"
    "dbhandler"
    "database/sql"
)

func main() {
    fmt.Println("app start")
    conf.LoadConfigFile("../conf/diff.ini")
    var connMap map[string]*sql.DB
    connMap = dbhandler.InitConn(conf.DbMap)

    var dbstructMap map[string]string
    dbstructMap = dbhandler.GetAllStruct(connMap)

    for k, v := range conf.CompareMap {
       if dbstructMap[k] != dbstructMap[v] {
           fmt.Println("different")
           notify.SendDingDing("different")
       } else {
           fmt.Println("same")
           notify.SendDingDing("same")
       }
    }
}

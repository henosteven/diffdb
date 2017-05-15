package main

import (
    "fmt"
    "conf"
    //"notify"
    "dbhandler"
    "database/sql"
    "regexp"
)

func main() {
    fmt.Println("app start")
    conf.LoadConfigFile("../conf/diff.ini")
    var connMap map[string]*sql.DB
    connMap = dbhandler.InitConn(conf.DbMap)

    var dbstructMap map[string]interface{}
    dbstructMap = dbhandler.GetAllStruct(connMap)

    var difftable []string
    for k, v := range conf.CompareMap {
        tmpMap := dbstructMap[k].(map[string]string)
        tmpTargetMap := dbstructMap[v].(map[string]string)
        for  _table, _tablesrtuct := range tmpMap {
            if diffStruct(_tablesrtuct, tmpTargetMap[_table]) {
                difftable = append(difftable, _table)
            }
        }
    }

    if difftable != nil {
        fmt.Println(difftable)
        notify.SendDingDing(difftable)
    }
}

func diffStruct(table1 string, table2 string) bool {
    var result bool = true
    if conf.IgnoreAutoIncrement {
        reg := regexp.MustCompile("AUTO_INCREMENT=\\d")
        table1 = reg.ReplaceAllString(table1, "")
        table2 = reg.ReplaceAllString(table2, "")
    }

    if table1 == table2 {
        result = false
    }

    return result
}

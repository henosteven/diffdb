package dbhandler

import (
    "conf"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
)

func InitConn(dbMap map[string]conf.Db) map[string]*sql.DB{
    var connMap = make(map[string]*sql.DB)
    for k, tmpconf := range dbMap  {
        dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", tmpconf.User, tmpconf.Password, tmpconf.Host, tmpconf.Port, tmpconf.Dbname) 
        db, err := sql.Open("mysql", dsn) 
        if err != nil {
            panic("failed to open db\n")
        }

        err = db.Ping()
        if err != nil {
            panic("failed to connect to db\n")
        }
        connMap[k] = db
    }
    return connMap
}

func GetAllStruct(connMap map[string]*sql.DB) map[string]interface{}{
    var dbstructMap = make(map[string]interface{})
    for k, v := range connMap {
       dbstructMap[k] = getDBStruct(v) 
    }
    return dbstructMap
}

func getDBStruct(conn *sql.DB) map[string]string {
    var tablestruct = make(map[string]string) 
    rows, err := conn.Query("show tables")
    if err != nil {
        fmt.Println("query failed")
    }
    defer rows.Close()

    var tableList []string 
    for rows.Next() {
        var tmpTable string
        err := rows.Scan(&tmpTable)
        if err != nil {
            fmt.Println("scan failed")
        }
        tableList = append(tableList, tmpTable)
    }

    for _, v := range tableList {
        var tmpstr string
        tmpstr = getTableStruct(conn, v)
        tablestruct[v] = tmpstr
    }

    return tablestruct
}

func getTableStruct(conn *sql.DB, table string) string {
    var tablestring string
    rows, err := conn.Query("show create table " + table)
    if err != nil {
        fmt.Println("query failed")
    }

    defer rows.Close()

    for rows.Next() {
        var tmptable string
        err := rows.Scan(&tmptable, &tablestring)
        if err != nil {
            fmt.Println("scan table failed")
        }
    }
    return tablestring
}

package utils

import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"github.com/gohouse/gorose"
	"github.com/cihub/seelog"
)

var SqlDB *sql.DB

func _init()  {
	SqlDB,err1 := sql.Open("mysql","root:1q2w3e4r@tcp(45.77.82.85:3306)/jpress")
	SqlDB.SetMaxIdleConns(2)
	SqlDB.SetMaxOpenConns(10)
	err2 := SqlDB.Ping()
	if err1 != nil {log.Fatal(err1)}
	if err2 != nil {log.Fatal(err2)}
}
func loadCon() *gorose.Connection {
	var dbConfig = &gorose.DbConfigSingle{
		Driver:          "mysql", // driver: mysql/sqlite/oracle/mssql/postgres
		EnableQueryLog:  true,    // if enable sql logs
		SetMaxOpenConns: 0,       // connection pool of max Open connections, default zero
		SetMaxIdleConns: 0,       // connection pool of max sleep connections
		Prefix:          "",      // prefix of table
		Dsn:             "root:root@tcp(localhost:3306)/test?charset=utf8", // db dsn
	}
	var connection,err = gorose.Open(dbConfig)
	if err != nil {
		seelog.Info("mysql connection error=",err)
	}
	return connection
}
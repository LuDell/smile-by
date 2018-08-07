package utils

import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
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

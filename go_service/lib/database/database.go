package database

import (
	"fmt"

	"github.com/gohouse/gorose/v2"

	// "github.com/go-redis/redis"
	G "go_api/lib/global"

	_ "github.com/go-sql-driver/mysql"
)

// var db *sql.DB
var err error
var engin *gorose.Engin

//Init 数据库初始化
func Init() {
	conn := fmt.Sprintf("%s:%s@%s?parseTime=true", G.Config("mysql", "username"), G.Config("mysql", "password"), G.Config("mysql", "host"))
	engin, err = gorose.Open(&gorose.Config{Driver: "mysql", Dsn: conn, SetMaxOpenConns: 10})
	return
}

//DB 获取数据库连接
func DB() gorose.IOrm {
	if engin == nil {
		Init()
	}
	return engin.NewOrm()
}

//DBT 带表名的数据库连接
func DBT(table string) gorose.IOrm {
	if engin == nil {
		Init()
	}
	e := engin.NewOrm()
	if len(table) > 0 {
		e.Table(table)
	}
	return e
}

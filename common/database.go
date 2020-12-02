package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

// 初始化数据库
func InitDB(){
	driverName:="mysql"
	host:="127.0.0.1"
	port:="3306"
	database:="bookstore"
	username:="root"
	password:="believe9407"
	charset:="utf8"

	dsn:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s&parseTime=true",
		username,password,host,port,database,charset)

	// 连接数据库
	db,err:=sql.Open(driverName,dsn)
	if err!=nil{
		panic(err.Error())
	}
	DB=db
}

// 获取DB
func GetDB()*sql.DB{
	return DB
}



package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	ConnHost       = "localhost"
	ConnPort       = "3306"
	DriverName     = "mysql"
	DataSourceName = "root:root@/qudash"
)

var db *sql.DB
var connectionError error

//FOR RUN MODE DEV
/*var (
	mysqluser      = beego.AppConfig.String("dev::mysqluser")
	mysqlpass      = beego.AppConfig.String("dev::mysqlpass")
	mysqldb        = beego.AppConfig.String("dev::mysqldb")
	DataSourceName = mysqluser + ":" + mysqlpass + "@/" + mysqldb
)*/

/*var (
	mysqluser      = beego.AppConfig.String("mysqluser")
	mysqlpass      = beego.AppConfig.String("mysqlpass")
	mysqldb        = beego.AppConfig.String("mysqldb")
	DataSourceName = mysqluser + ":" + mysqlpass + "@/" + mysqldb
)*/

func Connect() {
	log.Println("Datasource:: " + DataSourceName)
	log.Println("Driver:: " + DriverName)
	db, connectionError = sql.Open(DriverName, DataSourceName)
	if connectionError != nil {
		log.Fatal("error connecting to database : ", connectionError)
	}
}

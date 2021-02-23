package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/namsral/flag"
)

var DB *gorm.DB
var err error

func InitDB() {

	var username string
	var userpass string
	var dbname string
	var host string
	var port string
	flag.StringVar(&username, "username", "", "User")
	flag.StringVar(&userpass, "userpass", "", "Pass")
	flag.StringVar(&dbname, "dbname", "", "Db Name")
	flag.StringVar(&host, "dbhost", "", "Db host")
	flag.StringVar(&port, "dbport", "", "Db port")
	flag.Parse()
	db_config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", username, userpass, host, port, dbname)
	DB, err = gorm.Open("mysql", db_config)
	err = DB.DB().Ping()
	if err != nil {
		panic(err)
	}
	DB.LogMode(true)
	dur := time.Duration(0)
	DB.DB().SetConnMaxLifetime(dur)
}

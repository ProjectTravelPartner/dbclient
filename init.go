package dbclient

import (
	"database/sql"
	"fmt"
	"time"

	toml "github.com/pelletier/go-toml"
)

var dbGlobal *sql.DB

func Init() {
	const dbdetails = "conf/dbdetails.toml"
	initialiseDB(dbdetails)
}

func initialiseDB(dbdetails string) {
	var conf *toml.Tree
	var err error
	if conf, err = toml.LoadFile(dbdetails); err != nil {
		fmt.Println("error loading db config")
	}

	dbDriver := conf.Get("database.driver").(string)
	dbDatasource := conf.Get("database.dataSource").(string)

	var db *sql.DB

	if db, err = sql.Open(dbDriver, dbDatasource); err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 10; i++ {
		err := db.Ping()
		if err == nil {
			fmt.Println("DB ping success")
			break
		}
		time.Sleep(5 * time.Second)
	}
	dbGlobal = db
}

func Close() {
	dbGlobal.Close()
}

func GetDB() *sql.DB {
	return dbGlobal
}

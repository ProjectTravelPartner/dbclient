package dbclient

import "database/sql"

func Query(qry string, args ...interface{}) (*sql.Rows, error) {
	db := GetDB()
	return db.Query(qry, args...)
}

package dbclient

import "database/sql"

func Query(qry string, args ...interface{}) (*sql.Rows, error) {
	db := GetDB()
	return db.Query(qry, args...)
}

func QueryRow(qry string, args ...interface{}) *sql.Row {
	db := GetDB()
	return db.QueryRow(qry, args...)
}

func ExecGetID(qry string, args ...interface{}) (uint64, error) {
	db := GetDB()
	var err error
	var res sql.Result
	if res, err = db.Exec(qry, args...); err != nil {
		return 0, err
	}
	var id int64
	id, err = res.LastInsertId()
	return uint64(id), err
}

func ExecGetAffRows(qry string, args ...interface{}) (uint64, error) {
	db := GetDB()
	var err error
	var res sql.Result
	if res, err = db.Exec(qry, args...); err != nil {
		return 0, err
	}
	var id int64
	id, err = res.RowsAffected()
	return uint64(id), err
}

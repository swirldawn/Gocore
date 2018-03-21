package gocore

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitMysql() {

	db_user := GetConfig("database", "DB_USERNAME")
	db_password := GetConfig("database", "DB_PASSWORD")
	db_host := GetConfig("database", "DB_HOST")
	db_name := GetConfig("database", "DB_DATABASE")

	var err error

	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", db_user, db_password, db_host, db_name))

	if err != nil {
		panic(err.Error())
	}
}

//插入
func insert(sqlstr string, args ...interface{}) (int64, error) {

	stmtIns, err := DB.Prepare(sqlstr)

	if err != nil {
		panic(err.Error())
	}

	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)

	if err != nil {
		panic(err.Error())
	}

	return result.LastInsertId()

}

//修改和删除
func exec(sqlstr string, args ...interface{}) (int64, error) {

	stmtIns, err := DB.Prepare(sqlstr)

	if err != nil {
		panic(err.Error())
	}

	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)

	if err != nil {
		panic(err.Error())
	}

	return result.RowsAffected()
}

func fetchOne(sqlstr string, args ...interface{}) (string, error) {
	stmtOut, err := DB.Prepare(sqlstr)

	if err != nil {
		panic(err.Error())
	}

	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)

	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()

	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	var ret string

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string

		for _, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			ret = value
			break
		}
		break //get the first row only
	}

	return ret, nil
}

//取一行数据，注意这类取出来的结果都是string
func fetchRow(sqlstr string, args ...interface{}) (map[string]string, error) {

	stmtOut, err := DB.Prepare(sqlstr)

	if err != nil {
		panic(err.Error())
	}

	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)

	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()

	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	ret := make(map[string]string, len(scanArgs))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string

		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			ret[columns[i]] = value
		}
		break //get the first row only
	}

	return ret, nil
}

//取多行，<span style="font-family: Arial, Helvetica, sans-serif;">注意这类取出来的结果都是string </span>
func fetchAll(sqlstr string, args ...interface{}) ([]map[string]string, error) {

	stmtOut, err := DB.Prepare(sqlstr)

	if err != nil {
		panic(err.Error())
	}

	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)

	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()

	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := make([]map[string]string, 0)

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {

		err = rows.Scan(scanArgs...)

		if err != nil {
			panic(err.Error())
		}

		var value string

		vmap := make(map[string]string, len(scanArgs))

		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			vmap[columns[i]] = value
		}
		ret = append(ret, vmap)
	}

	return ret, nil
}

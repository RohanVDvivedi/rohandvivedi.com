package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var mysqlDatabase *sql.DB = nil

func Initialize() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/dbname?parseTime=true");
	if(err != nil) {
		panic(err.Error());
	} else {
		mysqlDatabase = db;
	}
}

func QueryDatabase(nativeQuery string, scanOperation func(*sql.Rows)(interface{}) ) (interface{}) {
	result, err := mysqlDatabase.Query(nativeQuery);
	if(err != nil) {
		panic(err.Error());
	} else {
		defer result.Close()
		if(scanOperation != nil) {
			return scanOperation(result);
		} else {
			return nil;
		}
	}
}

func Close() {
	mysqlDatabase.Close();
}
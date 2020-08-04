package data

import (
    "database/sql"
)

var Db *sql.DB = nil;

type Row interface {
    Scan(...interface{}) error
}
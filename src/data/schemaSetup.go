package data

import (
    "database/sql"
)

func InitializeSchema(db *sql.DB) {
    statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
    statement.Exec()
}


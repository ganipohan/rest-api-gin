package models

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var db *sql.DB

// InitDB menginisialisasi koneksi ke database MySQL
func InitDB(dataSourceName string) {
    var err error
    db, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
}

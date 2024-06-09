package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Conn() {
	db, _ := sql.Open("mysql", "root:@tcp/crud_pasien")

	if err := db.Ping(); err != nil {
		panic(err)
	}

	print("Berhasil Terkoneksi Ke Database")

	DB = db

}

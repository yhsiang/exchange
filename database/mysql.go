package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// type Database interface {
// 	New() *sql.DB
// }

func New() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	db, err := sql.Open("mysql", "root:"+pass+"@tcp("+host+":3306)/sbi_exchange?parseTime=true")
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

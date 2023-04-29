package mysql

import (
	"database/sql"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/oldme")
	if err != nil {
		return nil, err
	}
	return db, err
}

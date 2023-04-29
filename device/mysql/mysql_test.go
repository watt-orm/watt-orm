package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestOpen(t *testing.T) {
	db, err := Open()
	if err != nil {
		panic(err)
	}
	t.Log(db)
}

func TestQuery(t *testing.T) {
	//db, err := sql.Open("mysql", "root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/oldme")

}

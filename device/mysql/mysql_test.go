package mysql

import "testing"

func TestOpen(t *testing.T) {
	db, err := Open()
	if err != nil {
		panic(err)
	}
	t.Log(db)
}

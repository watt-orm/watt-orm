package test

import (
	"testing"
	"watt-orm"
	_ "watt-orm/device/mysql"
)

func getDb() (watt_orm.DB, error) {
	conf := watt_orm.GetDbConf()
	return watt_orm.New(conf)
}

func TestNewDb(t *testing.T) {
	_, err := getDb()
	if err != nil {
		t.Error(err)
	}
}

func TestFind(t *testing.T) {

}

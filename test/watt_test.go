package test

import (
	"testing"
	"watt-orm"
	_ "watt-orm/device/mysql"
)

func TestNewDb(t *testing.T) {
	conf := watt_orm.GetDbConf()
	_, err := watt_orm.New(conf)
	if err != nil {
		panic(err)
	}
}

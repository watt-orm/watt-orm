package watt_orm

import (
	"database/sql"
	"fmt"
	"watt-orm/werr"
)

func New(conf *Conf) (db DB, err error) {
	core := &Core{
		debug: conf.Debug,
		conf:  conf,
	}
	if v, ok := DriverList[conf.Driver]; ok {
		if core.db, err = v.New(core, conf); err != nil {
			return nil, err
		}
		return core.db, nil
	}
	var (
		errMsg = fmt.Sprintf("%s驱动未注册", conf.Driver)
	)
	return nil, werr.New(errMsg)
}

type DB interface {
	Open(conf *Conf) (*sql.DB, error)

	//Close()
	// Model 当前模型
	//Model() *Model
}

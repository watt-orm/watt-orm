package watt_orm

import (
	"context"
	"database/sql"
	"fmt"
	"watt-orm/werr"
)

type DB interface {
	Open(conf *Conf) (*sql.DB, error)

	// Model 当前模型
	Model(tableName string) *Model

	//Insert(ctx context.Context, table string, condition string)
	//Update(ctx context.Context)
	Select(ctx context.Context, table string, condition string) (*sql.Rows, error)
	//Delete(ctx context.Context)

	//Prepare()
	//Exec()
}

// New 创建并返回一个新的DB
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

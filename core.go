package watt_orm

import "context"

type Core struct {
	db    DB
	ctx   context.Context
	debug bool
	conf  *Conf
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) Model(tableName string) *Model {
	return &Model{}
}

package watt_orm

import (
	"context"
	"database/sql"
)

func (c *Core) Select(ctx context.Context, table string, condition string) (*sql.Rows, error) {
	return nil, nil
}

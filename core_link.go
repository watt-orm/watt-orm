package watt_orm

import "database/sql"

type DBLink struct {
	*sql.DB
}

package watt_orm

type DB struct {
	// 当前表名
	TableName string
	// 当前模型
	Model *Model
}

// Table 注入当前表名
func (d *DB) Table(name string) *Model {
	d.TableName = name
	return &Model{
		DB: d,
	}
}

// Commit 执行SQL语句
func (d *DB) Commit(sql string) {

}

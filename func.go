package watt_orm

// New 创建一个Db
func New() *DB {
	return &DB{
		TableName: "",
		Model:     nil,
	}
}

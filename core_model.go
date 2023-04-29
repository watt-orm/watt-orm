package watt_orm

import "fmt"

type Model struct {
	DB *DB
}

// Get 获取全部数据
func (m *Model) Get() {
	sql := fmt.Sprintf("select * from %s", m.DB.TableName)
	fmt.Println(sql)
}

func (m *Model) Where(where string, args interface{}) *Model {
	return nil
}

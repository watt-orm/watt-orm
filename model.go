package watt_orm

import "fmt"

type Model struct {
}

// Get 获取全部数据
func (m *Model) Get() {
	sql := fmt.Sprintf("select * from %s", "s")
	fmt.Println(sql)
}

func (m *Model) Where(where string, args interface{}) *Model {
	return nil
}

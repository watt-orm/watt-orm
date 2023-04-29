package test

import (
	"testing"
	w "watt-orm"
)

// 使用Db.Table创建一个Model
func TestCreModel(t *testing.T) {
	w.New().Table("article").Get()
}

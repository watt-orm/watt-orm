package mysql

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	watt "watt-orm"
)

type Driver struct {
	*watt.Core
}

// 将mysql系列的数据库注册到driverList
func init() {
	var (
		driverName = [3]string{"mysql", "mariadb", "tidb"}
	)
	for _, v := range driverName {
		watt.Register(v, &Driver{})
	}
}

// New 实例化DB,实现惰性加载
func (w *Driver) New(core *watt.Core, conf *watt.Conf) (watt.DB, error) {
	return &Driver{
		watt.NewCore(),
	}, nil
}

func (w *Driver) Open(conf *watt.Conf) (*sql.DB, error) {
	var (
		source string
	)
	source = fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=%s",
		conf.User, conf.Pass, conf.Protocol, conf.Host, conf.Port, conf.Name, conf.Charset)
	if conf.Timezone != "" {
		if strings.Contains(conf.Timezone, "/") {
			conf.Timezone = url.QueryEscape(conf.Timezone)
		}
		source = fmt.Sprintf("%s&loc=%s", source, conf.Timezone)
	}
	if conf.Extra != "" {
		source = fmt.Sprintf("%s&%s", source, conf.Extra)
	}
	db, err := sql.Open(conf.Driver, source)
	if err != nil {
		return nil, err
	}
	return db, err
}

package watt_orm

type Conf struct {
	Protocol string // 协议
	Timezone string // 时区
	Host     string // 主机
	Port     string // 端口
	Name     string // 数据库名称
	Charset  string // 数据库编码
	User     string // 用户名
	Pass     string // 密码
	Driver   string // 驱动类型
	Prefix   string // 表前缀
	Extra    string // 各种不同的数据库驱动额外配置
	Debug    bool   // Debug 模式
}

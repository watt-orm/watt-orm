package watt_orm

// GetDbConf 从配置文件中读取数据库相关配置
func GetDbConf() *Conf {
	return &Conf{
		Protocol: "tcp",
		Timezone: "Asia/Shanghai",
		Host:     "192.168.10.49",
		Port:     "3306",
		Name:     "oldme",
		Charset:  "utf8mb4,utf8",
		User:     "root",
		Pass:     "JCPHqknyy8ATR5ME",
		Driver:   "mysql",
		Extra:    "allowAllFiles=true",
		Debug:    false,
	}
}

package watt_orm

var (
	DriverList = make(map[string]Driver)
)

type Driver interface {
	New(core *Core, conf *Conf) (DB, error)
}

// Register 注册驱动
func Register(name string, driver Driver) {
	DriverList[name] = driver
}

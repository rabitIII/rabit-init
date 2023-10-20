package config

import "fmt"

type Mysql struct {
	Host     string `yaml:"host"`     // 服务器地址
	Port     int    `yaml:"port"`     // 服务器地址的端口
	Config   string `yaml:"config"`   // 高级配置
	DB       string `yaml:"db"`       // 所连接数据库的库名
	Username string `yaml:"username"` // 连接的用户名
	Password string `yaml:"password"` // 密码
	LogLevel string `yaml:"logLevel"` // 是否开启Gorm全局日志
}

func (m *Mysql) Dsn() string {

	// "username:password@tcp(x.x.x.x:post)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.DB, m.Config)
}

//func (m *Mysql) GetLogMode() string {
//	// 操作数据库的行为日志记录
//	return m.LogLevel
//}

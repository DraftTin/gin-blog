package config

import "strconv"

type MySQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Config   string `yaml:"config"` // 高级配置: 例如charset
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` // output all sql when debugging
}

func (m *MySQL) GetDSN() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.Db + "?" + m.Config
}

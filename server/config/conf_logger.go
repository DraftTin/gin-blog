package config

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     string `yaml:"showline"`       // show line number
	LogInConsole string `yaml:"log_in_console"` // show log path
}

package response

type ErrorCode int

const (
	SettingError ErrorCode = iota
)

var (
	ErrorMsg = map[ErrorCode]string{
		SettingError: "系统错误",
	}
)

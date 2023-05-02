package templates

var ErrorxTempalte = `package errorx

type ErrorCode int


const (
	SuccessCode   ErrorCode = 1000 // 成功
	SettingsError ErrorCode = 1001 //系统错误
	ArgumentError ErrorCode = 1002 //参数错误
	FailedCode    ErrorCode = 1999 // 返回失败
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
		ArgumentError: "参数错误",
		SuccessCode:   "成功",
		FailedCode:    "失败",
	}
)`

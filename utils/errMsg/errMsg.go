package errMsg

const (
	SUCCESS = 200
	ERROR = 500

	// ErrorUsernameUsed code = 1000... 用户模块的错误
	ErrorUsernameUsed    = 1001
	ErrorPasswordWrong = 1002
	ErrorUserNotExist = 1003
	ErrorTokenExist   = 1004
	ErrorTokenRuntime      = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007
	ErrorUserNoRight = 1008
	// code = 2000... 文章模块的错误

	// ErrorCateNameUsed  code = 3000... 分类模块的错误
	ErrorCateNameUsed = 3001
)

var CodeMsg = map[int]string{
	SUCCESS:             "OK",
	ERROR:               "FAIL",
	ErrorUsernameUsed:   "用户名已存在",
	ErrorPasswordWrong:  "密码错误",
	ErrorUserNotExist:   "用户不存在",
	ErrorTokenExist:     "TOKEN不存在",
	ErrorTokenRuntime:   "TOKEN已过期",
	ErrorTokenWrong:     "TOKEN不正确",
	ErrorTokenTypeWrong: "TOKEN格式错误",
	ErrorCateNameUsed: "分类已存在",
	ErrorUserNoRight: "该用户不具备管理员权限",
}


func GetErrMsg(code int) string {


	return CodeMsg[code]
}
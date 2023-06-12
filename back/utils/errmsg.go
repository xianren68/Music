package utils

// 定义一些常量，用作错误状态码
const (
	SUCCESS = 200
	ERROR   = 500
	// 关于用户数据的错误 1000
	ERROR_USER_USED      = 1001
	ERROR_USER_NOT_EXIST = 1002
	ERROR_PASSWORD_WRONG = 1003
	ERROR_USER_NO_RIGHT  = 1004
)

// 建立错误信息与错误码的映射表
var CodeMsg = map[int]string{
	SUCCESS:              "ok",
	ERROR:                "fail",
	ERROR_USER_USED:      "用户名被占用",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_USER_NO_RIGHT:  "该用户无权限",
	ERROR_PASSWORD_WRONG: "密码有误",
}

package controller

// 定义错误码
type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeError
	CodeInvalidParam
	CodeInvalidToken
	CodeUserNameExist
	CodeUserNotExist
	CodeUserPasswordError
	CodeServerBusy
	CodeSignUpFailed
	CodeNeedLogin
)

// 错误码与错误信息映射
var codeMap = map[ResCode]string{
	CodeSuccess:           "success",
	CodeError:             "error",
	CodeInvalidParam:      "invalid param",
	CodeInvalidToken:      "invalid token",
	CodeUserNameExist:     "user name exist",
	CodeUserNotExist:      "user not exist",
	CodeUserPasswordError: "username or password is error",
	CodeServerBusy:        "server is busy",
	CodeSignUpFailed:      "sign up failed",
	CodeNeedLogin:         "need login",
}

// 获取错误信息
func (c ResCode) Msg() string {
	msg, ok := codeMap[c]
	if !ok {
		msg = codeMap[CodeServerBusy]
	}
	return msg
}

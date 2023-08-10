package hsc

// 提供错误信息
var MsgFlags = map[int]string{
	SUCCESS:        "成功",
	ERROR:          "服务器出现错误",
	INVALID_PARAMS: "请求参数异常",
	NOT_PROMISE:    "鉴权失败",

	NOT_FOUND:                      "没有找到请求",
	METHOD_FAILS:                   "Method 请求错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token 检查异常",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token 超时了",
	ERROR_AUTH_TOKEN:               "Token 创建异常",
	ERROR_AUTH:                     "认证失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

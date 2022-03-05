package model

type ApiJson struct {
	Status bool        `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

func ApiResponse(status bool, objects interface{}, msg string) *ApiJson {
	return &ApiJson{Status: status, Data: objects, Msg: msg}
}

// ErrorInsertDatabase 插入数据库失败
func ErrorInsertDatabase(err error) *ApiJson {
	return ApiResponse(false, err.Error(), "插入数据库失败")
}

// ErrorQueryDatabase 查询数据库失败
func ErrorQueryDatabase(err error) *ApiJson {
	return ApiResponse(false, err.Error(), "查询数据库失败")
}

// ErrorUpdateDatabase 更新数据库失败
func ErrorUpdateDatabase(err error) *ApiJson {
	return ApiResponse(false, err.Error(), "更新数据库失败")
}

// ErrorDeleteDatabase 更新数据库失败
func ErrorDeleteDatabase(err error) *ApiJson {
	return ApiResponse(false, err.Error(), "删除数据库失败")
}

// ErrorIncompleteData 数据不完整
func ErrorIncompleteData(err error) *ApiJson {
	return ApiResponse(false, err.Error(), "数据不完整")
}

// ErrorVerification 数据检验失败
func ErrorVerification(err error) *ApiJson {
	return ApiResponse(false, err.Error(), "数据检验失败")
}

// ErrorBuildJWT 生成凭证错误
func ErrorBuildJWT(err error) *ApiJson {
	return ApiResponse(false, err.Error(), "生成凭证错误")
}

// ErrorUnauthorized 未认证登录
func ErrorUnauthorized(err error) *ApiJson {
	return ApiResponse(false, err.Error(), "未认证登录")
}

// ErrorWebSocket 无法升级为长连接
func ErrorWebSocketUpgrade(err error) *ApiJson {
	return ApiResponse(false, err.Error(), "无法升级为长连接")
}

// ErrorWebSocket 服务器内部错误
func ErrorInternalServer(err error) *ApiJson {
	return ApiResponse(false, err.Error(), "服务器内部错误")
}

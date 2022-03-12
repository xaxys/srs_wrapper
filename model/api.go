package model

import (
	"strings"
)

type ApiJson struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func combineError(errs []error) string {
	sb := &strings.Builder{}
	for _, err := range errs {
		sb.WriteString(err.Error())
		sb.WriteString("\n")
	}
	return sb.String()
}

func ApiResponse(code int, status bool, objects interface{}, msg string) *ApiJson {
	return &ApiJson{Code: code, Status: status, Data: objects, Msg: msg}
}

// Success 成功
func Success(objects interface{}, msg string) *ApiJson {
	return ApiResponse(200, true, objects, msg)
}

// SuccessUpdate 成功更新/删除资源
func SuccessUpdate(objects interface{}, msg string) *ApiJson {
	return ApiResponse(204, true, objects, msg)
}

// SuccessCreate 成功创建资源
func SuccessCreate(objects interface{}, msg string) *ApiJson {
	return ApiResponse(201, true, objects, msg)
}

// ErrorInsertDatabase 插入数据库失败
func ErrorInsertDatabase(errs ...error) *ApiJson {
	return ApiResponse(500, false, combineError(errs), "插入数据库失败")
}

// ErrorQueryDatabase 查询数据库失败
func ErrorQueryDatabase(errs ...error) *ApiJson {
	return ApiResponse(500, false, combineError(errs), "查询数据库失败")
}

// ErrorUpdateDatabase 更新数据库失败
func ErrorUpdateDatabase(errs ...error) *ApiJson {
	return ApiResponse(500, false, combineError(errs), "更新数据库失败")
}

// ErrorDeleteDatabase 删除数据库失败
func ErrorDeleteDatabase(errs ...error) *ApiJson {
	return ApiResponse(500, false, combineError(errs), "删除数据库失败")
}

// ErrorNotFound 未找到数据记录
func ErrorNotFound(errs ...error) *ApiJson {
	return ApiResponse(500, false, combineError(errs), "未找到数据记录")
}

// ErrorInvalidData 数据解析失败
func ErrorInvalidData(errs ...error) *ApiJson {
	return ApiResponse(400, false, combineError(errs), "数据解析失败")
}

// ErrorIncompleteData 数据不完整
func ErrorIncompleteData(errs ...error) *ApiJson {
	return ApiResponse(422, false, combineError(errs), "数据不完整")
}

// ErrorVerification 数据检验失败
func ErrorVerification(errs ...error) *ApiJson {
	return ApiResponse(422, false, combineError(errs), "数据检验失败")
}

// ErrorBuildJWT 生成凭证错误
func ErrorBuildJWT(errs ...error) *ApiJson {
	return ApiResponse(500, false, combineError(errs), "生成凭证错误")
}

// ErrorUnauthorized 未认证登录
func ErrorUnauthorized(errs ...error) *ApiJson {
	return ApiResponse(401, false, combineError(errs), "未认证登录")
}

// ErrorNoPermissions 账号权限不足
func ErrorNoPermissions(errs ...error) *ApiJson {
	return ApiResponse(403, false, combineError(errs), "账号权限不足")
}

// ErrorInternalServer 服务器内部错误
func ErrorInternalServer(errs ...error) *ApiJson {
	return ApiResponse(500, false, combineError(errs), "服务器内部错误")
}

package api

// bluebell/api/docs_models.go

// _ResponseLogin 用户登录返回
type _ResponseLogin struct {
	Code    int    `json:"code"`    // 业务响应状态码
	Message string `json:"message"` // 提示信息
	Data    string `json:"data"`    // 数据
}

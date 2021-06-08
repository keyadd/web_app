package request

//注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required,alphanumunicode"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

//登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required,alphanumunicode"`
	Password string `json:"password" binding:"required"`
}

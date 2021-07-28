package request

//注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required,alphanumunicode"`     //用户名
	Password   string `json:"password" binding:"required"`                     //用户名密码
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` //确认用户密码
}

//登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required,alphanumunicode"` //用户名
	Password string `json:"password" binding:"required"`                 //用户密码
}

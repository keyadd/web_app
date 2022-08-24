package request

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required,alphanumunicode"`     //用户名
	Password   string `json:"password" binding:"required"`                     //用户名密码
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` //确认用户密码
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required,alphanumunicode"` //用户名
	Password string `json:"password" binding:"required,timing"`          //用户密码
}

//var validate *validator.Validate

var LoginValidate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(string)
	if ok {
		pattern := `^\d+$|^\d+[.]\d+$` //反斜杠要转义
		result, _ := regexp.MatchString(pattern, date)
		return result
	}
	return false
}

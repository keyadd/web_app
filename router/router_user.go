package router

import (
	"github.com/gin-gonic/gin"
	"web_app/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("")
	{
		UserRouter.POST("/signup", api.SignUpHandler) //注册接口
		UserRouter.POST("/login", api.LoginHandler)   //登录接口
	}
}

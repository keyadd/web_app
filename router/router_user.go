package router

import (
	"github.com/gin-gonic/gin"
	"web_app/api/v1"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("")
	{
		UserRouter.POST("/signup", v1.SignUpHandler) //注册接口
		UserRouter.POST("/login", v1.LoginHandler)   //登录接口

		UserRouter.POST("/image", v1.PutHeaderImage) //图片上传接口
	}
}

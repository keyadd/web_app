package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/global"
	"web_app/middleware"

	_ "web_app/docs" // 千万不要忘了导入把你上一步生成的docs

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(middleware.CorsMiddleware()).Use(middleware.GinLogger()).Use(middleware.GinRecovery(true)).Use(middleware.Sentinel())
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	global.GVA_LOG.Info("register swagger handler")
	v1 := r.Group("/api/v1")
	{

		//无token 访问
		PublicGroup := v1.Group("")
		{
			InitUserRouter(PublicGroup)      // 注册用户路由
			InitCommunityRouter(PublicGroup) // 社区接口

		}
		//token 访问
		PrivateGroup := v1.Group("")
		PrivateGroup.Use(middleware.JWTAuthMiddleware())
		{
			InitPostRouter(PrivateGroup)
			InitCommentRouter(PrivateGroup)

		}
	}
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")

	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "not found",
			"status": "404",
		})

	})
	return r

}

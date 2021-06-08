package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"web_app/controller"
	"web_app/global"
	"web_app/middleware"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	v1 := r.Group("/api/v1")

	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts/", controller.GetPostListHandler)
		//帖子投票
		v1.POST("/vote", controller.PostVoteController)

		//用户对帖子回复

	}

	// r.POST("/usernames")

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")

	})

	r.GET("/ping", func(c *gin.Context) {
		err := global.GVA_REDIS.Set("key", "value23432432", 100000*time.Second).Err()
		if err != nil {
			fmt.Println(err)
		}

		result, err := global.GVA_REDIS.Get("key").Result()

		//fmt.Println(cmd1)

		c.JSON(http.StatusOK, gin.H{
			"key": result,
		})

	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "not found",
			"status": "404",
		})

	})
	return r

}

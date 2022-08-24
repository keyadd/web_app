package router

import (
	"github.com/gin-gonic/gin"
	"web_app/api"
)

func InitPostRouter(Router *gin.RouterGroup) {
	PostRouter := Router.Group("")
	{
		//社区相关接口
		PostRouter.POST("/post", api.CreatePostHandler)       //创建帖子接口
		PostRouter.GET("/post/:id", api.GetPostDetailHandler) //查看帖子详情
		PostRouter.GET("/posts2/", api.GetPostListHandler2)   //帖子列表
		PostRouter.GET("/posts/", api.GetPostListHandler)     // 帖子列表
		PostRouter.POST("/vote", api.PostVoteController)      //帖子投票接口
	}
}

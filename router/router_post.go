package router

import (
	"github.com/gin-gonic/gin"
	"web_app/api/v1"
)

func InitPostRouter(Router *gin.RouterGroup) {
	PostRouter := Router.Group("")
	{
		//社区相关接口
		PostRouter.POST("/post", v1.CreatePostHandler)       //创建帖子接口
		PostRouter.GET("/post/:id", v1.GetPostDetailHandler) //查看帖子详情
		PostRouter.GET("/posts2/", v1.GetPostListHandler2)   //帖子列表
		PostRouter.GET("/posts/", v1.GetPostListHandler)     // 帖子列表
		PostRouter.POST("/vote", v1.PostVoteController)      //帖子投票接口
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"web_app/api/v1"
)

func InitCommunityRouter(Router *gin.RouterGroup) {
	CommunityRouter := Router.Group("")
	{
		CommunityRouter.GET("/community", v1.CommunityHandler)           //查询所有社区接口
		CommunityRouter.GET("/community/:id", v1.CommunityDetailHandler) //社区分类详情
		CommunityRouter.POST("/community_post", v1.GetCommunityPostList) //根据分类id 查询帖子
	}
}

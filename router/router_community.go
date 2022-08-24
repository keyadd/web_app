package router

import (
	"github.com/gin-gonic/gin"
	"web_app/api"
)

func InitCommunityRouter(Router *gin.RouterGroup) {
	CommunityRouter := Router.Group("")
	{
		CommunityRouter.GET("/community", api.CommunityHandler)           //查询所有社区接口
		CommunityRouter.GET("/community/:id", api.CommunityDetailHandler) //社区分类详情
		CommunityRouter.POST("/community_post", api.GetCommunityPostList) //根据分类id 查询帖子
	}
}

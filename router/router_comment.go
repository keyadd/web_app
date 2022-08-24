package router

import (
	"github.com/gin-gonic/gin"
	"web_app/api"
)

func InitCommentRouter(Router *gin.RouterGroup) {
	CommentRouter := Router.Group("")
	{
		//用户对帖子回复
		CommentRouter.POST("/comment/post_comment", api.PostCommentHandle) //发表评论接口
		CommentRouter.POST("/comment/post_reply", api.PostReplyHandle)     //发表回复接口
		CommentRouter.POST("/comment/list", api.CommentListHandler)        //获取评论列表接口
		CommentRouter.POST("/comment/reply_list", api.ReplyListHandler)    //获取回复列表接口
		CommentRouter.POST("/comment/like", api.LikeHandler)               //回复点赞接口
	}
}

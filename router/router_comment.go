package router

import (
	"github.com/gin-gonic/gin"
	"web_app/api/v1"
)

func InitCommentRouter(Router *gin.RouterGroup) {
	CommentRouter := Router.Group("")
	{
		//用户对帖子回复
		CommentRouter.POST("/comment/post_comment", v1.PostCommentHandle) //发表评论接口
		CommentRouter.POST("/comment/post_reply", v1.PostReplyHandle)     //发表回复接口
		CommentRouter.POST("/comment/list", v1.CommentListHandler)        //获取评论列表接口
		CommentRouter.POST("/comment/reply_list", v1.ReplyListHandler)    //获取回复列表接口
		CommentRouter.POST("/comment/like", v1.LikeHandler)               //回复点赞接口
	}
}

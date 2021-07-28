package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controller"
	"web_app/middleware"

	_ "web_app/docs" // 千万不要忘了导入把你上一步生成的docs

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")

	v1.POST("/signup", controller.SignUpHandler) //注册接口
	v1.POST("/login", controller.LoginHandler)   //登录接口

	v1.Use(middleware.JWTAuthMiddleware())
	{
		//社区相关接口
		v1.GET("/community", controller.CommunityHandler)           //查询所有社区接口
		v1.GET("/community/:id", controller.CommunityDetailHandler) //社区分类详情
		//帖子相关接口
		v1.POST("/post", controller.CreatePostHandler)       //创建帖子接口
		v1.GET("/post/:id", controller.GetPostDetailHandler) //查看帖子详情
		v1.GET("/posts2/", controller.GetPostListHandler2)   //
		v1.GET("/posts/", controller.GetPostListHandler)
		v1.POST("/vote", controller.PostVoteController) //帖子投票接口
		//用户对帖子回复
		v1.POST("/comment/post_comment", controller.PostCommentHandle) //发表评论接口
		v1.POST("/comment/post_reply", controller.PostReplyHandle)     //发表回复接口
		v1.POST("/comment/list", controller.CommentListHandler)        //获取评论列表接口
		v1.POST("/comment/reply_list", controller.ReplyListHandler)    //获取回复列表接口
		v1.POST("/comment/like", controller.LikeHandler)               //回复点赞接口

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

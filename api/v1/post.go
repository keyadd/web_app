package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/global"
	"web_app/model/request"
	"web_app/service"
	"web_app/utils"
)

// CreatePostHandler CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	// 获取参数及参数的校验
	p := new(request.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param ", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	userID, err := utils.GetUserId(c)
	if err != nil {
		utils.ResponseError(c, global.CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 创建帖子
	if err := service.CreatePost(p); err != nil {
		zap.L().Error("service.CreatePost(p) failed ", zap.Error(err))
		utils.ResponseError(c, global.CodeServerBusy)
		return
	}

	//返回响应
	utils.ResponseSuccess(c, nil)

}

//查看当前帖子
func GetPostDetailHandler(c *gin.Context) {
	//从URL获取帖子id
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	//根据id取出帖子的数据
	data, err := service.GetPostById(pid)
	if err != nil {
		zap.L().Error("service.GetPostById(pid) failed", zap.Error(err))
		utils.ResponseError(c, global.CodeServerBusy)
		return
	}

	//返回响应
	utils.ResponseSuccess(c, data)
}

//GetPostListHandler 获取帖子列表的处理方法
func GetPostListHandler(c *gin.Context) {

	//分页参数
	page, size := utils.GetPageInfo(c)
	//获取数据
	data, err := service.GetPostList(page, size)
	if err != nil {
		global.GVA_LOG.Error("service.GetPostList() failed", zap.Error(err))
		utils.ResponseError(c, global.CodeServerBusy)
		return
	}

	//返回响应
	utils.ResponseSuccess(c, data)

}

//GetPostListHandler 获取帖子列表的处理方法
func GetPostListHandler2(c *gin.Context) {

	//分页参数
	p := &request.PostList{
		request.PageInfo{
			Page:     1,
			PageSize: 10,
		},
		request.OrderTime,
	} //获取数据
	if err := c.ShouldBind(p); err != nil {
		global.GVA_LOG.Error("GetPostListHandler2 with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	fmt.Println(p)
	data, err := service.GetPostList2(p)
	if err != nil {
		global.GVA_LOG.Error("service.GetPostList() failed", zap.Error(err))
		utils.ResponseError(c, global.CodeServerBusy)
		return
	}

	//返回响应
	utils.ResponseSuccess(c, data)

}

// GetCommunityPostList 根据社区去查询
func GetCommunityPostList(c *gin.Context) {

	//分页参数
	p := &request.CommunityPostList{
		PostList: request.PostList{
			PageInfo: request.PageInfo{
				Page:     1,
				PageSize: 10,
			},
			Order: request.OrderTime,
		},
		CommunityId: 1,
	} //获取数据

	if err := c.ShouldBind(p); err != nil {
		global.GVA_LOG.Error("GetPostListHandler2 with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	fmt.Println(p)
	data, err := service.GetCommunityPostList(p)
	if err != nil {
		global.GVA_LOG.Error("service.GetPostList() failed", zap.Error(err))
		utils.ResponseError(c, global.CodeServerBusy)
		return
	}

	//返回响应
	utils.ResponseSuccess(c, data)
}

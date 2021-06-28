package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/global"
	"web_app/model"
	"web_app/model/request"
	"web_app/service"
	"web_app/utils"
)

//发表评论接口
func PostCommentHandle(c *gin.Context) {
	var p request.PostComment
	if err := c.BindJSON(&p); err != nil {
		global.GVA_LOG.Error("PostCommentHandle with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	userID, err := utils.GetUserId(c)
	if err != nil {
		utils.ResponseError(c, global.CodeNeedLogin)
		return
	}
	p.AuthorId = userID

	if err := service.GetPostComment(&p); err != nil {
		global.GVA_LOG.Error("service.GetPostComment(&p) error", zap.Error(err))
		if errors.Is(err, model.ErrorBeginx) {
			utils.ResponseError(c, global.CodeBeginxError)
			return
		}
		utils.ResponseError(c, global.CodeServerBusy)
		return
	}
	//3.返回响应
	utils.ResponseSuccess(c, nil)

}

//发表回复接口
func PostReplyHandle(c *gin.Context) {
	var p request.PostReply
	if err := c.BindJSON(&p); err != nil {
		global.GVA_LOG.Error("PostCommentHandle with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	userID, err := utils.GetUserId(c)
	if err != nil {
		utils.ResponseError(c, global.CodeNeedLogin)
		return
	}
	p.AuthorId = userID

	if err := service.GetPostReply(&p); err != nil {
		global.GVA_LOG.Error("service.GetPostReply error", zap.Error(err))
		if errors.Is(err, model.ErrorBeginx) {
			utils.ResponseError(c, global.CodeBeginxError)
			return
		}
		utils.ResponseError(c, global.CodeServerBusy)
		return
	}
	//3.返回响应
	utils.ResponseSuccess(c, nil)

}

//评论列表
func CommentListHandler(c *gin.Context) {

	//var p request.GetByPostId
	p := &request.GetByPostId{
		request.PageInfo{
			Page:     1,
			PageSize: 10,
		},
		0,
	}

	if err := c.BindJSON(&p); err != nil {
		global.GVA_LOG.Error("CommentListHandler with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	data, err := service.GetCommentList(p)
	if err != nil {
		global.GVA_LOG.Error("service.GetCommentList(p) failed", zap.Error(err))
		utils.ResponseError(c, global.CodeServerBusy)
		return
	}

	//返回响应
	utils.ResponseSuccess(c, data)
}

//回复评论列表
func ReplyListHandler(c *gin.Context) {

	p := &request.GetByCommentId{
		request.PageInfo{
			Page:     1,
			PageSize: 10,
		},
		0,
	}

	if err := c.BindJSON(&p); err != nil {
		global.GVA_LOG.Error("CommentListHandler with invalid params", zap.Error(err))
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	data, err := service.GetReplyList(p)
	if err != nil {
		global.GVA_LOG.Error("service.GetCommentList(p) failed", zap.Error(err))
		utils.ResponseError(c, global.CodeServerBusy)
		return
	}

	//返回响应
	utils.ResponseSuccess(c, data)

}

func LikeHandler(c *gin.Context) {}

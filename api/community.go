package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/global"
	"web_app/service"
	"web_app/utils"
)

func CommunityHandler(c *gin.Context) {
	//查询所以社区 以列表形式返回
	data, err := service.GetCommunityList()
	if err != nil {
		global.GVA_LOG.Error("service.GetCommunityList() failed", zap.Error(err))
		utils.ResponseError(c, global.CodeServerBusy)
	}
	utils.ResponseSuccess(c, data)
}

//CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	//获取社区ID
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseError(c, global.CodeInvalidParam)
		return
	}

	//查询所以社区 以列表形式返回
	data, err := service.GetCommunityListDetail(id)
	if err != nil {
		global.GVA_LOG.Error("service.GetCommunityListDetail() failed", zap.Error(err))
		utils.ResponseError(c, global.CodeServerBusy)
	}
	utils.ResponseSuccess(c, data)
}

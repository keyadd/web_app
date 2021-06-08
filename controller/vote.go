package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/global"
	"web_app/model/request"
	"web_app/service"
	"web_app/utils"
)

func PostVoteController(c *gin.Context) {
	p := new(request.VoteData)
	fmt.Println(p)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			utils.ResponseError(c, global.CodeInvalidParam)
			return
		}
		utils.ResponseErrorWithMsg(c, global.CodeInvalidParam, errs.Translate(global.TRANS))
		return

	}
	fmt.Println(p)

	userID, err := utils.GetUserId(c)
	if err != nil {
		utils.ResponseError(c, global.CodeNeedLogin)
		return
	}
	if err := service.VoteForPost(userID, p); err != nil {
		global.GVA_LOG.Error("service.VoteForPost(userID, p) failed", zap.Error(err))
		utils.ResponseError(c, global.CodeServerBusy)
	}
	utils.ResponseSuccess(c, nil)

}

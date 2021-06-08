package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/global"
	"web_app/model"
	"web_app/model/request"
	"web_app/service"
	"web_app/utils"
)

func SignUpHandler(c *gin.Context) {
	//1.参数校验
	var p = new(request.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			utils.ResponseError(c, global.CodeInvalidParam)
			return

		}
		utils.ResponseErrorWithMsg(c, global.CodeInvalidParam, errs.Translate(global.TRANS))
		return

	}

	//2.业务处理
	if err := service.SignUp(p); err != nil {
		if errors.Is(err, model.ErrorUserExist) {
			utils.ResponseError(c, global.CodeUserExist)
			return
		}
		utils.ResponseError(c, global.CodeServerBusy)
		return
	}

	//3.返回响应
	utils.ResponseSuccess(c, nil)

}

func LoginHandler(c *gin.Context) {
	//获取请求参数及参数校验
	p := new(request.ParamLogin)

	if err := c.ShouldBindJSON(p); err != nil {
		global.GVA_LOG.Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			utils.ResponseError(c, global.CodeInvalidParam)
			return
		}
		utils.ResponseErrorWithMsg(c, global.CodeInvalidParam, errs.Translate(global.TRANS))
		return
	}

	//业务逻辑处理
	token, err := service.Login(p)
	if nil != err {
		global.GVA_LOG.Error("service.Login(p) failed", zap.Error(err))

		if errors.Is(err, model.ErrorUserNotExist) {
			utils.ResponseError(c, global.CodeUserNotExist)
		}
		utils.ResponseError(c, global.CodeInvalidPassword)
		return
	}
	//返回响应
	utils.ResponseSuccess(c, token)

}

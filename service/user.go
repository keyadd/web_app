package service

import (
	"fmt"
	"web_app/core"
	"web_app/model"
	"web_app/model/request"
	"web_app/model/response"
)

func SignUp(p *request.ParamSignUp) (err error) {

	//判断用户不存在
	err = model.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	fmt.Println(p.Username)
	//生成Uid

	userID := core.GenID()

	u := &response.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//保存进数据库
	return model.InsertUser(u)

}

func Login(p *request.ParamLogin) (token string, err error) {
	user := &response.User{
		Username: p.Username,
		Password: p.Password,
	}

	if err := model.Login(user); err != nil {
		return "", err
	}
	//生成token
	return core.GenToken(user.UserID, user.Username)

}

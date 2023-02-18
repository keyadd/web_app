package model

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"web_app/global"
	"web_app/model/response"
	"web_app/utils"
)

var (
	ErrorUserExist        = errors.New("用户已存在")
	ErrorUserNotExist     = errors.New("not is username")
	ErrorPasswordNotExist = errors.New("not is password")
)

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := global.SQLX_DB.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return

}

func InsertUser(user *response.User) (err error) {
	user.Password = utils.MD5V([]byte(user.Password))
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = global.SQLX_DB.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return

}

func Login(user *response.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id,username,password from user where username=?`
	err = global.SQLX_DB.Get(user, sqlStr, user.Username)

	//用户不存在
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	//判断密码是否正确
	password := utils.MD5V([]byte(oPassword))
	if password != user.Password {
		return ErrorPasswordNotExist
	}
	return
}

// GetUserById 根据id 获取用户信息
func GetUserById(uid int64) (user *response.User, err error) {
	user = new(response.User)
	sqlStr := `select user_id,username from user where user_id=?`
	err = global.SQLX_DB.Get(user, sqlStr, uid)
	return

}

// GetUserInfoList 批量查询用户信息
func GetUserInfoList(userIdList []int64) (userInfoList []*response.UserInfo, err error) {

	if len(userIdList) == 0 {
		return
	}
	sqlstr := `select user_id, gender, username, email from  user where user_id in(?)`
	var userIdTmpArr []interface{}
	for _, userId := range userIdList {
		userIdTmpArr = append(userIdTmpArr, userId)
	}
	query, args, err := sqlx.In(sqlstr, userIdTmpArr)
	if err != nil {
		return nil, err
	}
	err = global.SQLX_DB.Select(&userInfoList, query, args...)
	if err != nil {
		return nil, err
	}
	return
}

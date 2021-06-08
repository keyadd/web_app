package service

import (
	"web_app/model"
	"web_app/model/response"
)

func GetCommunityList() ([]*response.Community, error) {

	//查数据库 查找所有的community 返回
	return model.GetCommunityList()

}

func GetCommunityListDetail(id int64) (*response.CommunityDetail, error) {
	return model.GetCommunityDetailByID(id)
}

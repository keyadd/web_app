package model

import (
	"database/sql"
	"web_app/global"
	"web_app/model/response"
)

func GetCommunityList() (communityList []*response.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err := global.SQLX_DB.Select(&communityList, sqlStr); nil != err {
		if err == sql.ErrNoRows {
			global.GVA_LOG.Warn("there is no community in db")
			err = nil
		}
	}
	return

}

func GetCommunityDetailByID(id int64) (*response.CommunityDetail, error) {
	community := new(response.CommunityDetail)

	sqlStr := `select community_id,community_name,introduction from community where community_id = ?`
	err := global.SQLX_DB.Get(community, sqlStr, id)
	if err == sql.ErrNoRows {
		err = global.ErrorInvalidID
	}
	return community, err

}

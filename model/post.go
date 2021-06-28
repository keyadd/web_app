package model

import (
	"github.com/jmoiron/sqlx"
	"strings"
	"web_app/global"
	"web_app/model/request"
)

func CreatePost(p *request.Post) (err error) {
	sqlStr := `insert into post (post_id,title,content,author_id,community_id) values(?,?,?,?,?)`
	_, err = global.SQLX_DB.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return

}

//查询单篇帖子的数据
func GetPostById(pid int64) (data *request.Post, err error) {
	post := new(request.Post)
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post where post_id =?`
	err = global.SQLX_DB.Get(post, sqlStr, pid)
	return post, err

}

func GetPostList(page int64, size int64) (posts []*request.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post limit ?,?`

	posts = make([]*request.Post, 0, 2)
	err = global.SQLX_DB.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

func GetPostListById(result []string) (postList []*request.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post where post_id in (?) order by FIND_IN_SET(post_id,?)`
	query, args, err := sqlx.In(sqlStr, result, strings.Join(result, ","))
	if err != nil {
		return nil, err
	}
	query = global.SQLX_DB.Rebind(query)
	err = global.SQLX_DB.Select(&postList, query, args...)
	return

}

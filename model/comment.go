package model

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"web_app/global"
	"web_app/model/request"
	"web_app/model/response"
)

var (
	ErrorBeginx = errors.New("InsertPostComment() beginx is error")
	ErrorQuery  = errors.New("query comment list failed")
)

//插入评论

func InsertPostComment(p *response.PostComment) (err error) {

	beginx, err := global.SQLX_DB.Beginx()
	if err != nil {
		return ErrorBeginx
	}
	sqlStr := `insert into comment(comment_id,content,author_id,like_count,comment_count,create_time) values(?,?,?,?,?,?)`
	_, err = beginx.Exec(sqlStr, p.CommentId, p.Content, p.AuthorId, p.LikeCount, p.CommentCount, p.CreateTime)
	if err != nil {
		fmt.Println(err)
		beginx.Rollback()
		return err
	}
	sqlStr = `insert into comment_rel(comment_id,parent_id,level,post_id,reply_author_id) values(?,?,?,?,?)`
	_, err = beginx.Exec(sqlStr, p.PostCommentRel.CommentId, p.PostCommentRel.ParentId, p.PostCommentRel.Level, p.PostCommentRel.PostId, p.PostCommentRel.ReplyAuthorId)
	if err != nil {
		fmt.Println(err)
		beginx.Rollback()
		return err
	}
	sqlStr = `update comment set comment_count = comment_count+1 where comment_id = ?`
	_, err = beginx.Exec(sqlStr, p.CommentId)
	if err != nil {
		fmt.Println(err)
		beginx.Rollback()
		return err
	}

	err = beginx.Commit()
	if err != nil {
		return err
	}
	return

}

//插入回复

func InsertPostReply(p *response.PostComment) (err error) {

	beginx, err := global.SQLX_DB.Beginx()
	if err != nil {
		return ErrorBeginx
	}

	sqlStr := `insert into comment(comment_id,content,author_id,like_count,comment_count,create_time) values(?,?,?,?,?,?)`
	_, err = beginx.Exec(sqlStr, p.CommentId, p.Content, p.AuthorId, p.LikeCount, p.CommentCount, p.CreateTime)
	if err != nil {
		fmt.Println(err)
		beginx.Rollback()
		return err
	}
	sqlStr = `insert into comment_rel(comment_id,parent_id,level,post_id,reply_author_id,reply_comment_id) values(?,?,?,?,?,?)`
	_, err = beginx.Exec(sqlStr, p.PostCommentRel.CommentId, p.PostCommentRel.ParentId, p.PostCommentRel.Level, p.PostCommentRel.PostId, p.PostCommentRel.ReplyAuthorId, p.ReplyCommentId)
	if err != nil {
		fmt.Println(err)
		beginx.Rollback()
		return err
	}
	sqlStr = `update comment set comment_count = comment_count+1 where comment_id = ?`
	_, err = beginx.Exec(sqlStr, p.CommentId)
	if err != nil {
		fmt.Println(err)
		beginx.Rollback()
		return err
	}

	err = beginx.Commit()
	if err != nil {
		return err
	}
	return

}

// SelectPostComment 查询评论的作者的author_id
func SelectPostComment(p *response.PostCommentRel) (AuthorId int64, err error) {
	var replyAuthorId int64
	sqlStr := `select author_id from comment where comment_id = ?`
	err = global.SQLX_DB.Get(&replyAuthorId, sqlStr, p.ReplyCommentId)
	fmt.Println(err)
	return replyAuthorId, err
}

// SelectCommentList 查询评论第一层列表
func SelectCommentList(p *request.GetByPostId) (apicommentList *response.ApiCommentList, err error) {
	var commentIdList []int64

	sqlStr := `select comment_id from comment_rel where post_id = ? and level = 1 limit ?,?`

	start := (p.PageInfo.Page - 1) * p.PageInfo.PageSize
	end := start + p.PageInfo.PageSize - 1

	err = global.SQLX_DB.Select(&commentIdList, sqlStr, p.PostId, start, end)
	if err != nil {
		return nil, err
	}
	sqlStr = `select comment_id,content,author_id,like_count,comment_count,create_time from comment where comment_id in (?)`
	sqlstr, paramList, err := sqlx.In(sqlStr, commentIdList)
	if err != nil {
		return nil, err
	}
	var list = &response.ApiCommentList{}
	err = global.SQLX_DB.Select(&list.CommentList, sqlstr, paramList...)
	if err != nil {
		return nil, err
	}
	//查询总的记录条数
	sqlStr = `select count(comment_id) from comment_rel where post_id = ? and level = 1`

	err = global.SQLX_DB.Get(&list.Count, sqlStr, p.PostId)
	if err != nil {
		return nil, err
	}

	return list, err

}

func SelectReplyList(p *request.GetByCommentId) (apicommentList *response.ApiCommentList, err error) {
	var commentIdList []int64

	sqlStr := `select comment_id from comment_rel where parent_id = ? and level = 2 limit ?,?`

	start := (p.PageInfo.Page - 1) * p.PageInfo.PageSize
	end := start + p.PageInfo.PageSize - 1

	err = global.SQLX_DB.Select(&commentIdList, sqlStr, p.CommentId, start, end)
	if err != nil {
		return nil, err
	}
	if len(commentIdList) == 0 {
		return nil, ErrorQuery
	}
	sqlStr = `select c.comment_id,c.content,c.author_id,c.like_count,c.comment_count,c.create_time,r.reply_author_id from comment as c, comment_rel as r where c.comment_id=r.comment_id and c.comment_id in (?)`
	sqlstr, paramList, err := sqlx.In(sqlStr, commentIdList)
	if err != nil {
		return nil, err
	}
	commentList := []*response.CommentList{}
	var list = &response.ApiCommentList{}

	err = global.SQLX_DB.Select(&commentList, sqlstr, paramList...)
	if err != nil {
		return nil, err
	}
	list.CommentList = commentList
	//查询总的记录条数
	sqlStr = `select count(comment_id) from comment_rel where parent_id = ? and level = 2`

	err = global.SQLX_DB.Get(&list.Count, sqlStr, p.CommentId)
	if err != nil {
		return nil, err
	}

	return list, err

}

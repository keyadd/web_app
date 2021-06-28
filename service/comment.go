package service

import (
	"fmt"
	"go.uber.org/zap"
	"time"
	"web_app/core"
	"web_app/global"
	"web_app/model"
	"web_app/model/request"
	"web_app/model/response"
)

func GetPostComment(p *request.PostComment) (err error) {

	GenID := core.GenID()

	CommentRel := response.PostCommentRel{
		CommentId:     GenID,
		ParentId:      p.ParentId,
		Level:         1,
		PostId:        p.PostId,
		ReplyAuthorId: p.AuthorId,
	}

	comment := &response.PostComment{
		CommentRel,
		GenID,
		p.Content,
		p.AuthorId,
		0,
		0,
		time.Now(),
	}
	err = model.InsertPostComment(comment)
	if err != nil {
		return
	}
	return

}

func GetPostReply(p *request.PostReply) (err error) {
	GenID := core.GenID()

	CommentRel := &response.PostCommentRel{
		CommentId:      GenID,
		ParentId:       p.ParentId,
		Level:          2,
		PostId:         p.PostId,
		ReplyAuthorId:  p.AuthorId,
		ReplyCommentId: p.ReplyCommentId,
	}
	fmt.Println(CommentRel)
	CommentRel.ReplyAuthorId, err = model.SelectPostComment(CommentRel)
	fmt.Println(err)
	if err != nil {
		return err
	}

	comment := &response.PostComment{
		*CommentRel,
		GenID,
		p.Content,
		p.AuthorId,
		0,
		0,
		time.Now(),
	}

	err = model.InsertPostReply(comment)
	if err != nil {
		return
	}
	return

}

func GetCommentList(p *request.GetByPostId) (commentList *response.ApiCommentList, err error) {

	list, err := model.SelectCommentList(p)
	if err != nil {
		global.GVA_LOG.Error("model.SelectCommentList(p) list", zap.Int64("PostId", p.PostId), zap.Error(err))
		return
	}

	var userIdList []int64
	for _, v := range list.CommentList {
		userIdList = append(userIdList, v.AuthorId)
	}
	infoList, err := model.GetUserInfoList(userIdList)
	if err != nil {
		global.GVA_LOG.Error("model.GetUserInfoList(userIdList) list error", zap.Error(err))
		return
	}

	userInfoMap := make(map[int64]*response.UserInfo, len(infoList))
	for _, user := range infoList {
		userInfoMap[user.UserID] = user
	}
	for _, v := range *&list.CommentList {

		user, ok := userInfoMap[v.AuthorId]
		if ok {
			v.AuthorName = user.Username
		}
		user, ok = userInfoMap[v.ReplyAuthorId]
		if ok {
			v.ReplyAuthorName = user.Username
		}
	}

	return list, nil
}

func GetReplyList(p *request.GetByCommentId) (commentList *response.ApiCommentList, err error) {

	list, err := model.SelectReplyList(p)

	if err != nil {
		global.GVA_LOG.Error("model.SelectCommentList(p) list", zap.Int64("CommentId", p.CommentId), zap.Error(err))
		return
	}

	var userIdList []int64
	for _, v := range list.CommentList {
		userIdList = append(userIdList, v.AuthorId)
	}
	//fmt.Println(list.CommentList)
	infoList, err := model.GetUserInfoList(userIdList)
	if err != nil {
		global.GVA_LOG.Error("model.GetUserInfoList(userIdList) list error", zap.Error(err))
		return
	}

	userInfoMap := make(map[int64]*response.UserInfo, len(infoList))
	for _, user := range infoList {
		userInfoMap[user.UserID] = user
	}
	for _, v := range *&list.CommentList {

		user, ok := userInfoMap[v.AuthorId]
		if ok {
			v.AuthorName = user.Username
		}
		user, ok = userInfoMap[v.ReplyAuthorId]
		fmt.Println(v)
		if ok {
			v.ReplyAuthorName = user.Username
		}
	}

	return list, nil
}

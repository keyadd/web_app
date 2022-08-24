package service

import (
	"fmt"
	"go.uber.org/zap"
	"web_app/core"
	"web_app/global"
	"web_app/model"
	"web_app/model/redis"
	"web_app/model/request"
	"web_app/model/response"
)

func CreatePost(p *request.Post) (err error) {
	// 生成Post id
	p.ID = int64(core.GenID())
	// 保存数据库
	err = model.CreatePost(p)
	if err != nil {
		return err
	}

	err = redis.CreatePost(p.ID, p.CommunityID)
	return

	//返回

}

//查询帖子详情数据
func GetPostById(pid int64) (data *response.ApiPostDetail, err error) {
	//查询数据组合数据

	post, err := model.GetPostById(pid)
	if err != nil {
		global.GVA_LOG.Error("model.GetPostById(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}
	user, err := model.GetUserById(post.AuthorID)
	if err != nil {
		global.GVA_LOG.Error("model.GetUserById(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return
	}
	//查询社区详细
	community, err := model.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		global.GVA_LOG.Error("model.GetCommunityDetailByID(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
		return

	}
	data = &response.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return

}

//GetPostList 获取帖子列表
func GetPostList(page int64, size int64) (data []*response.ApiPostDetail, err error) {
	posts, err := model.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*response.ApiPostDetail, 0, len(posts))

	for _, post := range posts {

		user, err := model.GetUserById(post.AuthorID)
		if err != nil {
			global.GVA_LOG.Error("model.GetUserById(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}
		//查询社区详细
		community, err := model.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			global.GVA_LOG.Error("model.GetCommunityDetailByID(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
			continue

		}
		postdetail := &response.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		fmt.Println(postdetail)

		data = append(data, postdetail)
	}
	return
}

//GetPostList 获取帖子列表
func GetPostList2(p *request.PostList) (data []*response.ApiPostDetail, err error) {
	//	redis查询id列表

	result, err := redis.GetPostList(p)
	if err != nil {
		return
	}

	global.GVA_LOG.Debug("GetPostList2 redis.GetPostList(p) result data", zap.Any("Post_Id", result))
	if len(result) == 0 {
		global.GVA_LOG.Warn("GetPostList2 redis.GetPostList(p) return 0 data")
		return
	}

	posts, err := model.GetPostListById(result)
	if err != nil {
		return
	}

	//根据id去数据库查询帖子

	for _, post := range posts {

		user, err := model.GetUserById(post.AuthorID)
		if err != nil {
			global.GVA_LOG.Error("model.GetUserById(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}
		//查询社区详细
		community, err := model.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			global.GVA_LOG.Error("model.GetCommunityDetailByID(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
			continue

		}
		postdetail := &response.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		fmt.Println(postdetail)

		data = append(data, postdetail)
	}
	return
}

func GetCommunityPostList(p *request.CommunityPostList) (data []*response.ApiPostDetail, err error) {
	//	redis查询id列表

	result, err := redis.GetCommunityPostList(p)
	fmt.Println(result)
	if err != nil {
		return
	}

	global.GVA_LOG.Debug("GetPostList2 redis.GetPostList(p) result data", zap.Any("Post_Id", result))
	if len(result) == 0 {
		global.GVA_LOG.Warn("GetPostList2 redis.GetPostList(p) return 0 data")
		return
	}

	posts, err := model.GetPostListById(result)
	if err != nil {
		return
	}

	//根据id去数据库查询帖子

	for _, post := range posts {

		user, err := model.GetUserById(post.AuthorID)
		if err != nil {
			global.GVA_LOG.Error("model.GetUserById(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}
		//查询社区详细
		community, err := model.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			global.GVA_LOG.Error("model.GetCommunityDetailByID(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
			continue

		}
		postdetail := &response.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		fmt.Println(postdetail)

		data = append(data, postdetail)
	}
	return
}

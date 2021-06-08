package service

import (
	"go.uber.org/zap"
	"strconv"
	"web_app/global"
	"web_app/model/redis"
	"web_app/model/request"
)

// 投票功能

//1 用户投票的数据
//2 投一票 就加432分 86400/200 ->200 张赞成票 就增加一天

/*

 */

func VoteForPost(userID int64, p *request.VoteData) error {
	//帖子投票

	global.GVA_LOG.Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postId", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))

}

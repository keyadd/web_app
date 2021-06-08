package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
	"web_app/global"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600 //一周多少秒
	scorePerVote     = 432           //每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

func CreatePost(postId int64) error {
	pipeline := global.GVA_REDIS.TxPipeline()
	//帖子时间
	pipeline.ZAdd(GetRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	//帖子分数
	pipeline.ZAdd(GetRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userId, postId string, value float64) error {

	//判断投票限制
	postTime := global.GVA_REDIS.ZScore(GetRedisKey(KeyPostTimeZSet), postId).Val()
	//判断这个帖子的时间是否大于一周
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	//更新帖子分数
	//查询当前用户 redis 的投票记录

	ov := global.GVA_REDIS.ZScore(GetRedisKey(KeyPostVotedZSetPrefix+postId), userId).Val()

	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	pipeline := global.GVA_REDIS.TxPipeline()
	pipeline.ZIncrBy(GetRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postId)

	//记录用户为改帖子投票的数据
	if value == 0 {
		pipeline.ZRem(GetRedisKey(KeyPostVotedZSetPrefix+postId), postId)
	} else {
		pipeline.ZAdd(GetRedisKey(KeyPostVotedZSetPrefix+postId), redis.Z{
			Score:  value,
			Member: userId,
		})

	}
	_, err := pipeline.Exec()
	return err
}

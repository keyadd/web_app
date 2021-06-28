package redis

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
	"web_app/global"
	"web_app/model/request"
)

func GetPostList(p *request.PostList) ([]string, error) {
	//获取redis id
	key := GetRedisKey(KeyPostTimeZSet)
	if p.Order == request.OrderScore {
		key = GetRedisKey(KeyPostScoreZSet)
	}
	start := (p.Page - 1) * p.PageSize
	end := start + p.PageSize - 1
	//ZRevRange select

	return global.GVA_REDIS.ZRevRange(key, start, end).Result()
}

//按社区查询数据
func GetCommunityPostList(p *request.CommunityPostList) ([]string, error) {

	orderKey := GetRedisKey(KeyPostTimeZSet)
	if p.Order == request.OrderScore {
		orderKey = GetRedisKey(KeyPostScoreZSet)
	}

	cKey := GetRedisKey(keyCommunityZSetPrefix + strconv.Itoa(int(p.CommunityId)))

	key := orderKey + strconv.Itoa(int(p.CommunityId))
	if global.GVA_REDIS.Exists(key).Val() < 1 {
		pipeline := global.GVA_REDIS.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey) //计算redis set 和zset 的交集
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}

	start := (p.Page - 1) * p.PageSize
	end := start + p.PageSize - 1
	//ZRevRange select

	return global.GVA_REDIS.ZRevRange(key, start, end).Result()
}

package redis

const (
	Prefix                 = "web_app:"
	KeyPostTimeZSet        = "post:time"  //zset 帖子 及发帖时间
	KeyPostScoreZSet       = "post:score" //zset 帖子 及投票的分数
	KeyPostVotedZSetPrefix = "post:voted" //zset 记录用户及投票的类型 参数是post_id
)

func GetRedisKey(key string) string {
	return Prefix + key
}

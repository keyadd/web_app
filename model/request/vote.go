package request

//投票
type VoteData struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction" binding:"oneof=1 0 -1"` //赞成票还是反对票 赞成票1 反对票-1 取消 0
}

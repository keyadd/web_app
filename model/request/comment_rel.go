package request

type PostCommentRel struct {
	ParentId int64 `json:"parent_id" db:"parent_id"` //父层ID
	PostId   int64 `json:"post_id" db:"post_id"`     //帖子评论id
}

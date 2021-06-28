package response

type PostCommentRel struct {
	CommentId      int64 `json:"comment_id" db:"comment_id"`             //评论id
	ParentId       int64 `json:"parent_id" db:"parent_id"`               //父层ID
	Level          int64 `json:"level" db:"level"`                       //层次
	PostId         int64 `json:"post_id" db:"post_id"`                   //帖子评论id
	ReplyAuthorId  int64 `json:"reply_author_id" ob:"reply_author_id"`   //回复评论作者的ID
	ReplyCommentId int64 `json:"reply_comment_id" db:"reply_comment_id"` //回复评论 的id
}

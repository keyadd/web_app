package request

type PostComment struct {
	//CommentId int64  `json:"comment_id" db:"comment_id"` //评论id
	Content  string `json:"content" db:"content"`     //评论内容
	AuthorId int64  `json:"author_id" db:"author_id"` //用户id
	//LikeCount    int       `json:"like_count" db:"like_count"`       //点赞数
	//CommentCount int       `json:"comment_count" db:"comment_count"` //评论数
	//CreateTime time.Time `json:"create_time" db:"create_time"` //创建时间
	PostCommentRel
}

type PostReply struct {
	PostComment
	ReplyCommentId int64 `json:"reply_comment_id" db:"reply_comment_id"`
}

type GetByPostId struct {
	PageInfo
	PostId int64 `json:"post_id" db:"post_id"`
}

type GetByCommentId struct {
	PageInfo
	CommentId int64 `json:"comment_id" db:"comment_id"`
}

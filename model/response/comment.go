package response

import (
	"time"
)

type PostComment struct {
	PostCommentRel
	CommentId    int64     `json:"comment_id" db:"comment_id"`       //评论id
	Content      string    `json:"content" db:"content"`             //评论内容
	AuthorId     int64     `json:"author_id" db:"author_id"`         //用户id
	LikeCount    int       `json:"like_count" db:"like_count"`       //点赞数
	CommentCount int       `json:"comment_count" db:"comment_count"` //评论数
	CreateTime   time.Time `json:"create_time" db:"create_time"`     //创建时间

}

//回复帖子
type CommentList struct {
	CommentId       int64     `json:"comment_id" db:"comment_id"`           //评论id
	Content         string    `json:"content" db:"content"`                 //评论内容
	AuthorId        int64     `json:"author_id" db:"author_id"`             //用户id
	LikeCount       int       `json:"like_count" db:"like_count"`           //点赞数
	CommentCount    int       `json:"comment_count" db:"comment_count"`     //评论数
	CreateTime      time.Time `json:"create_time" db:"create_time"`         //创建时间
	PostId          int64     `json:"post_id" db:"post_id"`                 //帖子评论id
	ReplyAuthorId   int64     `json:"reply_author_id" db:"reply_author_id"` //回复评论作者的ID
	AuthorName      string    `json:"author_name" `                         //发表用户名称
	ReplyAuthorName string    `json:"reply_author_name"`                    //回复的对象的名字
}

//评论回复
type ApiCommentList struct {
	CommentList []*CommentList `json:"comment_list"`     //回复帖子
	Count       int64          `json:"count" db:"count"` //回复总数
}

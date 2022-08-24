package request

import "time"

const (
	OrderTime  = "time"  //时间
	OrderScore = "score" //文章是否火热
)

type Post struct {
	ID          int64     `json:"id" db:"post_id"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id"`
	Status      int64     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// PostList 帖子列表参数获取
type PostList struct {
	PageInfo
	Order string `json:"order" form:"order"`
}

type CommunityPostList struct {
	PostList
	CommunityId int64 `json:"community_id" form:"community_id"`
}

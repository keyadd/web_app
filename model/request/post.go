package request

import "time"

type Post struct {
	ID          int64     `json:"id" db:"post_id"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id"`
	Status      int64     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

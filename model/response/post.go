package response

import "web_app/model/request"

// ApiPostDetail 帖子详情结构体
type ApiPostDetail struct {
	AuthorName       string             `json:"author_name"`
	*request.Post                       //帖子的结构体
	*CommunityDetail `json:"community"` //社区信息
}

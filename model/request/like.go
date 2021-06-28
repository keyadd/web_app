package request

const (
	LikeType
)

type Like struct {
	Id       int64 `json:"id"`
	LikeType int `json:"like_type"`
}

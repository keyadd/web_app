package response

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type UserInfo struct {
	UserID   int64  `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Gender   int64  `json:"gender" db:"gender"`
	Email    string `json:"email" db:"email"`
}

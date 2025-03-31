package todo

type User struct {
	Id           int    `json:"-" db:"id"`
	Name         string `json:"name" binding:"required"`
	Username     string `json:"username" binding:"required"`
	PasswordHash string `json:"password_hash" binding:"required"`
}

package structure

type User struct {
	Id       int    `json:"-" db:"id"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
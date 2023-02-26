package structure

type User struct {
	Id       int    `json:"-" db:"id"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UsersVariables struct {
	Id int `json:"id"  db:"id" `
	UserId int `json:"user_id" db:"user_id"`
	Variables string `json:"variables" db:"variables"`
}
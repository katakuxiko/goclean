package structure

type User struct{
	Id int `json"-"`
	Name string `json:"name"`
	User string `json:"username"`
	Password string `json:"password"`
}
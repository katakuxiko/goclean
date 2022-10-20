package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/katakuxiko/clean_go/structure"
)

type AuthPostgress struct {
	db *sqlx.DB
}
func NewAuthPostgress(db *sqlx.DB) *AuthPostgress{
	return &AuthPostgress{db: db}
}
func (r *AuthPostgress) CreateUser(user structure.User)(int, error){
	var id int
	query := fmt.Sprintf("Insert into %s(name,username,passwordHash) values($1,$2,$3) Returning id",userTable )
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err 
	}

	return id, nil
}

func (r *AuthPostgress) GetUser(username, password string)(structure.User,error){
	var user structure.User
	query := fmt.Sprintf("SELECT Id from %s WHERE username = $1 AND passwordHash = $2",userTable)
	err := r.db.Get(&user, query,username, password)

	return user, err
}
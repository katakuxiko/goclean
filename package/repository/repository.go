package repository

import (
	"github.com/katakuxiko/clean_go/structure"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user structure.User)(int, error)
	GetUser(username, password string)(structure.User,error)
}
type BooksList interface {
}
type BooksItem interface {
}
type Repository struct {
	Authorization
	BooksList
	BooksItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:NewAuthPostgress(db),
	}
}
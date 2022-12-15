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
	Create(userId int, books structure.BooksList)(int, error)
	GetAll(userId int)([]structure.BooksList,error)
	GetById(userId int, id int) (structure.BooksList,error)
	Delete(userId int, id int) error
	Update(userId, listId int, input structure.UpdateListInput) error
}
type BooksItem interface {
	Create(listId int, item structure.BookdItem)(int, error)
	GetAll(userId int, listId int) ([]structure.BookdItem,error)
	GetById(userId int, itemId int) (structure.BookdItem,error)
	Delete(userId int, id int) error
	Update(userId, itemId int, input structure.UpdateItemInput) error
}
type Repository struct {
	Authorization
	BooksList
	BooksItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:NewAuthPostgress(db),
		BooksList: NewBooksListPostgres(db),
		BooksItem: NewBooksItemPostgress(db),
	}
}
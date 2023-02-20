package service

import (
	"github.com/katakuxiko/clean_go/package/repository"
	"github.com/katakuxiko/clean_go/structure"
)

type Authorization interface {
	CreateUser(user structure.User)(int, error) 
	GenerateToken(username, password string)(string, error) 
	ParseToken(token string)(int, error)
	GetUser(username, password string)(structure.User,error)
	RefreshToken(oldToken string)(string,error)
}
type BooksList interface {
	Create(userId int, books structure.BooksList)(int, error)
	GetAll(userId int)([]structure.BooksList,error)
	GetById(userId int, listId int) (structure.BooksList,error)
	Delete(userId int, listId int)error
	Update(userId, listId int, input structure.UpdateListInput) error
}
type BooksItem interface {
	Create(userId,listId int, item structure.BookdItem)(int, error)
	GetAll(userId int, listId int) ([]structure.BookdItem,error)
	GetById(userId int, itemId int) (structure.BookdItem,error)
	Delete(userId int, itemId int)error
	Update(userId, itemId int, input structure.UpdateItemInput) error

}
type Service struct {
	Authorization
	BooksList
	BooksItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		BooksList: NewBooksListService(repo.BooksList),
		BooksItem: NewBooksItemService(repo.BooksItem, repo.BooksList),
	}
}

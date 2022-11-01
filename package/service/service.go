package service

import (
	"github.com/katakuxiko/clean_go/package/repository"
	"github.com/katakuxiko/clean_go/structure"
)

type Authorization interface {
	CreateUser(user structure.User)(int, error) 
	GenerateToken(username, password string)(string, error) 
	ParseToken(token string)(int, error)
}
type BooksList interface {
	Create(userId int, books structure.BooksList)(int, error)
}
type BooksItem interface {
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
	}
}

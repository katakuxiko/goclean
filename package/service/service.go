package service

import "github.com/katakuxiko/clean_go/package/repository"

type Authorization struct {
}
type BooksList struct {
}
type BooksItem struct {
}
type Service struct {
	Authorization
	BooksList
	BooksItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
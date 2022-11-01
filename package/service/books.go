package service

import (
	"github.com/katakuxiko/clean_go/package/repository"
	"github.com/katakuxiko/clean_go/structure"
)

type BooksListService struct {
	repo repository.BooksList
}

func NewBooksListService(repo repository.BooksList) *BooksListService {
	return &BooksListService{repo: repo}
}


func (s *BooksListService) Create(userId int, list structure.BooksList)(int, error){
	return s.repo.Create(userId, list)
}
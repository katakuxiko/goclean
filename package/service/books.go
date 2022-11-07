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

func (s *BooksListService) 	GetAll(userId int)([]structure.BooksList,error){
	return s.repo.GetAll(userId)
}
func (s *BooksListService) 	GetById(userId int, listId int) (structure.BooksList,error){
	return s.repo.GetById(userId,listId)
}
func (s *BooksListService) 	Delete(userId int, listId int)error{
	return s.repo.Delete(userId,listId)
}
func (s *BooksListService)  Update(userId, listId int, input structure.UpdateListInput) error{
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, listId, input)
}
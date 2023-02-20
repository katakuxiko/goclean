package service

import (
	"github.com/katakuxiko/clean_go/package/repository"
	"github.com/katakuxiko/clean_go/structure"
)

type BooksItemService struct {
	repo repository.BooksItem
	listRepo repository.BooksList 
}

// NewBooksItemService returns new BooksItemService.
func NewBooksItemService(repo repository.BooksItem, listRepo repository.BooksList) *BooksItemService {
	return &BooksItemService{repo: repo, listRepo: listRepo}
}

func (s *BooksItemService) Create(userId,listId int, item structure.BookdItem)(int, error){
	_, err := s.listRepo.GetById(userId,listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId,item)
}
func (s *BooksItemService) GetAll(userId,listId int)([]structure.BookdItemSelect,error){
	return s.repo.GetAll(userId,listId)
}
func (s *BooksItemService) GetById(userId int, itemId int) (structure.BookdItem,error){
	return s.repo.GetById(userId,itemId)
}
func (s *BooksItemService)	Delete(userId int, itemId int)error{
	return s.repo.Delete(userId,itemId)
}
func (s *BooksItemService) 	Update(userId, itemId int, input structure.UpdateItemInput) error{
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId,itemId,input)
}

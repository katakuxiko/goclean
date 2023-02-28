package service

import (
	"strconv"
	"strings"

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
func (s *BooksItemService) GetAll(userId,listId int)([]structure.BookdItem,error){
	return s.repo.GetAll(userId,listId, -1)
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
func (s *BooksItemService)	GetItemToNextPage(userId,listId int,variables string)(int,error){
	items, err := s.repo.GetAll(userId,listId, 3);
	// fmt.Print(items)
	// var filteredItems []structure.BookdItem 
	var id int

	variable := strings.Split(variables, ",");
	for _, bi := range items {
	splitetCond := strings.Split(bi.Condition, "; ")
	var checks []bool
	for _, cond := range splitetCond {
		splitetCondName := strings.Split(cond, " ")[0]+":"
		splitetCondSymb := strings.Split(cond, " ")[1]
		splitetCondValue, err := strconv.ParseInt(strings.Split(cond, " ")[2],0,32)
		if err != nil {
			return 0,err
		}
		for _, v := range variable {
		name := strings.Split(v, " ")[1]
		value, err := strconv.ParseInt(strings.Split(v, " ")[2],0,32)
		if err != nil {
				return 0,err
			}
		switch splitetCondSymb{
		case ">":
			if(splitetCondName==name&&value>splitetCondValue){
				checks = append(checks, true)
			} else {
				checks = append(checks, false)
			}
		case "<":
			if(splitetCondName==name&&value<splitetCondValue){
				checks = append(checks, true)
			} else {
				checks = append(checks, false)
			}
		case "=":
			if(splitetCondName==name&&value==splitetCondValue){
				checks = append(checks, true)
			} else {
				checks = append(checks, false)
			}
		}

	}
		
	}
	var counter int
	for _, v := range checks {
		if v {
			counter++
		}
	}
	if counter == len(splitetCond){
		id = bi.Id
	}
}
	return id, err
}

func containsFalse(bools []bool) bool {
    result := true
    for _, b := range bools {
        if !b {
            result = false
            break
        }
    }
    return result
}
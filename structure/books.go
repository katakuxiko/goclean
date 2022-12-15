package structure

import "errors"

type BooksList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type BookdItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListItem struct {
	Id     int `json:"id"`
	ListId int `json:"list"`
	ItemId int `json:"itemId"`
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}
type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`

}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no value")
	}
	return nil
}
func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil  && i.Done == nil {
		return errors.New("update structure has no value")
	}
	return nil
}
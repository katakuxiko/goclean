package structure

import "errors"

type BooksList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	ImgUrl string `json:"img" db:"img"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type ButtonStruct struct {
	BtnName string `json:"btnName" `
	BtnAction string `json:"btnAction" `
	BtnVar string `json:"btnVar"`
}

// type Variables struct {

// }

type BookdItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
	Buttons []ButtonStruct `json:"buttons" db:"buttons"`
	Condition string `json:"condition" db:"condition"`
	Page        int   `json:"page" db:"page"`

}
type BookdItemSelect struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
	Buttons []byte `json:"buttons" db:"buttons"`
	Condition string `json:"condition" db:"condition"`
	Page        int   `json:"page" db:"page"`
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
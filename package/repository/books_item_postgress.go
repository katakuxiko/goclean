package repository

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/katakuxiko/clean_go/structure"
)

type BooksItemPostgress struct {
	db *sqlx.DB
}

// NewBooksItemPostgress returns new BooksItemPostgress.
func NewBooksItemPostgress(db *sqlx.DB) *BooksItemPostgress {
	return &BooksItemPostgress{db: db}
}

func (r *BooksItemPostgress) Create(listId int, item structure.BookdItem)(int, error){
	tx,err := r.db.Begin()
	if err != nil {
		return 0, err
	}
    btn, err := json.Marshal(&item.Buttons);
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description,buttons,condition) values ($1, $2, $3, $4) RETURNING id",booksItemTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description, btn, item.Condition)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)",listItemsTable)
	_, err = tx.Exec(createListItemQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err

	}
	return itemId, tx.Commit()
}

func (r *BooksItemPostgress) GetAll(userId int, listId int) ([]structure.BookdItem, error){
	var items []structure.BookdItemSelect
	var itemsMarshal []structure.BookdItem
	var itemsMarshalSolo structure.BookdItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done, ti.buttons, ti.condition FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1`, booksItemTable, listItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId); err != nil {
		return nil, err
	}
	var buttons []structure.ButtonStruct
	for _, item := range items {
		itemsMarshalSolo.Id = item.Id
		itemsMarshalSolo.Title = item.Title
		itemsMarshalSolo.Description = item.Description
		itemsMarshalSolo.Condition = item.Condition
		itemsMarshalSolo.Done = item.Done
		json.Unmarshal(item.Buttons, &buttons)
		itemsMarshalSolo.Buttons = buttons
		itemsMarshal = append(itemsMarshal,itemsMarshalSolo)
	}
	return itemsMarshal, nil
}
func (r *BooksItemPostgress) GetById(userId int, itemId int) (structure.BookdItem,error){
	var item structure.BookdItemSelect
	var itemsMarshal structure.BookdItem
	var buttons []structure.ButtonStruct

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done, ti.Buttons, ti.condition FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 `, booksItemTable, listItemsTable, usersListsTable)
	err := r.db.Get(&item, query, itemId); 
	json.Unmarshal(item.Buttons, &buttons)
	itemsMarshal.Id=item.Id
	itemsMarshal.Description=item.Description
	itemsMarshal.Done=item.Done
	itemsMarshal.Title=item.Title
	itemsMarshal.Condition=item.Condition
	itemsMarshal.Buttons= buttons
	if err != nil {
		return itemsMarshal, err
	}
	

	return itemsMarshal, nil
}
func(r *BooksItemPostgress)	Delete(userId int, itemId int)error{
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND
									ul.user_id = $1 and ti.id = $2`, booksItemTable, listItemsTable, usersListsTable)
	_,err := r.db.Exec(query, userId, itemId)
	return err
}
func (r *BooksItemPostgress) 	Update(userId, itemId int, input structure.UpdateItemInput) error{
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		booksItemTable, setQuery, listItemsTable,usersListsTable, argId, argId+1)
	args = append(args, itemId, userId)



	_, err := r.db.Exec(query, args...)
	return err
}

package repository

import (
	"fmt"
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

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values($1, $2) RETURNING id",booksListsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemQuery := fmt.Sprintf("INSERT INTO %s (list_id,item_id) values($1, $2) RETURNING id",listItemsTable)
	_,err = tx.Exec(createListItemQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err

	}
	return itemId, err
}

func (r *BooksItemPostgress) GetAll(userId int, listId int) ([]structure.BookdItem,error){
	var items []structure.BookdItem
	query := fmt.Sprintf(`SELECT * FROM %s tl INNER JOIN %s li ON li.item_id = ti.id 
									INNER JOIN %s ul on ul.list_id = li_list_id WHERE li.list_id = $1 AND ul.user_id = $2`, booksItemTable, booksListsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}	
	return items, nil
}

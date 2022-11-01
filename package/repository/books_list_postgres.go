package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/katakuxiko/clean_go/structure"
)

type BooksListPostgres struct {
	db *sqlx.DB
}

// NewBooksListPostgres returns new BooksListPostgres.
func NewBooksListPostgres(db *sqlx.DB) *BooksListPostgres {
	
	return &BooksListPostgres{db: db}
}

func (r *BooksListPostgres) Create(userId int, list structure.BooksList)(int, error){
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createBooksQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2)",booksListsTable)
	row := tx.QueryRow(createBooksQuery,list.Title, list.Description)
	if err = row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createListQuery := fmt.Sprintf("INSERT INTO %s (user_id,list_id) VALUES($1,$2)", usersListsTable)
	_,err = tx.Exec(createListQuery, userId,id )
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
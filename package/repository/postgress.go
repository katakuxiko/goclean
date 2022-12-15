package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/stdlib"
)
const(
	userTable = "users"
	booksListsTable = "books_list"
	booksItemTable = "books_items"
	usersListsTable = "users_list"
	listItemsTable = "list_items"
)
func NewPostgresDB(url string) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping();
	if err != nil {
		return nil, err
	}
	return db, nil
}	
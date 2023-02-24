package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/katakuxiko/clean_go/structure"
	"github.com/sirupsen/logrus"
)

type BooksListPostgres struct {
	db *sqlx.DB
}

func NewBooksListPostgres(db *sqlx.DB) *BooksListPostgres {
	
	return &BooksListPostgres{db: db}
}

func (r *BooksListPostgres) Create(userId int, list structure.BooksList)(int, error){
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createBooksQuery := fmt.Sprintf("INSERT INTO %s (title, description, img) VALUES ($1, $2, $3) RETURNING id",booksListsTable)
	row := tx.QueryRow(createBooksQuery,list.Title, list.Description,list.ImgUrl)
	if err = row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createListQuery := fmt.Sprintf("INSERT INTO %s (user_id,list_id) VALUES($1,$2)", usersListsTable)
	_,err = tx.Exec(createListQuery, userId,id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (r *BooksListPostgres) GetAll(userId int,pageParam string)([]structure.BooksList,error){
	var lists []structure.BooksList
	
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.img, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id LIMIT %s",
		booksListsTable, usersListsTable, pageParam)
	err := r.db.Select(&lists, query)

	return lists, err
}
func (r *BooksListPostgres) GetUserBooksAll(userId int)([]structure.BooksList,error){
	var lists []structure.BooksList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.img, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		booksListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}
func (r *BooksListPostgres) GetById(userId int, listId int) (structure.BooksList,error){
	var list structure.BooksList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.img FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 and ul.list_id=$2",
		booksListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId,listId)

	return list, err
}
func (r *BooksListPostgres) Delete(userId int, listId int)error{
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",booksListsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}
func (r *BooksListPostgres) Update(userId, listId int, input structure.UpdateListInput) error {
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

	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		booksListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/katakuxiko/clean_go/structure"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {

	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(userId int, userVariables structure.UsersVariables) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	} 
	var id int
	createUserVariablesQuery := fmt.Sprintf("INSERT INTO %s (user_id, variables) VALUES ($1, $2) RETURNING id", usersVariablesTable)
	row := tx.QueryRow(createUserVariablesQuery, userId,userVariables.Variables)
	if err = row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	
	
	return id, tx.Commit()
}

func (r *UserPostgres) Update(userId int, input structure.UpdateUserVariables) error {
	query := fmt.Sprintf("UPDATE %s SET variables = $1 WHERE user_id = $2", usersVariablesTable)
	_, err := r.db.Exec(query,input.Variables,userId)
	return err
}

func (r *UserPostgres) GetAllVariables(userId int) (structure.UsersVariables,error) {
	var response structure.UsersVariables

	query := fmt.Sprintf("SELECT id, user_id, variables FROM %s WHERE user_id=$1", usersVariablesTable)
	err := r.db.Get(&response,query,userId);
	return response,err
	
}
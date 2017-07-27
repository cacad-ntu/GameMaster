package storage

import (
	"database/sql"
	"../models"
)

var addUserQuery string = "replace into User values(:name, :hashed_password)"
var getUserQuery string = "select * from User where name = ?"
var listUsersQuery string = "select * from User"
var deleteUserQuery string = "delete from User where name = ?"

func (db *dbImpl) CreateUser(user models.User) error {
    _, err := db.sqliteDB.NamedExec(addUserQuery, user)
    return err
}

func (db *dbImpl) GetUser(name string) (*models.User, error) {
    result := models.User{}
    err := db.sqliteDB.Get(&result, getUserQuery, name)
    
    if err == sql.ErrNoRows {
        return nil, nil
    }

    return &result, err
}

func (db *dbImpl) ListUsers() ([]models.User, error) {
	res := []models.User{}

	err := db.sqliteDB.Select(&res, listUsersQuery)
	return res, err
}

func (db *dbImpl) DeleteUser(name string) error {
	_, err := db.sqliteDB.Exec(deleteUserQuery, name)
	return err
}
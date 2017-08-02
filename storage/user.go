package storage

import (
	"database/sql"

	"github.com/cacad-ntu/GameMaster/models"
)

var addUserQuery string = "replace into User values(:user_name, :hashed_password)"
var getUserQuery string = "select * from User where user_name = ?"
var listUsersQuery string = "select * from User"
var deleteUserQuery string = "delete from User where user_name = ?"

func (db *dbImpl) CreateUser(user models.User) error {
	_, err := db.sqliteDB.NamedExec(addUserQuery, user)
	return err
}

func (db *dbImpl) GetUser(userName string) (*models.User, error) {
	result := models.User{}
	err := db.sqliteDB.Get(&result, getUserQuery, userName)

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

func (db *dbImpl) DeleteUser(userName string) error {
	_, err := db.sqliteDB.Exec(deleteUserQuery, userName)
	return err
}

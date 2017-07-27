package storage

import (
	"database/sql"
	"../models"
)

var addGroupQuery string = "replace into Group values(:name, :hashed_password)"
var getGroupQuery string = "select * from Group where name = ?"
var listGroupsQuery string = "select * from Group"
var deleteGroupQuery string = "delete from Group where name = ?"

func (db *dbImpl) CreateGroup(group models.Group) error {
    _, err := db.sqliteDB.NamedExec(addGroupQuery, group)
    return err
}

func (db *dbImpl) GetGroup(name string) (*models.Group, error) {
    result := models.Group{}
    err := db.sqliteDB.Get(&result, getGroupQuery, name)
    
    if err == sql.ErrNoRows {
        return nil, nil
    }

    return &result, err
}

func (db *dbImpl) ListGroups() ([]models.Group, error) {
	res := []models.Group{}

	err := db.sqliteDB.Select(&res, listGroupsQuery)
	return res, err
}

func (db *dbImpl) DeleteGroup(name string) error {
	_, err := db.sqliteDB.Exec(deleteGroupQuery, name)
	return err
}
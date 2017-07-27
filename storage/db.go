package storage

import (
    "github.com/jmoiron/sqlx"
    _ "github.com/mattn/go-sqlite3"
    "../models"
)
// TODO: Use singleton for db
type DB interface {
    CreateUser(user models.User) error
    GetUser(name string) (*models.User, error)
    ListUsers() ([]models.User, error)
    DeleteUser(name string) error
    CreateGroup(group models.Group) error
    GetGroup(name string) (*models.Group, error)
    ListGroups() ([]models.Group, error)
    DeleteGroup(name string) error
}

type dbImpl struct {
    sqliteDB *sqlx.DB
}

func NewDB(fileName string) (DB, error) {
    db, err := sqlx.Open("sqlite3", fileName)
    if err != nil {
        return nil, err
    }
    db.MustExec("PRAGMA foreign_keys = ON;")
    initDBSchema(db)
    return &dbImpl{db}, err
}
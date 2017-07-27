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
package storage

import (
	"github.com/jmoiron/sqlx"
)

var createUserTableQuery string = `create table if not exists User
(
    name TEXT,
    hashed_password TEXT,
    PRIMARY KEY(name)
)`

var createGroupTableQuery string = `create table if not exists Group
(
    name TEXT,
    hashed_password TEXT,
    PRIMARY KEY(name)
)`

func initDBSchema(db *sqlx.DB) error {
	_, err := db.Exec(createUserTableQuery)
    _, err = db.Exec(createGroupTableQuery)
	return err
}
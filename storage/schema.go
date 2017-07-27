package storage

import (
	"github.com/jmoiron/sqlx"
)

var createUserTableQuery string = `create table if not exists User
(
    user_name TEXT,
    hashed_password TEXT,
    PRIMARY KEY(user_name)
)`

var createTeamTableQuery string = `create table if not exists Team
(
    id TEXT,
    name TEXT,
    hashed_password TEXT,
    PRIMARY KEY(id)
)`

var createEventTableQuery string = `create table if not exists Event
(
    id TEXT,
    name TEXT,
    PRIMARY KEY(id)
)`

var createUserEventTableQuery string = `create table if not exists UserEvent
(
    id TEXT,
    user_name TEXT REFERENCES User(user_name),
    event_id TEXT REFERENCES Event(id),
    PRIMARY KEY(id)
)`

var createEventTeamTableQuery string = `create table if not exists EventTeam
(
    id TEXT,
    event_id TEXT REFERENCES Event(id),
    team_id TEXT REFERENCES Team(id),
    PRIMARY KEY(id)
)`

var createGameTableQuery string = `create table if not exists Game
(
    id TEXT,
    name TEXT,
    event_id TEXT REFERENCES Event(id),
    score INTEGER,
    PRIMARY KEY(id)
)`

var createTeamGameTableQuery string = `create table if not exists TeamGame
(
    id TEXT,
    team_id TEXT REFERENCES Team(id),
    game_id TEXT REFERENCES Game(id),
    PRIMARY KEY(id)
)`

func initDBSchema(db *sqlx.DB) error {
	_, err := db.Exec(createUserTableQuery)
    _, err = db.Exec(createTeamTableQuery)
    _, err = db.Exec(createEventTableQuery)
    _, err = db.Exec(createUserEventTableQuery)
    _, err = db.Exec(createEventTeamTableQuery)
    _, err = db.Exec(createGameTableQuery)
    _, err = db.Exec(createTeamGameTableQuery)
	return err
}
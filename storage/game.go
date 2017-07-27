package storage

import (
	"database/sql"
	"../models"
	"github.com/satori/go.uuid"
)

var addGameQuery string = "replace into Game values(:id, :name, :event_id, :score)"
var getGameQuery string = "select * from Game where id = ?"
var listGamesQuery string = "select * from Game"
var deleteGameQuery string = "delete from Game where id = ?"

func (db *dbImpl) CreateGame(game models.Game) error {
    _, err := db.sqliteDB.NamedExec(addGameQuery, game)
    return err
}

func (db *dbImpl) GetGame(id uuid.UUID) (*models.Game, error) {
    result := models.Game{}
    err := db.sqliteDB.Get(&result, getGameQuery, id)
    
    if err == sql.ErrNoRows {
        return nil, nil
    }

    return &result, err
}

func (db *dbImpl) ListGames() ([]models.Game, error) {
	res := []models.Game{}

	err := db.sqliteDB.Select(&res, listGamesQuery)
	return res, err
}

func (db *dbImpl) DeleteGame(id uuid.UUID) error {
	_, err := db.sqliteDB.Exec(deleteGameQuery, id)
	return err
}
package storage

import (
	"database/sql"
    "github.com/satori/go.uuid"
	"../models"
)

var addTeamQuery string = "replace into Team values(:id, :name, :hashed_password)"
var getTeamQuery string = "select * from Team where id = ?"
var listTeamsQuery string = "select * from Team"
var deleteTeamQuery string = "delete from Team where id = ?"

func (db *dbImpl) CreateTeam(team models.Team) error {
    _, err := db.sqliteDB.NamedExec(addTeamQuery, team)
    return err
}

func (db *dbImpl) GetTeam(id uuid.UUID) (*models.Team, error) {
    result := models.Team{}
    err := db.sqliteDB.Get(&result, getTeamQuery, id)
    
    if err == sql.ErrNoRows {
        return nil, nil
    }

    return &result, err
}

func (db *dbImpl) ListTeams() ([]models.Team, error) {
	res := []models.Team{}

	err := db.sqliteDB.Select(&res, listTeamsQuery)
	return res, err
}

func (db *dbImpl) DeleteTeam(id uuid.UUID) error {
	_, err := db.sqliteDB.Exec(deleteTeamQuery, id)
	return err
}
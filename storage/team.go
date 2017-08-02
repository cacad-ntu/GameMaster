package storage

import (
	"database/sql"

	"github.com/cacad-ntu/GameMaster/models"
)

var addTeamQuery string = "replace into Team values(:id, :name, :hashed_password)"
var getTeamQuery string = "select * from Team where name = ?"
var listTeamsQuery string = "select * from Team"
var deleteTeamQuery string = "delete from Team where name = ?"

func (db *dbImpl) CreateTeam(team models.Team) error {
	_, err := db.sqliteDB.NamedExec(addTeamQuery, team)
	return err
}

func (db *dbImpl) GetTeam(name string) (*models.Team, error) {
	result := models.Team{}
	err := db.sqliteDB.Get(&result, getTeamQuery, name)

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

func (db *dbImpl) DeleteTeam(name string) error {
	_, err := db.sqliteDB.Exec(deleteTeamQuery, name)
	return err
}

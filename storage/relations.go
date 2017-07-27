package storage

import (
	"github.com/satori/go.uuid"
)

var addUserEventQuery string = "replace into UserEvent values(?, ?, ?)"
var addEventTeamQuery string = "replace into EventTeam values(?, ?, ?)"
var addTeamGameQuery string = "replace into TeamGame values(?, ?, ?)"


func (db *dbImpl) CreateUserEvent(id uuid.UUID, user_name string, event_id uuid.UUID) error {
    _, err := db.sqliteDB.Exec(addUserEventQuery, id, user_name, event_id)
    return err
}

func (db *dbImpl) CreateEventTeam(id uuid.UUID, event_id uuid.UUID, team_id uuid.UUID) error {
    _, err := db.sqliteDB.Exec(addEventTeamQuery, id, event_id, team_id)
    return err
}

func (db *dbImpl) CreateTeamGame(id uuid.UUID, team_id uuid.UUID, game_id uuid.UUID) error {
    _, err := db.sqliteDB.Exec(addTeamGameQuery, id, team_id, game_id)
    return err
}
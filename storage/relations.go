package storage

import (
	"github.com/satori/go.uuid"
    "../models"
)

var addUserEventQuery string = "replace into UserEvent values(?, ?, ?)"
var addEventTeamQuery string = "replace into EventTeam values(?, ?, ?)"
var addTeamGameQuery string = "replace into TeamGame values(?, ?, ?)"

var listUsersFromEventQuery string = "select user_name from UserEvent where event_id = ?"
var deleteUserFromEventQuery string = "delete from UserEvent where user_name = ?"

var listTeamsFromEventQuery string = "select team_id from EventTeam where event_id = ?"
var deleteTeamFromEventQuery string = "delete from EventTeam where team_id = ?"

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

func (db *dbImpl) ListUsersFromEvent(event_id uuid.UUID) ([]models.User, error) {
    res := []string{}

	err := db.sqliteDB.Select(&res, listUsersFromEventQuery, event_id)
    users := []models.User{}
    for _, elem := range res {
        user, _ := db.GetUser(elem)
        users = append(users, *user)
    }
	return users, err
}

func (db *dbImpl) DeleteUserFromEvent(userName string) error {
    _, err := db.sqliteDB.Exec(deleteUserFromEventQuery, userName)
	return err

}

func (db *dbImpl) ListTeamsFromEvent(event_id uuid.UUID) ([]models.Team, error) {
    res := []uuid.UUID{}

    err := db.sqliteDB.Select(&res, listTeamsFromEventQuery, event_id)
    teams := []models.Team{}
    for _, elem := range res {
        team, _ := db.GetTeam(elem)
        teams = append(teams, *team)
    }
    return teams, err
}

func (db *dbImpl) DeleteTeamFromEvent(team_id uuid.UUID) error {
    _, err := db.sqliteDB.Exec(deleteTeamFromEventQuery, team_id)
    return err
}
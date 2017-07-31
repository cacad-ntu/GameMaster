package storage

import (
    "github.com/jmoiron/sqlx"
    "github.com/satori/go.uuid"
    _ "github.com/mattn/go-sqlite3"
    "../models"
)
// TODO: Use singleton for db
type DB interface {
    CreateUser(user models.User) error
    GetUser(userName string) (*models.User, error)
    ListUsers() ([]models.User, error)
    DeleteUser(userName string) error
    CreateTeam(team models.Team) error
    GetTeam(id uuid.UUID) (*models.Team, error)
    ListTeams() ([]models.Team, error)
    DeleteTeam(id uuid.UUID) error
    CreateEvent(event models.Event) error
    GetEvent(id uuid.UUID) (*models.Event, error)
    ListEvents() ([]models.Event, error)
    DeleteEvent(id uuid.UUID) error
    CreateGame(game models.Game) error
    GetGame(id uuid.UUID) (*models.Game, error)
    ListGames() ([]models.Game, error)
    DeleteGame(id uuid.UUID) error
    CreateUserEvent(id uuid.UUID, user_name string, event_id uuid.UUID) error
    CreateEventTeam(id uuid.UUID, event_id uuid.UUID, team_id uuid.UUID) error
    CreateTeamGame(id uuid.UUID, team_id uuid.UUID, game_id uuid.UUID) error
    ListUsersFromEvent(event_id uuid.UUID) ([]models.User, error)
    DeleteUserFromEvent(userName string) error
    ListTeamsFromEvent(event_id uuid.UUID) ([]models.Team, error)
    DeleteTeamFromEvent(team_id uuid.UUID) error
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
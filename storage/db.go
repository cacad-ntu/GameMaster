package storage

import (
	"github.com/cacad-ntu/GameMaster/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/satori/go.uuid"
)

// TODO: Use singleton for db
type DB interface {
	CreateUser(user models.User) error
	GetUser(userName string) (*models.User, error)
	ListUsers() ([]models.User, error)
	DeleteUser(userName string) error
	CreateTeam(team models.Team) error
	GetTeam(name string) (*models.Team, error)
	ListTeams() ([]models.Team, error)
	DeleteTeam(name string) error
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

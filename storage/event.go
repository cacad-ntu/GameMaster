package storage

import (
	"database/sql"
	"../models"
	"github.com/satori/go.uuid"
)

var addEventQuery string = "replace into Event values(:id, :name)"
var getEventQuery string = "select * from Event where id = ?"
var listEventsQuery string = "select * from Event"
var deleteEventQuery string = "delete from Event where id = ?"

func (db *dbImpl) CreateEvent(event models.Event) error {
    _, err := db.sqliteDB.NamedExec(addEventQuery, event)
    return err
}

func (db *dbImpl) GetEvent(id uuid.UUID) (*models.Event, error) {
    result := models.Event{}
    err := db.sqliteDB.Get(&result, getEventQuery, id)
    
    if err == sql.ErrNoRows {
        return nil, nil
    }

    return &result, err
}

func (db *dbImpl) ListEvents() ([]models.Event, error) {
	res := []models.Event{}

	err := db.sqliteDB.Select(&res, listEventsQuery)
	return res, err
}

func (db *dbImpl) DeleteEvent(id uuid.UUID) error {
	_, err := db.sqliteDB.Exec(deleteEventQuery, id)
	return err
}
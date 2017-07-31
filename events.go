package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"./storage"
	"./models"
)

func CreateEvent(eventName string) (uuid.UUID, error) {
	db, err := storage.NewDB("./test.sqlite3")
    if err != nil {
	    fmt.Printf(err.Error())
    }

	id := uuid.NewV4()
	event := models.Event{id, eventName}
	err = db.CreateEvent(event)
	
	if err == nil {
		return id, nil
	}
	return uuid.UUID{}, err
}

func GetEvent(eventID uuid.UUID) *models.Event {
	db, err := storage.NewDB("./test.sqlite3")
    if err != nil {
	    fmt.Printf(err.Error())
    }
	res, err := db.GetEvent(eventID)
	return res
}

func ListEvents() []models.Event {
	db, err := storage.NewDB("./test.sqlite3")
    if err != nil {
	    fmt.Printf(err.Error())
    }
	res, err := db.ListEvents()
	return res
}

func DeleteEvent(eventID uuid.UUID) error {
	db, err := storage.NewDB("./test.sqlite3")
    if err != nil {
	    fmt.Printf(err.Error())
    }
	err = db.DeleteEvent(eventID)
	return err
}

func AddGMToEvent(userName string, eventID uuid.UUID) error {
	db, err := storage.NewDB("./test.sqlite3")
    if err != nil {
	    fmt.Printf(err.Error())
    }

	id := uuid.NewV4()
	err = db.CreateUserEvent(id, userName, eventID)
	return err
}

func GetGMsFromEvent(eventID uuid.UUID) []models.User {
	db, err := storage.NewDB("./test.sqlite3")
    if err != nil {
	    fmt.Printf(err.Error())
    }

	res, err := db.ListUsersFromEvent(eventID)
	return res
}

func RemoveGMFromEvent(userName string) error {
	db, err := storage.NewDB("./test.sqlite3")
    if err != nil {
	    fmt.Printf(err.Error())
    }

	err = db.DeleteUserFromEvent(userName)
	return err
}

func AddTeamToEvent(eventID uuid.UUID, teamID uuid.UUID) error {
	db, err := storage.NewDB("./test.sqlite3")
    if err != nil {
	    fmt.Printf(err.Error())
    }

	id := uuid.NewV4()
	err = db.CreateEventTeam(id, eventID, teamID)
	return err
}

func GetTeamsFromEvent(eventID uuid.UUID) []models.Team {
	db, err := storage.NewDB("./test.sqlite3")
    if err != nil {
	    fmt.Printf(err.Error())
    }

	res, err := db.ListTeamsFromEvent(eventID)
	return res
}

func RemoveTeamFromEvent(teamID uuid.UUID) error {
	db, err := storage.NewDB("./test.sqlite3")
    if err != nil {
	    fmt.Printf(err.Error())
    }

	err = db.DeleteTeamFromEvent(teamID)
	return err
}
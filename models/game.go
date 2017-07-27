package models

import (
    "github.com/satori/go.uuid"
)

type Game struct {
    Id uuid.UUID `json:"id" db:"id"`
    Name string `json:"name" db:"name"`
	EventId uuid.UUID `json:"event_id" db:"event_id"`
    Score int `json:"score" db:"score"`
}
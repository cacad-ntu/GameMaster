package models

import (
    "github.com/satori/go.uuid"
)

type Event struct {
    Id uuid.UUID `json:"id" db:"id"`
    Name string `json:"name" db:"name"`
}
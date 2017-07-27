package models

import (
    "github.com/satori/go.uuid"
)

// TODO: does team need password? is it for login? seems redundant
type Team struct {
    Id uuid.UUID `json:"id" db:"id"`
    Name string `json:"name" db:"name"`
    HashedPassword []byte `json:"hashedPassword" db:"hashed_password"`
}
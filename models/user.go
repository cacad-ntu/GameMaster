package models

type User struct {
    Name string `json:"name" db:"name"`
    HashedPassword []byte `json:"hashedPassword" db:"hashed_password"`
}
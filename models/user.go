package models

type User struct {
    UserName string `json:"user_name" db:"user_name"`
    HashedPassword []byte `json:"hashed_password" db:"hashed_password"`
}
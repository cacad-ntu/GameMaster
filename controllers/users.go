package controllers

import (
	"fmt"

	"github.com/cacad-ntu/GameMaster/models"
	"github.com/cacad-ntu/GameMaster/storage"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(userName string, password []byte) error {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	db, err := storage.NewDB("./test.sqlite3")
	if err != nil {
		fmt.Printf(err.Error())
	}

	user := models.User{userName, hashedPassword}
	err = db.CreateUser(user)
	return err
}

func Login(userName string, password []byte) (bool, error) {
	db, err := storage.NewDB("./test.sqlite3")
	if err != nil {
		fmt.Printf(err.Error())
	}

	user, _ := db.GetUser(userName)

	err = bcrypt.CompareHashAndPassword(user.HashedPassword, password)

	// nil means it is a match
	if err != nil {
		return false, err
	}
	return true, nil
}

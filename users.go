package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(userName string, password []byte) error {
	// Hashing the password with the default cost of 10
    hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
	
	// TODO: store in DB
	fmt.Println(string(hashedPassword))
	return err
}

func Login(userName string, password []byte) (bool, error) {
	// TODO: retrieve data from DB and match them
	// for now it is hardcoded that the correct password is "test", feel free to change it for dev purpose
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)

    // nil means it is a match
	if (err != nil) {
		return false, err
	}
	return true, nil
}

package handler

import (
	"net/http"

	"github.com/cacad-ntu/GameMaster/controllers"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	secretKey = "cacad-ntu"
	issuer    = "cacad-ntu"
)

//LoginClaims to generate token
type LoginClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// SignupHandler handle http request to signup as new user
func SignupHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		renderTemplate(res, "signup", nil)
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")
	confirmPassword := req.FormValue("confirm-password")

	if password != confirmPassword {
		//TODO: Better error message when password missmatch
		//check on jquery
		http.Redirect(res, req, "/signup", http.StatusFound)
		return
	}

	err := controllers.SignUp(username, []byte(password))
	if err != nil {
		//TODO: Better error message when failed to signup
		http.Redirect(res, req, "/signup", http.StatusFound)
		return
	}

	token := generateToken(username)
	setToken(res, token)

	http.Redirect(res, req, "/home", http.StatusSeeOther)
	return
}

// LoginHandler handle http request to login for existing user
func LoginHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		renderTemplate(res, "login", nil)
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	ok, err := controllers.Login(username, []byte(password))
	if !ok || err != nil {
		//TODO: Better error message when failed to login
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}

	token := generateToken(username)
	setToken(res, token)

	http.Redirect(res, req, "/home", http.StatusSeeOther)
	return
}

// LogoutHandler handle user log out
func LogoutHandler(res http.ResponseWriter, req *http.Request, username string) {
	RemoveToken(res)
	http.Redirect(res, req, "/login", http.StatusSeeOther)
	return
}

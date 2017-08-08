package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// Templates store html template that can be rendered
	Templates *template.Template
)

type view struct {
	Username       string
	ProfilePictURL string
	DisplayName    string
}

func renderTemplate(res http.ResponseWriter, tmpl string, viewFields *view) {
	if viewFields != nil && viewFields.ProfilePictURL == "" {
		viewFields.ProfilePictURL = "default.png"
	}

	err := Templates.ExecuteTemplate(res, tmpl+".html", viewFields)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func generateToken(username string) string {
	expireToken := time.Now().Add(time.Hour * 1).Unix()

	claims := LoginClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte(secretKey))

	return signedToken
}

func setToken(res http.ResponseWriter, token string) {
	expireCookie := time.Now().Add(time.Hour * 1)
	cookie := http.Cookie{Name: "Auth", Value: token, Expires: expireCookie, HttpOnly: true}
	http.SetCookie(res, &cookie)
}

// GetToken retrieve token from cookie with "Auth" name
func GetToken(req *http.Request) string {
	cookie, err := req.Cookie("Auth")
	if err != nil {
		return "none"
	}
	return cookie.Value
}

// RemoveToken replace cookies with empty token
func RemoveToken(res http.ResponseWriter) {
	deleteCookie := http.Cookie{Name: "Auth", Value: "none", Expires: time.Now()}
	http.SetCookie(res, &deleteCookie)
}

// ValidateToken and return the corespondent username
func ValidateToken(strToken string) (string, bool) {
	if strToken == "none" {
		return "", false
	}

	token, err := jwt.ParseWithClaims(strToken, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	claims, ok := token.Claims.(*LoginClaims)
	if !ok {
		return "", false
	}

	return claims.Username, err == nil && token.Valid
}

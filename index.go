package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"

	"github.com/cacad-ntu/GameMaster/handler"
)

const (
	fileServerDir = "resources"
)

var (
	httpServerPort *string
)

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fn(res, req)
	}
}

func makeSecureHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		token := handler.GetToken(req)
		username, ok := handler.ValidateToken(token)
		if !ok {
			handler.RemoveToken(res)
			http.Redirect(res, req, "/login", http.StatusFound)
			return
		}
		fn(res, req, username)
	}
}

func init() {
	httpServerPort = flag.String("port", "8080", "Port to serve http")
	flag.Parse()

	handler.Templates = template.Must(template.ParseFiles("html/home.html", "html/login.html", "html/signup.html"))
}

func main() {
	// Router function handler
	http.HandleFunc("/", makeSecureHandler(handler.ViewHomeHandler))
	http.HandleFunc("/home", makeSecureHandler(handler.ViewHomeHandler))
	http.HandleFunc("/login", makeHandler(handler.LoginHandler))
	http.HandleFunc("/signup", makeHandler(handler.SignupHandler))
	http.HandleFunc("/logout", makeSecureHandler(handler.LogoutHandler))

	// File server
	fs := http.FileServer(http.Dir(fileServerDir))
	http.Handle("/resources/", http.StripPrefix("/resources/", fs))

	fmt.Println("Serving http server at :", *httpServerPort)
	http.ListenAndServe(":"+*httpServerPort, nil)
}

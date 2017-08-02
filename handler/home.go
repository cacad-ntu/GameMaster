package handler

import (
	"net/http"
)

// ViewHomeHandler handle http request to view home page
func ViewHomeHandler(res http.ResponseWriter, req *http.Request, username string) {

	viewFields := &view{
		Username:    username,
		DisplayName: username,
	}

	renderTemplate(res, "home", viewFields)
	return
}

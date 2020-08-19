package api

import (
	"net/http"
	"rohandvivedi.com/src/session"
)

// this file provides owner (and only owner) with apis to access user sessions

// api handlers in this file
var PrintAllUserSessions = http.HandlerFunc(printAllUserSessions)

func printAllUserSessions(w http.ResponseWriter, r *http.Request) {
	session.PrintAllSessionValues()
}
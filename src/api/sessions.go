package api

import (
	"net/http"
	"rohandvivedi.com/src/session"
	"encoding/json"
)

// this file provides owner (and only owner) with apis to access user sessions

// api handlers in this file
var PrintAllUserSessions = http.HandlerFunc(printAllUserSessions)

func printAllUserSessions(w http.ResponseWriter, r *http.Request) {
	if(session.GlobalSessionStore == nil) {
		return
	}

	session.GlobalSessionStore.Lock.Lock()
	json, _ := json.Marshal(session.GlobalSessionStore.Sessions);
	session.GlobalSessionStore.Lock.Unlock()
	w.Write(json);
}
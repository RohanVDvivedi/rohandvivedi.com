package chatter

import (
	"net/http"
	/*"encoding/json"
	"golang.org/x/net/websocket"
	"time"
	"fmt"*/
)

import (
	"rohandvivedi.com/src/session"
	"rohandvivedi.com/src/config"
)

func AuthorizeChat(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		nameList, existsName := r.URL.Query()["name"];

		if(config.GetGlobalConfig().Create_user_sessions) {
			s := session.GlobalSessionStore.GetExistingSession(r);
			if(s != nil) {
				// if a name is present in the request, store it in the session values
				if(existsName) {
					s.SetValue("name", nameList[0]);
				}
				// allow the request to be served only if the name exists in the session values
				_, nameSessionExists := s.GetValue("name");
				if(nameSessionExists) {
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		// if any thing fails, just unautorize
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You are not authorized to chat on rohandvivedi.com"))
	})
}
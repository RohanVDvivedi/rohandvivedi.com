package chatter

import (
	"net/http"
	/*"encoding/json"
	"golang.org/x/net/websocket"
	"time"
	"fmt"*/
)

import (
	//"rohandvivedi.com/src/session"
	"rohandvivedi.com/src/config"
)

func AuthorizeChat(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// in the dev env, auth is to be bypassed
		if(config.GetGlobalConfig().Environment == "dev") {
			next.ServeHTTP(w, r)
			return
		}

		/*nameList, existsName := r.URL.Query()["name"];

		if(config.GetGlobalConfig().Create_user_sessions) {
			s := session.GlobalSessionStore.GetExistingSession(r);
			if(s != nil) {
				name, nameExists := s.GetValue("name");
				if(nameExists) {
					valName, ok := name.(bool)
					if(ok && valIsOwner){
						next.ServeHTTP(w, r)
						return
					}
				}
			}
		}*/

		// if any thing fails, just unautorize
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You are not authorized to chat on rohandvivedi.com"))
	})
}
package chatter

import (
	"net/http"
	"golang.org/x/net/websocket"
)

import (
	"rohandvivedi.com/src/session"
	"rohandvivedi.com/src/config"
)

var AuthorizeAndStartChatHandler = AuthorizeChat(websocket.Handler(ChatConnectionHandler))

func AuthorizeChat(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		nameList, existsName := r.URL.Query()["name"];

		if(config.GetGlobalConfig().Create_user_sessions) {
			s := session.GlobalSessionStore.GetExistingSession(r);
			if(s != nil) {
				if(existsName) { // if a name is present in the request, store it in the session values
					s.SetValue("name", nameList[0]);
				}

				// allow the request to be served only if the name exists in the session values
				// and a corresponding chat session does not exist
				nameIntr, nameSessionExists := s.GetValue("name");
				if(nameSessionExists) {
					name, ok := nameIntr.(string)
					if(ok) {
						next.ServeHTTP(w, r)
						return
					}
				}
			}
		}

		// if any thing fails, just unautorize
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You are not authorized to chat on rohandvivedi.com"))
	})
}
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
		publicKeyList, existsPublicKeyList := r.URL.Query()["publicKey"];

		if(config.GetGlobalConfig().Create_user_sessions) {
			s := session.GlobalSessionStore.GetExistingSession(r);
			if(s != nil) {
				if(existsName && existsPublicKeyList) {
					InsertNameAndPublicKeyToSession(s, nameList[0], publicKeyList[0])
				} else if ( (existsName && !existsPublicKeyList) || (!existsName && existsPublicKeyList) ) {
					RemoveNameAndPublicKeyFromSession(s)
				}

				next.ServeHTTP(w, r)
				return
			}
		}

		// if any thing fails, just unautorize
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You are not authorized to chat on rohandvivedi.com"))
	})
}

func GetNameAndPublicKeyFromSession(s *session.Session) (string, string, bool) {
	nameIntr, foundName := s.GetValue("name")
	publicKeyIntr, foundPublicKey := s.GetValue("publicKey")
	if(foundName && foundPublicKey) {
		name, nameOk := nameIntr.(string)
		publicKey, publicKeyOk := publicKeyIntr.(string)
		if(nameOk && publicKeyOk) {
			return name, publicKey, true
		}
	}
	RemoveNameAndPublicKeyFromSession(s)
	return "", "", false
}

func InsertNameAndPublicKeyToSession(s *session.Session, name string, publicKey string) {
	s.SetValue("name", name);
	s.SetValue("publicKey", publicKey);
}

func RemoveNameAndPublicKeyFromSession(s *session.Session) {
	s.RemoveValue("name");
	s.RemoveValue("publicKey");
}
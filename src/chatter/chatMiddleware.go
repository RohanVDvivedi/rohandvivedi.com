package chatter

import (
	"net/http"
	"golang.org/x/net/websocket"
)

import (
	"rohandvivedi.com/src/session"
)

var AuthorizeAndStartChatHandler = AuthorizeChat(websocket.Handler(ChatConnectionHandler))

func AuthorizeChat(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		nameList, existsName := r.URL.Query()["name"];
		publicKeyList, existsPublicKeyList := r.URL.Query()["publicKey"];

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

		// if any thing fails, just unautorize
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You are not authorized to chat on rohandvivedi.com"))
	})
}

func GetNameAndPublicKeyFromSession(s *session.Session) (string, string, bool) {
	name := ""
	publicKey := ""
	found := false
	s.ExecuteOnValues(func (values map[string]interface{}, additional_params interface{}) interface{} {
		nameIntr, foundName := values["name"]
		publicKeyIntr, foundPublicKey := values["publicKey"]
		if(foundName && foundPublicKey) {
			nameV, nameOk := nameIntr.(string)
			publicKeyV, publicKeyOk := publicKeyIntr.(string)
			if(nameOk && publicKeyOk) {
				name = nameV
				publicKey = publicKeyV
				found = true
			}
		}
		return nil
	}, nil)
	return name, publicKey, found
}

func InsertNameAndPublicKeyToSession(s *session.Session, name string, publicKey string) {
	s.ExecuteOnValues(func (values map[string]interface{}, additional_params interface{}) interface{} {
		values["name"] = name;
		values["publicKey"] = publicKey;
		return nil
	}, nil)
}

func RemoveNameAndPublicKeyFromSession(s *session.Session) {
	s.ExecuteOnValues(func (values map[string]interface{}, additional_params interface{}) interface{} {
		delete(values, "name");
		delete(values, "publicKey");
		return nil
	}, nil)
}
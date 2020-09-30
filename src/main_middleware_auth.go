package main

// go utilities
import (
	"net/http"
)

// maintains session, (in memory)
import (
	"rohandvivedi.com/src/session"
)


func AuthorizeIfOwner(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// you must need a session to allow me to check if you are the owner of the website
		s := session.GlobalSessionStore.GetExistingSession(r);
		if(s != nil) {
			isOwner, ownerKeyExists := s.GetValue("owner");
			if(ownerKeyExists) {
				valIsOwner, ok := isOwner.(bool)
				if(ok && valIsOwner){
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		// if any thing fails, just unautorize
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You are not an authorized owner of rohandvivedi.com"))

	})
}

func AuthorizeIfHasSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if(session.GlobalSessionStore.GetExistingSession(r) != nil) {
			next.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You are not an authorized to use this functionality of rohandvivedi.com, without a valid session id"))
	})
}
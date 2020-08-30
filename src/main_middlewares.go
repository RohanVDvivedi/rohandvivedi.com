package main

// go utilities
import (
	"net/http"
	"strings"
)

// maintains global configuration for the application
import (
	"rohandvivedi.com/src/config"
)

// maintains session, (in memory)
import (
	"rohandvivedi.com/src/session"
)

// this function is a middleware to send 404 response, if the requested path is a folder
// i.e. request path ending in "/"
func Send404OnFolderRequest(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if strings.HasSuffix(r.URL.Path, "/") && len(r.URL.Path) != 1 {
            http.NotFound(w, r)
            return
        }
        next.ServeHTTP(w, r)
    })
}

// this middleware lets you maintain data regarding api access that each sessioned user has made
func CountApiHitsInSessionValues(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
        if(config.GetGlobalConfig().Create_user_sessions) {
	        // you must need a session to allow me to maintain the count
	        s := session.GlobalSessionStore.GetOrCreateSession(w, r);
			_ = s.ExecuteOnValues(func (SessionVals map[string]interface{}, add_par interface{}) interface{} {
				reqPathCountKey := "<" + r.URL.Path + ">_count"		// this is the key we will use to store count of hits in session values
				count, exists := SessionVals[reqPathCountKey]
				if(exists) {
					intCount, isInt := count.(int)
					if isInt {
						SessionVals[reqPathCountKey] = intCount + 1
						return nil
					}
				}
				SessionVals[reqPathCountKey] = 1
				return nil
			}, nil);
		}

        next.ServeHTTP(w, r)
    })
}


func AuthorizeIfOwner(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// in the dev env, autorization for owner is bypassed
		if(config.GetGlobalConfig().Environment == "dev") {
			next.ServeHTTP(w, r)
			return
		}
        
        if(config.GetGlobalConfig().Create_user_sessions) {
	        // you must need a session to allow me to check if you are the owner of the website
	        s := session.GlobalSessionStore.GetExistingSession(w, r);
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
		}

		// if any thing fails, just unautorize
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You are not an authorized owner of rohandvivedi.com"))

    })
}

func AuthorizeIfHasSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// in the dev env, autorization for owner is bypassed
		if(config.GetGlobalConfig().Environment == "dev") {
			next.ServeHTTP(w, r)
			return
		}
        
        if(config.GetGlobalConfig().Create_user_sessions) {
	        // you must need a session to allow me to allow you to reach the handler function
	        s := session.GlobalSessionStore.GetExistingSession(w, r);
	        if(s != nil) {
				next.ServeHTTP(w, r)
				return
			}
		}

		// if any thing fails, just unautorize
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You are not authorized to access this api of rohandvivedi.com"))

    })
}
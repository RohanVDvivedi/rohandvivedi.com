package main

// go utilities
import (
	"time"
	"net/http"
	"strings"
	"strconv"
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

// this function is a middleware to send 404 response, if the requested path is a folder
// i.e. request path ending in "/"
func SetRequestCacheControl(maxAge time.Duration, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Cache-Control", "public, max-age=" + strconv.Itoa(int(maxAge.Seconds())))
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

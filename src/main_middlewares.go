package main

// go utilities
import (
	"time"
	"net/http"
	"strings"
	"strconv"
)

import (
	"rohandvivedi.com/src/useractlogger"
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

// this middleware lets you maintain log data regarding api access that each sessioned user has made
func LogUserActivity(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        s := session.GlobalSessionStore.GetExistingSession(r);
        sessionId := "UNKNOWN_OR_NEW_USER"
        if(s != nil) {
            sessionId = s.SessionId
        }
        go useractlogger.LogUserActivity(sessionId, r.URL.Path, r.URL.RawQuery);
        next.ServeHTTP(w, r)
    })
}

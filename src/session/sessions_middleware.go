package session

import (
	"time"
	"net/http"
)

var GlobalSessionStore *SessionStore = nil;

func InitGlobalSessionStore(CookieName string, MaxLifeDuration time.Duration) {
	GlobalSessionStore = NewSessionStore(CookieName, MaxLifeDuration)
}

func SessionManagerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := GlobalSessionStore.acquireSession(w,r);
		next.ServeHTTP(w, r)
		GlobalSessionStore.releaseAcquiredSession(s);
	})
}
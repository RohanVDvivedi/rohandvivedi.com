package session

import (
	"time"
	"sync"
	"net/http"
)

/******************************/
import "math/rand"
func InitRand() {
    rand.Seed(time.Now().UnixNano())
}
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
/******************************/

// this is a per user per session struct
// the Values map will store a map to key values that we will use to store certain data corresponding to the user
type Session struct {
	SessionId		string
	FirstAccessed	time.Time
	LastAccessed	time.Time
	Lock			sync.Mutex
	Values			map[string]interface{}
}

type SessionStore struct {
	CookieName		string 					// name of the cookie, against which session_id will be stored on client side
	MaxLifeDuration	time.Duration 			// this is the maximum time a cookie session can and will survive	
											// it will be set on cookie's expiry, aswell as any cookie unused for MaxLifeTime will expire on server side
	Lock			sync.Mutex				// lock to make Sessions map thread safe
	Sessions		map[string]*Session		// the map -> it stores Session.SessionId vs Session
	// session id creator function
}

var GlobalSessionStore *SessionStore = nil;

func InitGlobalSessionStore(CookieName string, MaxLifeDuration time.Duration) {
	GlobalSessionStore = &SessionStore{
		CookieName: CookieName, 
		MaxLifeDuration: MaxLifeDuration,
		Sessions: make(map[string]*Session),
	};
	InitRand()
}

func GetOrCreateSession(w http.ResponseWriter, r *http.Request) *Session {
	// if the application SessionStore is not setup, before getting the first request
	// of if the owner does not want the user sessions to be maintained
	if(GlobalSessionStore == nil) {
		return nil
	}

	sessionCookie, errUserCookie := r.Cookie(GlobalSessionStore.CookieName)

	GlobalSessionStore.Lock.Lock()

	if(errUserCookie == nil) {	// no error means client remembers the session_id as cookie
		session_id := sessionCookie.Value

		// use the session id to find corresponding session
		session, errServerSession := GlobalSessionStore.Sessions[session_id]

		if(errServerSession) { // if a session is found, return it
			GlobalSessionStore.Lock.Unlock()
			return session;
		}
	}

	// create a new session id
	session_id := RandStringBytes(12)

	// create a new session
	session := &Session{
		SessionId: session_id,
		FirstAccessed: time.Now(),
		LastAccessed: time.Now(),
		Values: make(map[string]interface{}),
	}

	// create a new cookie
	cookie := http.Cookie{
		Name: GlobalSessionStore.CookieName,
		Value: session.SessionId,
		Expires: session.FirstAccessed.Add(GlobalSessionStore.MaxLifeDuration),
		HttpOnly: true,
	}

	// set a ccookie on the http response
	http.SetCookie(w, &cookie)

	// store the corresponding session
	GlobalSessionStore.Sessions[session.SessionId] = session

	GlobalSessionStore.Lock.Unlock()
	
	return session
}

func (s *Session) GetValue(key string) (interface{}, bool) {
	s.Lock.Lock()
	s.LastAccessed =time.Now()
	val, err := s.Values[key]
	s.Lock.Unlock()
	return val, err
}

func (s *Session) SetValue(key string, value interface{}) {
	s.Lock.Lock()
	s.LastAccessed =time.Now()
	s.Values[key] = value
	s.Lock.Unlock()
}

func (s *Session) ExecuteOnValues(operation_function func (values map[string]interface{}, additional_params interface{}) interface{}, additional_params interface{}) interface{} {
	s.Lock.Lock()
	s.LastAccessed = time.Now()
	result := operation_function(s.Values, additional_params)
	s.Lock.Unlock()
	return result
} 
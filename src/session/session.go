package session

import (
	"time"
	"sync"
)

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
	MaxLifeTime		time.Time 				// this is the maximum time a cookie session can and will survive	
											// it will be set on cookie's expiry, aswell as any cookie unused for MaxLifeTime will expire on server side
	Lock			sync.Mutex				// lock to make Sessions map thread safe
	Sessions		map[string]*Session		// the map -> it stores Session.SessionId vs Session
	// session id creator function
}

var GlobalSessionStore *SessionStore = nil;

func InitGlobalSessionStore(CookieName string, MaxLifeTime time.Time) {
	GlobalSessionStore = &SessionStore{
		CookieName: CookieName, 
		MaxLifeTime: MaxLifeTime,
		Sessions: make(map[string]*Session)
	};
}

func GetOrCreateSession(w http.ResponseWriter, r *http.Request) *Session {

	sessionCookie, errUserCookie := r.Cookie(GlobalSessionStore.CookieName)

	SessionStore.Lock.Lock()

	if(errUserCookie == nil) {	// no error means client remembers the session_id as cookie
		session_id := sessionCookie.Value

		// use the session id to find corresponding session
		session, errServerSession := SessionStore.Sessions[session_id]

		if(errServerSession == nil) { // if a session is found, return it
			SessionStore.Lock.Unlock()
			return session;
		}
	}

	// create a new session id
	session_id := ""

	// create a new session
	session := &Session{
		SessionId: session_id,
		FirstAccessed: time.Now(),
		LastAccessed: time.Now(),
		Values: make(map[string]interface{})
	}

	// create a new cookie
	cookie := http.Cookie{
		Name: GlobalSessionStore.CookieName,
		Value: session.SessionId,
		Expires: session.FirstAccessed.Add(GlobalSessionStore.MaxLifeTime),
		HttpOnly: true
	}

	// set a ccookie on the http response
	http.SetCookie(w, cookie)

	// store the corresponding session
	SessionStore.Sessions[session.SessionId] = session

	SessionStore.Lock.Unlock()
	
	return session
}

func (s *Session) GetValue(key string) interface{}, Error {
	s.Lock.Lock()
	s.LastAccessed =time.Now()
	val, err := s.Values[key]
	s.Lock.Unlock()
	return val
}

func (s *Session) SetValue(key string, value interface{}) {
	s.Lock.Lock()
	s.LastAccessed =time.Now()
	s.Values[key] = value
	s.Lock.Unlock()
}

func (s *Session) ExecuteOnValues(
		// a function to operate  on values of this session
		operation func (values map[string]interface{}, additional_params interface{}) interface{}, 

		// may be some additional values to the operation
		additional_params interface{}
	) 
	// this function will return what ever is returned from the operator
	interface{} {

	s.Lock.Lock()
	s.LastAccessed = time.Now()
	result := operation(s.Values, additional_params)
	s.Lock.Unlock()
	
	return result
} 
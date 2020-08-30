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
const letterBytes = "+-/<[abcdefghijklmnopqrstuvwxyz](ABCDEFGHIJKLMNOPQRSTUVWXYZ){0123456789}>-_^#"
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

// unsafe
func (ss *SessionStore) createNewUniquelyRandomSessionId() string {
	// create a new session id composed of 32 character random string, 
	// which does not exist in SessionStore before
	sessionIdLength := 32
	sessionId := RandStringBytes(sessionIdLength)
	_, sessionIdExists := ss.Sessions[sessionId]
	for(sessionIdExists) {
		sessionId = RandStringBytes(sessionIdLength)
		_, sessionIdExists = ss.Sessions[sessionId]
	}

	return sessionId;
}

// unsafe
func (ss *SessionStore) createNewSession() *Session {
	// create a new session
	session := &Session{
		SessionId: ss.createNewUniquelyRandomSessionId(),
		FirstAccessed: time.Now(),
		LastAccessed: time.Now(),
		Values: make(map[string]interface{}),
	}
	return session
}

// unsafe
func (ss *SessionStore) createNewSessionCookie(session *Session) *http.Cookie {
	// create a new cookie
	cookie := &http.Cookie{
		Name: ss.CookieName,
		Value: session.SessionId,
		Expires: session.FirstAccessed.Add(ss.MaxLifeDuration),
		HttpOnly: true,
		Path: "/",
	}
	return cookie
}

func (ss *SessionStore) InitializeOwnerSession() *Session {
	ss.Lock.Lock()

	// create a new session
	session := ss.createNewSession();

	// store the corresponding session, in the session store
	ss.Sessions[session.SessionId] = session

	session.SetValue("owner", true)

	ss.Lock.Unlock()

	return session;
}

func (ss *SessionStore) GetExistingSession(w http.ResponseWriter, r *http.Request) *Session {
	sessionCookie, errUserCookie := r.Cookie(ss.CookieName)

	ss.Lock.Lock()

	if(errUserCookie == nil) {	// no error means client remembers the session_id as cookie
		sessionId := sessionCookie.Value

		// use the session id to find corresponding session
		session, serverSessionExists := ss.Sessions[sessionId]

		if(serverSessionExists) { // if a session is found, return it
			ss.Lock.Unlock()
			return session;
		}
	}

	ss.Lock.Unlock()

	return nil
}

func (ss *SessionStore) GetOrCreateSession(w http.ResponseWriter, r *http.Request) *Session {

	sessionCookie, errUserCookie := r.Cookie(ss.CookieName)

	ss.Lock.Lock()

	if(errUserCookie == nil) {	// no error means client remembers the session_id as cookie
		sessionId := sessionCookie.Value

		// use the session id to find corresponding session
		session, serverSessionExists := ss.Sessions[sessionId]

		if(serverSessionExists) { // if a session is found, return it
			ss.Lock.Unlock()
			return session;
		}
	}

	// create a new session
	session := ss.createNewSession();

	// store the corresponding session, in the session store
	ss.Sessions[session.SessionId] = session

	// create a new cookie
	cookie := ss.createNewSessionCookie(session)

	// set a ccookie on the http response
	http.SetCookie(w, cookie)

	ss.Lock.Unlock()
	
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
package session

import (
	"time"
	"sync"
	"net/http"
)

// this is a per user session struct
// the Values map will store a map to key values that we will use to store certain data corresponding to the user
type Session struct {
	SessionId			string

	tLock				sync.Mutex				// lock to protect the LastAccessed time stamps and the ActiveRequestCount

	LastAccessed		time.Time
	ActiveRequestCount	uint64 					// this is incremented as soon as a new request is hit to the server
												// a session may not be deleted as long as it has atlest 1 active request
	
	vLock				sync.Mutex				// lock to protect the session values

	Values				map[string]interface{}

	next	*Session
	prev	*Session
}

// unsafe
func (ss *SessionStore) createNewSession() *Session {
	return &Session{
		SessionId: ss.createNewUniquelyRandomSessionId(),
		LastAccessed: time.Now(),
		ActiveRequestCount: 0
		Values: make(map[string]interface{}),
		next: nil,
		prev: nil,
	}
}

// unsafe
func (ss *SessionStore) createNewSessionCookie(session *Session) *http.Cookie {
	return &http.Cookie{
		Name: ss.CookieName,
		Value: session.SessionId,
		Expires: session.LastAccessed.Add(ss.MaxLifeDuration),
		HttpOnly: true,
		Path: "/",
	}
}

// unsafe
func (s *Session) acquireSession() {
	s.LastAccessed = time.Now();
	s.ActiveRequestCount += 1
}

// unsafe
func (s *Session) releaseSession() {
	s.ActiveRequestCount -= 1
}

/* to operate on the session, the api handler may call only the below provided functions */

func (s *Session) GetValue(key string) (interface{}, bool) {
	s.vLock.Lock()
	val, err := s.Values[key]
	s.vLock.Unlock()
	return val, err
}

func (s *Session) SetValue(key string, value interface{}) {
	s.vLock.Lock()
	s.Values[key] = value
	s.vLock.Unlock()
}

func (s *Session) RemoveValue(key string) {
	s.vLock.Lock()
	delete(s.Values, key);
	s.vLock.Unlock()
}

func (s *Session) ExecuteOnValues(operation_function func (values map[string]interface{}, additional_params interface{}) interface{}, additional_params interface{}) interface{} {
	s.vLock.Lock()
	result := operation_function(s.Values, additional_params)
	s.vLock.Unlock()
	return result
} 
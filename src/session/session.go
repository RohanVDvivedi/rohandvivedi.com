package session

import "time"
import "sync"

// this is a per user per session struct
// the Values map will store a map to key values that we will use to store certain data corresponding to the user
type Session struct {
	SessionId		string
	FirstAccessed	time.Time
	LastAccessed	time.Time
	Lock			sync.Mutex
	Values			map[interface{}]interface{}
}

type SessionStore struct {
	CookieName		string 					// name of the cookie, against which session_id will be stored on client side
	MaxLifeTime		time.Time 				// this is the maximum time a cookie session can and will survive	
											// it will be set on cookie's expiry, aswell as any cookie unused for MaxLifeTime will expire on server side
	Lock			sync.RWMutex			// lock to make Sessions map thread safe
	Sessions		map[string]Session		// the map -> it stores Session.SessionId vs Session
	// session id creator function
}

func InitSessionStore(string CookieName, MaxLifeTime time.Time) *SessionStore {

}

func (ss *SessionStore) GetSessionValuesIfExists() *Session {
	return nil
}

func (ss *SessionStore) GetOrCreateSessionValues() *Session {
	return nil
}

func (s *Session) GetValue(key interface{}) interfae{} {
	return nil
}

func (s *Session) SetValue(key interface{}, value interface{}) {
	
}

func (s *Session) ExecuteOnValues(func (values map[interface{}]interface{}, additional_params interface{})) {
} 
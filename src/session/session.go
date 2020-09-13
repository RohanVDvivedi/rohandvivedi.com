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

	next	*Session
	prev	*Session
}

type SessionStore struct {
	CookieName		string 					// name of the cookie, against which session_id will be stored on client side
	MaxLifeDuration	time.Duration 			// this is the maximum time a cookie session can and will survive	
											// it will be set on cookie's expiry, aswell as any cookie unused for MaxLifeTime will expire on server side
	Lock			sync.Mutex				// lock to make Sessions map thread safe
	Sessions		map[string]*Session		// the map -> it stores Session.SessionId vs Session
	
	listLRU struct {						// a doubly linkedlist of session values
		head 		*Session 				// head = Least recently used session
		tail 		*Session 				// tail = Most recently used session
	}
}

var GlobalSessionStore *SessionStore = nil;

func InitGlobalSessionStore(CookieName string, MaxLifeDuration time.Duration) {
	GlobalSessionStore = &SessionStore{
		CookieName: CookieName, 
		MaxLifeDuration: MaxLifeDuration,
		Sessions: make(map[string]*Session),
		listLRU: struct {head *Session 
		tail *Session}{head: nil, tail: nil},
	};
	go GlobalSessionStore.GarbageCollectionRoutine()
	InitRand()
}

// unsafe
func (ss *SessionStore) createNewUniquelyRandomSessionId() string {
	// create a new session id composed of 32 character random string, 
	// which does not exist in SessionStore before
	sessionIdLength := 40
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
		next: nil,
		prev: nil,
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

/* Below functions are required to manage lru */

// unsafe
func (ss *SessionStore) removeSessionFromLRU(session *Session) {
	if(session.next != nil) {
		session.next.prev = session.prev
	}
	if(session.prev != nil) {
		session.prev.next = session.next
	}
	if(ss.listLRU.head == session) {
		ss.listLRU.head = session.next
	}
	if(ss.listLRU.tail == session) {
		ss.listLRU.head = session.prev
	}
	session.next = nil
	session.prev = nil
}

// unsafe, the session must not already exist in lru
func (ss *SessionStore) insertSessiontoLRUtail(session *Session) {
	tailSess := ss.listLRU.tail
	if(tailSess == nil) {
		ss.listLRU.head = session
	} else {
		tailSess.next = session
		session.prev = tailSess
		session.next = nil
	}
	ss.listLRU.tail = session
}

// unsafe
func (ss *SessionStore) getSessionFromLRUhead() *Session {
	return ss.listLRU.head
}

/* Above functions are required to manage lru */

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

func (ss *SessionStore) GetExistingSession(r *http.Request) *Session {
	sessionCookie, errUserCookie := r.Cookie(ss.CookieName)

	ss.Lock.Lock()

	if(errUserCookie == nil) {	// no error means client remembers the session_id as cookie
		sessionId := sessionCookie.Value

		// use the session id to find corresponding session
		session, serverSessionExists := ss.Sessions[sessionId]

		if(serverSessionExists) { // if a session is found, return it
			// if a server session exists and is found, you need to put it at the 
			ss.removeSessionFromLRU(session)
			ss.insertSessiontoLRUtail(session)

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
			// if a server session exists and is found, you need to put it at the 
			ss.removeSessionFromLRU(session)
			ss.insertSessiontoLRUtail(session)

			ss.Lock.Unlock()
			return session;
		}
	}

	// create a new session, and insert it at the tail of the lru
	session := ss.createNewSession();
	ss.insertSessiontoLRUtail(session)

	// store the corresponding session, in the session store
	ss.Sessions[session.SessionId] = session

	// create a new cookie
	cookie := ss.createNewSessionCookie(session)

	// set a ccookie on the http response
	http.SetCookie(w, cookie)

	ss.Lock.Unlock()
	
	return session
}

func (ss *SessionStore) GarbageCollectionRoutine() {
	for (true) {
		loop_exit := false
		sessionsProcessed := 0

		for(!loop_exit) {
			ss.Lock.Lock()

			LRUhead := ss.getSessionFromLRUhead()
			if(LRUhead == nil) {
				loop_exit = true
			} else {
				LRUhead.Lock.Lock()

				/*valOwnerIntr, foundOwnerKey := s.Values["owner"]
				valOwner, isBoolvalOwnerIntr := valOwnerIntr.(bool)

				valChatActiveIntr, foundChatActiveKey := s.Values["chat_active"]
				valChatActive, isBoolvalChatActiveIntr := valChatActiveIntr.(bool)*/

				hasSessionExpiryElapsed := time.Now().Sub(LRUhead.LastAccessed) > ss.MaxLifeDuration

				// bump the session in the list if it is an owner session or an active chat session
				/*if((foundOwnerKey && isBoolvalOwnerIntr && valOwner) ||
					(foundChatActiveKey && isBoolvalChatActiveIntr && valChatActive)) {
					ss.removeSessionFromLRU(LRUhead)
					ss.insertSessiontoLRUtail(LRUhead)
				}*/
				// remove if the session has not been accessed for more than its life time amount of time
				/*else*/
				if(hasSessionExpiryElapsed) {
					ss.removeSessionFromLRU(LRUhead)
				} else if (!hasSessionExpiryElapsed) {
					loop_exit = true
				}
				
				LRUhead.Lock.Unlock()
			}

			ss.Lock.Unlock()
		}
		time.Sleep(5 * time.Minute)	// garbage collection running every 3 minutes
	}
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

func (s *Session) RemoveValue(key string) {
	s.Lock.Lock()
	s.LastAccessed =time.Now()
	delete(s.Values, key);
	s.Lock.Unlock()
}

func (s *Session) ExecuteOnValues(operation_function func (values map[string]interface{}, additional_params interface{}) interface{}, additional_params interface{}) interface{} {
	s.Lock.Lock()
	s.LastAccessed = time.Now()
	result := operation_function(s.Values, additional_params)
	s.Lock.Unlock()
	return result
} 
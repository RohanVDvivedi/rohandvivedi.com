package session

import (
	"time"
	"sync"
	"net/http"
	"rohandvivedi.com/src/randstring"
)

type SessionStore struct {
	CookieName		string 					// name of the cookie, against which session_id will be stored on client side
	MaxLifeDuration	time.Duration 			// this is the maximum time a cookie session can and will survive	
											// it will be set on cookie's expiry, aswell as any cookie unused for MaxLifeTime will expire on server side
	sLock			sync.Mutex				// lock to make Sessions map thread safe
	Sessions		map[string]*Session		// the map -> it stores Session.SessionId vs Session
	
	listLRU struct {						// a doubly linkedlist of session values
		head 		*Session 				// head = Least recently used session
		tail 		*Session 				// tail = Most recently used session
	}
}

func NewSessionStore(CookieName string, MaxLifeDuration time.Duration) *SessionStore {
	ss := &SessionStore{
		CookieName: CookieName, 
		MaxLifeDuration: MaxLifeDuration,
		Sessions: make(map[string]*Session),
		listLRU: struct {head *Session 
		tail *Session}{head: nil, tail: nil},
	};
	go ss.garbageCollectionRoutine()
	return ss
}

func (ss *SessionStore) GetExistingSession(r *http.Request) *Session {
	sessionCookie, errUserCookie := r.Cookie(ss.CookieName)

	ss.sLock.Lock()

	if(errUserCookie == nil) {	// no error means client remembers the session_id as cookie
		sessionId := sessionCookie.Value

		// use the session id to find corresponding session
		session, serverSessionExists := ss.Sessions[sessionId]

		if(serverSessionExists) { // if a session is found, return it
			ss.sLock.Unlock()
			return session;
		}
	}

	ss.sLock.Unlock()

	return nil
}

func (ss *SessionStore) acquireSession(w http.ResponseWriter, r *http.Request) *Session {

	sessionCookie, errUserCookie := r.Cookie(ss.CookieName)

	ss.sLock.Lock()
		// session pointer we found or we created
		var session *Session = nil

		// try to find the session using the user provided session value
		if(errUserCookie == nil) {
			sessionFound, serverSessionExists := ss.Sessions[sessionCookie.Value]
			if(serverSessionExists) {
				session = sessionFound
			}
		}
		// else if not found, create a new session, and insert it in the session store
		if(session == nil) {
			session = ss.createNewSession();
			ss.Sessions[session.SessionId] = session
			ss.insertSessiontoLRUtail(session)
		}

		if(session != nil) {
			session.tLock.Lock()

				if(session.ActiveRequestCount == 0) {
					ss.removeSessionFromLRU(session)
				}
				session.acquireSession()

			session.tLock.Unlock()
		}

	ss.sLock.Unlock()

	// create a new cookie, with updated expiry, from the session itself and send it
	cookie := ss.createNewSessionCookie(session)
	http.SetCookie(w, cookie)
	
	return session
}

func (ss *SessionStore) releaseAcquiredSession(session *Session) {
	ss.sLock.Lock()

		sessionFound, serverSessionExists := ss.Sessions[session.SessionId]
		if(serverSessionExists) {

			sessionFound.tLock.Lock()

				sessionFound.releaseSession()
				if(sessionFound.ActiveRequestCount == 0) {
					ss.insertSessiontoLRUtail(sessionFound)
				}

			sessionFound.tLock.Unlock()

		}

	ss.sLock.Unlock()
}

func (ss *SessionStore) garbageCollectionRoutine() {
	for (true) {
		loop_exit := false

		for(!loop_exit) {
			ss.sLock.Lock()

			LRUhead := ss.getSessionFromLRUhead()
			if(LRUhead == nil) {
				loop_exit = true
			} else {
				LRUhead.tLock.Lock()

				hasSessionExpiryElapsed := time.Now().Sub(LRUhead.LastAccessed) > ss.MaxLifeDuration

				if(hasSessionExpiryElapsed && LRUhead.ActiveRequestCount == 0) {
					ss.removeSessionFromLRU(LRUhead)
					delete(ss.Sessions, LRUhead.SessionId);
				} else if (LRUhead.ActiveRequestCount > 0) {
					ss.removeSessionFromLRU(LRUhead)
				} else {
					loop_exit = true
				}
				
				LRUhead.tLock.Unlock()
			}

			ss.sLock.Unlock()
		}
		time.Sleep(5 * time.Minute)	// garbage collection running every 5 minutes
	}
}

// unsafe
func (ss *SessionStore) createNewUniquelyRandomSessionId() string {
	const sessionIdLength = 40

	sessionId := randstring.GetRandomString(sessionIdLength)
	_, sessionIdExists := ss.Sessions[sessionId]
	for(sessionIdExists) {
		sessionId = randstring.GetRandomString(sessionIdLength)
		_, sessionIdExists = ss.Sessions[sessionId]
	}

	return sessionId;
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
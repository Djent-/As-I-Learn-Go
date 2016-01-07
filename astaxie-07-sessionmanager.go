package main

import (
	"net/http"
	"net/url"
	"html/template"
	"sync"
	"fmt"
	"io"
	"crypto/rand"
	"encoding/base64"
)

type Manager struct {
	cookieName  string     // private cookie
	lock        sync.Mutex // protects session
	provider    Provider
	maxlifetime int64
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	// check to make sure the provider actually provides that provideName
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	// Return a new Manager
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

type Provider interface {
	// implements the initialization of a session
	SessionInit(sid string) (Session, error)
	// returns a session by corresponding sid
	// makes a new session if the given one does not exist
	SessionRead(sid string) (Session, error)
	// given an sid, deletes the corresponding session
	SessionDestroy(sid string) error
	// deletes expired session variables
	SessionGC(maxLifeTime int64)
}

type Session interface {
	// There are only four operations for sessions
	Set(key, value interface{}) error // set session value
	Get(key interface{}) interface{}  // get session value
	Delete(key interface{}) error     // delete session value
	SessionID() string                // get current sessionID
}

/* Not sure if needed. Book is not clear */
var provides = make (map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provide " + name)
	}
	provides[name] = provider
}
/**/

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	// lock the mutex and unlock it when done
	manager.lock.Lock()
	defer manager.lock.Unlock()
	// get the cookie from the request
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		// get the session from the session id
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
		return session
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ := manager.provider.SessionRead(sid)
		return session
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		// Set the session and redirect the user to /
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}

// Not sure this goes here
// Moved from inside main()
var globalSessions *Manager

func main() {
	globalSessions, err := NewManager("memory", "gosessionid", 3600)

}

func init() {
}

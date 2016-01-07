package main

type Manager struct {
	cookieName string // private cookie
	lock sync.Mutex // protects session
	provider Provider
	maxlifetime int64
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	// check to make sure the provider actually provides that provideName
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	// Return a new Manager
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}
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
	Get(key interface{}) interface{} // get session value
	Delete(key interface{}) error // delete session value
	SessionID() string // get current sessionID
}

/* Not sure if needed. Book is not clear *\
var provides = make (map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provide is nil")
	}
	if _. dup := provides[name]; dup {
		panic("session: Register called twice for provide " + name)
	}
	provides[name] = provider
}
**/

func (manager *Manager) sessionId() string {

}

func main() {
	// Not sure this goes here
	var globalSessions *session.Manager
}

func init() {
	globalSessions = NewManager("memory", "gosessionid", 3600)
}


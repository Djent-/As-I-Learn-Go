package memory

import (
	"container/list"
	"github.com/astaxie/session"
	"sync"
	"time"
)

var pder = &Provider{list: list.New()}

type SessionStore struct {
	sid string // unique session id
	timeAccessed time.Time // last access
	value map[interface{}]interface{}
}

func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
	// why does this need to return error?
}

func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
	return nil
}

func (st *SessionStore) Delete(key interface{}) error {
	delete()
}
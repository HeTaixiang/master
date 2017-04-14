package db

import (
	"sync"

	"github.com/HeTaixiang/master/utils"
	mgo "gopkg.in/mgo.v2"
)

// ProcessFn execute the real database operation
type ProcessFn func(c *mgo.Collection) error

// SessionInterface manager session
type SessionInterface interface {
	Destroy()
	ExecuteFor(collectionName string, processFn ProcessFn) error
}

type sessionManager struct {
	sessionMux sync.Mutex
	session    *mgo.Session
}

// NewManager create the instace of SessionManager
func NewManager(config *utils.DBConfig) {

}

// Destroy destroy global session
func (s *sessionManager) Destroy() {
	s.sessionMux.Lock()
	defer s.sessionMux.Unlock()
	if s.session == nil {
		return
	}
	s.session.Close()
	s.session = nil
}

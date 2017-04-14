package db

import (
	"errors"
	"sync"

	"fmt"

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
	sessionMux *sync.Mutex
	config     *utils.DBConfig
	session    *mgo.Session
}

// NewManager create the instace of SessionManager
func NewManager(c *utils.DBConfig) (SessionInterface, error) {
	if c == nil {
		return nil, errors.New("config must not nil")
	}
	if len(c.DB) == 0 || len(c.URL) == 0 || len(c.Password) == 0 || len(c.Password) == 0 {
		return nil, errors.New("database params is invalid")
	}
	url := fmt.Sprintf("mongodb://%s:%s@%s", c.Username, c.Password, c.URL)
	s, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	manager := &sessionManager{
		sessionMux: new(sync.Mutex),
		config:     c,
		session:    s}
	return manager, nil
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

func (s *sessionManager) ExecuteFor(collectionName string, processFn ProcessFn) error {
	sessionClone := s.session.Clone()
	defer sessionClone.Close()
	c := sessionClone.DB(s.config.DB).C(collectionName)
	return processFn(c)
}

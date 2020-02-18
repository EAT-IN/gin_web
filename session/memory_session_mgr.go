package session

import (
	"errors"
	"github.com/satori/go.uuid"
	"sync"
)

type MemorySessionMgr struct {
	sessionMap map[string]interface{}
	relock     sync.RWMutex
}

func NewMemorySessionMgr() *MemorySessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]interface{}, 1024),
	}
	return sr
}

func (s *MemorySessionMgr) Init(addr string, options ...string) error {
	return nil
}

func (s *MemorySessionMgr) CreateSession() (Session, error) {
	s.relock.Lock()
	defer s.relock.Unlock()
	// 生成uuid
	// github.com/satori/go.uuid
	id := uuid.NewV4()
	sessionID := id.String()
	// 创建seesion对象, 因为MemorySession实现了session的所有接口，所以也就是接口类型了
	session := NewMemorySession(sessionID)
	// session.data是一个map用来存session里面的信息
	// 然后还有一个大map用来存储素有的session。key为sessionid，值为单个的session对象
	s.sessionMap[sessionID] = session
	return session, nil

}

func (s *MemorySessionMgr) Get(sessionID string) (Session, error) {
	// 用seesionid来回去单个的session对象，并取出里面的值
	s.relock.Lock()
	defer s.relock.Unlock()
	session, ok := s.sessionMap[sessionID].(Session)
	if !ok {
		return nil, errors.New("session not exist")
	}
	return session, nil

}

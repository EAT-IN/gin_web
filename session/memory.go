package session

import (
	"errors"
	"sync"
)

// 把session存入内存
type MemorySession struct {
	sessionID string
	data      map[string]interface{}
	rwlock    sync.RWMutex
}

func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		sessionID: id,
		data:      make(map[string]interface{}, 16),
		rwlock:    sync.RWMutex{},
	}
	return s
}

func (m *MemorySession) Set(key string, value interface{}) error {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	// 设置session
	m.data[key] = value
}

func (m *MemorySession) Get(key string) (interface{}, error) {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	value, ok := m.data[key]
	if !ok {
		return nil, errors.New("key not exists in session")
	}
	return value, nil
}

func (m *MemorySession) Del(key string) error {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data, key)
	return nil
}

func (m *MemorySession) Save() error {
	return nil
}

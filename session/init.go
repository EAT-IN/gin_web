package session

import (
	"errors"
)

var (
	sessionMgr SessionMgr
)

func Init(provider string, addr string, options ...string) error {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		return errors.New("不支持的方式")
	}
	err := sessionMgr.Init(addr, options...)
	if err != nil {
		return err
	}
	return nil
}

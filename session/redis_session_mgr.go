package session

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

//redis seesionMgr接口实现类
//Init(addr string, options ...string) error
//CreateSession() (Session, error)
//Get(sessionID string) (Session, error)

type RedisSessionMgr struct {
	// todo redis地址
	addr string
	// todo 密码
	passwd string
	// todo 连接池
	pool *redis.Pool
	// todo 锁
	rwlock sync.RWMutex
	// todo 大map
	sessionMap map[string]interface{}
}

func NewRedisSessionMgr() *RedisSessionMgr {
	sr := &RedisSessionMgr{
		sessionMap: make(map[string]interface{}, 16),
	}
	return sr
}

func (r *RedisSessionMgr) Init(addr string, options ...string) error {
	// 如果有参数
	if len(options) > 0 {
		r.passwd = options[0]
	}
	// 创建连接池
	r.pool = myPoll(addr, r.passwd)
	r.addr = addr
	return nil

}

func myPoll(addr, password string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}

			// 进行密码的判断
			if _, err := conn.Do("AUTH", password); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow:    nil,
		MaxIdle:         64,
		MaxActive:       1000,
		IdleTimeout:     time.Second * 3,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}

func (r *RedisSessionMgr) CreateSession() (Session, error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	id := uuid.NewV4()

	sessionID := id.String()
	session := NewRedisSession(sessionID, r.pool)
	// 把session保存到大map中，也就是内存中
	r.sessionMap[sessionID] = session
	return session, nil
}
func (r *RedisSessionMgr) Get(sessionID string) (Session, error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	session, ok := r.sessionMap[sessionID].(Session)
	if !ok {
		err := errors.New("session not exists")
		return nil, err
	}
	return session, nil
}

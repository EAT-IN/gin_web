package session

import (
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"sync"
)

//redis seesion接口实现类

type RedisSession struct {
	sessionID string
	// redis连接池对象
	redisPool *redis.Pool
	// 存入session值信息
	sessionMap map[string]interface{}
	rwlock     sync.RWMutex
	// 用来标记map是否被操作
	flag int
}

// 用常量定义状态
const (
	// 内存数据没变化
	SessionFlagNone = iota
	SessionFlagModify
)

//构造函数
func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	s := &RedisSession{
		sessionID:  id,
		redisPool:  pool,
		sessionMap: make(map[string]interface{}, 16),
		flag:       SessionFlagNone,
	}
	return s
}

//接下来就要实现session接口里面四个方法了
//Set(key string, value interface{}) error
//Get(key string) (interface{}, error)
//Del(key string) error
//Save() error

//session储存到内存中
func (r *RedisSession) Set(key string, value interface{}) error {
	//因为涉及map读写 并发不安全需要枷锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	// 设置值
	r.sessionMap[key] = value
	r.flag = SessionFlagModify
	return nil
}
func (r *RedisSession) Get(key string) (interface{}, error) {
	//因为涉及map读写 并发不安全需要枷锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 取值的话 先判断一下内存有没有，如果没有再去redis里面取值
	result, ok := r.sessionMap[key]
	if !ok {
		err := errors.New("key not exists")
		return nil, err
	}
	return result, nil
}

func (r *RedisSession) loadFromRedis() error {
	//因为涉及map读写 并发不安全需要枷锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	conn := r.redisPool.Get()
	reply, err := conn.Do("GET", r.sessionID)
	if err != nil {
		return err
	} else {
		data, err := redis.String(reply, err)
		if err != nil {
			return err
		} else {
			// 把redis中的json数据反序列化到map中方便代码取值,反加载到sessio的map中
			err = json.Unmarshal([]byte(data), &r.sessionMap)
			if err != nil {
				return err
			} else {
				return nil
			}
		}
	}
}

func (r *RedisSession) Del(key string) error {
	//因为涉及map读写 并发不安全需要枷锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	r.flag = SessionFlagModify
	delete(r.sessionMap, key)
	return nil
}

//把内存中的session存储到redis中
func (r *RedisSession) Save() (err error) {
	//因为涉及map读写 并发不安全需要枷锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	if r.flag != SessionFlagModify {
		return
	}
	//把内存中的seesionmap进行序列化成json对面，方便redis k-v存储
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}

	//接下来进行redis的存储操作
	conn := r.redisPool.Get()
	_, err = conn.Do("SET", r.sessionID, string(data))
	return

}

package session

// 管理所有session 并列举出来相应的方法, 只要相应的结构体实现下面的方法就是session管理者接口
type SessionMgr interface {
	Init(addr string, options ...string) error
	CreateSession() (Session, error)
	Get(sessionID string) (Session, error)
}

package session

// 定义接口，并列举出来相应的方法, 只要相应的结构体实现下面的方法就是session接口
// 之所以使用接口来管理，是因为session有内存 redis mongo等等实现方法，但是他们这些工具类都是拥有同样的方法
type Session interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Del(key string) error
	Save() error
}

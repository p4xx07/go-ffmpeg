package lock

import "sync"

type Lock[T any] struct {
	Value *T
	mutex sync.Mutex
}

func Invoke[T any, TK any](lock *Lock[T], f func(*T) TK) TK {
	lock.mutex.Lock()
	defer lock.mutex.Unlock()
	return f(lock.Value)
}

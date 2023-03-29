package log

import (
	"errors"
	"sync"
)

var (
	mu          sync.RWMutex
	ns          map[string]*Logger
	ErrNotExist = errors.New("namespaced logger does not exist")
)

func init() {
	ns = make(map[string]*Logger)
}

func GetNSLogger(namespace string) (*Logger, error) {
	mu.RLock()
	defer mu.RUnlock()
	if l, ok := ns[namespace]; ok {
		return l, nil
	}
	return Default(), ErrNotExist
}

func SetNSLogger(namespace string, logger *Logger) {
	mu.Lock()
	defer mu.Unlock()
	ns[namespace] = logger
}

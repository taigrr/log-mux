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

// GetNSLogger returns the logger for the given namespace.
// If the namespace does not exist, a new default logger is created and
// ErrNotExist is returned alongside the newly created logger.
func GetNSLogger(namespace string) (*Logger, error) {
	mu.RLock()
	if l, ok := ns[namespace]; ok {
		mu.RUnlock()
		return l, nil
	}
	mu.RUnlock()

	mu.Lock()
	defer mu.Unlock()
	// Double-check after acquiring write lock.
	if l, ok := ns[namespace]; ok {
		return l, nil
	}
	ns[namespace] = Default()
	return ns[namespace], ErrNotExist
}

// SetNSLogger sets the sub-loggers for the given namespace.
// If the namespace does not exist, it is created.
func SetNSLogger(namespace string, logger Logger) {
	mu.Lock()
	defer mu.Unlock()
	if l, ok := ns[namespace]; ok {
		l.SubLoggers = logger.SubLoggers
	} else {
		ns[namespace] = &logger
	}
}

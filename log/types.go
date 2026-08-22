package log

import (
	"reflect"
	"sync"
	"sync/atomic"
	"unsafe"
)

// Default returns a new Logger with no sub-loggers.
func Default() *Logger {
	return &Logger{
		mu:         &sync.RWMutex{},
		SubLoggers: []LevelLogger{},
	}
}

// Compile-time assertions that *Logger satisfies the interfaces it advertises.
// Note: the level methods have pointer receivers (required for the internal
// mutex), so only *Logger — not a Logger value — satisfies LevelLogger. In
// v1.x the methods had value receivers; this changed in v2.0.0.
var (
	_ LevelLogger                             = (*Logger)(nil)
	_ interface{ Write([]byte) (int, error) } = (*Logger)(nil)
)

// EnrichLogger wraps a StdLogger (such as the standard library's log.Logger)
// to satisfy the LevelLogger interface by mapping all level methods to the
// basic Print/Println/Printf methods.
func EnrichLogger(weak StdLogger) LevelLogger {
	return gapLogger{subLogger: weak}
}

// Logger multiplexes log calls to multiple sub-loggers.
//
// The mutex is stored as a plain pointer (rather than an atomic.Pointer or an
// embedded sync.RWMutex) so that Logger remains free of noCopy fields and can
// still be passed by value, e.g. to SetNSLogger. Lazy initialization of the
// mutex for zero-value Loggers is made race-free via atomic operations on the
// pointer field.
type Logger struct {
	mu *sync.RWMutex
	// SubLoggers holds the registered sub-loggers. Use AddSubLogger/
	// RemoveSubLogger for concurrent-safe management; assigning this field
	// directly is not synchronized and races with concurrent logging.
	SubLoggers []LevelLogger
}

func (l *Logger) ensureMu() *sync.RWMutex {
	addr := (*unsafe.Pointer)(unsafe.Pointer(&l.mu))
	if mu := atomic.LoadPointer(addr); mu != nil {
		return (*sync.RWMutex)(mu)
	}
	mu := &sync.RWMutex{}
	if atomic.CompareAndSwapPointer(addr, nil, unsafe.Pointer(mu)) {
		return mu
	}
	return (*sync.RWMutex)(atomic.LoadPointer(addr))
}

// AddSubLogger appends a sub-logger in a thread-safe manner.
func (l *Logger) AddSubLogger(sl LevelLogger) {
	mu := l.ensureMu()
	mu.Lock()
	defer mu.Unlock()
	l.SubLoggers = append(l.SubLoggers, sl)
}

// RemoveSubLogger removes the first occurrence of the given sub-logger.
// It returns true if the sub-logger was found and removed. Matching is by
// equality (==). Sub-loggers whose dynamic type is not comparable are skipped.
func (l *Logger) RemoveSubLogger(sl LevelLogger) bool {
	mu := l.ensureMu()
	mu.Lock()
	defer mu.Unlock()
	for i, existing := range l.SubLoggers {
		if sameSubLogger(existing, sl) {
			l.SubLoggers = append(l.SubLoggers[:i], l.SubLoggers[i+1:]...)
			return true
		}
	}
	return false
}

func sameSubLogger(existing, target LevelLogger) (eq bool) {
	if existing == nil || target == nil {
		return existing == target
	}
	if reflect.TypeOf(existing) != reflect.TypeOf(target) {
		return false
	}
	// reflect.Type.Comparable() only reports static comparability; a struct
	// with an interface field is statically comparable yet can panic at == if
	// that field holds a non-comparable dynamic value. Recover keeps eq false.
	defer func() { _ = recover() }()
	return existing == target
}

// Len returns the number of registered sub-loggers.
func (l *Logger) Len() int {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	return len(l.SubLoggers)
}

// Write implements io.Writer by delegating to Print on each sub-logger.
// This allows a Logger to be used as the output for the standard library's
// log.Logger or as an http.Server.ErrorLog writer.
func (l *Logger) Write(p []byte) (int, error) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	msg := string(p)
	for _, sl := range l.SubLoggers {
		sl.Print(msg)
	}
	return len(p), nil
}

type gapLogger struct {
	subLogger StdLogger
}

// StdLogger is the interface satisfied by the standard library's log.Logger.
type StdLogger interface {
	Println(v ...any)
	Printf(format string, v ...any)
	Print(v ...any)
	Panic(v ...any)
	Panicf(format string, v ...any)
	Panicln(v ...any)
	Fatal(v ...any)
	Fatalf(format string, v ...any)
	Fatalln(v ...any)
}

// LevelLogger is the interface that level-aware loggers must implement.
type LevelLogger interface {
	Debug(v ...any)
	Debugf(format string, v ...any)
	Debugln(v ...any)
	Error(v ...any)
	Errorf(format string, v ...any)
	Errorln(v ...any)
	Fatal(v ...any)
	Fatalf(format string, v ...any)
	Fatalln(v ...any)
	Info(v ...any)
	Infof(format string, v ...any)
	Infoln(v ...any)
	Notice(v ...any)
	Noticef(format string, v ...any)
	Noticeln(v ...any)
	Panic(v ...any)
	Panicf(format string, v ...any)
	Panicln(v ...any)
	Print(v ...any)
	Printf(format string, v ...any)
	Println(v ...any)
	Trace(v ...any)
	Tracef(format string, v ...any)
	Traceln(v ...any)
	Warn(v ...any)
	Warnf(format string, v ...any)
	Warnln(v ...any)
}

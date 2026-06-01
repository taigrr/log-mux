package log

import "sync"

// Default returns a new Logger with no sub-loggers.
func Default() *Logger {
	return &Logger{
		mu:         &sync.RWMutex{},
		SubLoggers: []LevelLogger{},
	}
}

// EnrichLogger wraps a StdLogger (such as the standard library's log.Logger)
// to satisfy the LevelLogger interface by mapping all level methods to the
// basic Print/Println/Printf methods.
func EnrichLogger(weak StdLogger) LevelLogger {
	return gapLogger{subLogger: weak}
}

// Logger multiplexes log calls to multiple sub-loggers.
type Logger struct {
	mu         *sync.RWMutex
	SubLoggers []LevelLogger
}

func (l *Logger) ensureMu() *sync.RWMutex {
	if l.mu == nil {
		l.mu = &sync.RWMutex{}
	}
	return l.mu
}

// AddSubLogger appends a sub-logger in a thread-safe manner.
func (l *Logger) AddSubLogger(sl LevelLogger) {
	mu := l.ensureMu()
	mu.Lock()
	defer mu.Unlock()
	l.SubLoggers = append(l.SubLoggers, sl)
}

// RemoveSubLogger removes the first occurrence of the given sub-logger.
// It returns true if the sub-logger was found and removed.
func (l *Logger) RemoveSubLogger(sl LevelLogger) bool {
	mu := l.ensureMu()
	mu.Lock()
	defer mu.Unlock()
	for i, existing := range l.SubLoggers {
		if existing == sl {
			l.SubLoggers = append(l.SubLoggers[:i], l.SubLoggers[i+1:]...)
			return true
		}
	}
	return false
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

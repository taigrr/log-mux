package log

// Trace prints out logs on trace level
func (l *Logger) Trace(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Trace(args...)
	}
}

// Tracef is a formatted print for Trace
func (l *Logger) Tracef(format string, args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Tracef(format, args...)
	}
}

// Traceln prints out logs on trace level with newline
func (l *Logger) Traceln(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Traceln(args...)
	}
}

// Debug prints out logs on debug level
func (l *Logger) Debug(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Debug(args...)
	}
}

// Debugf is a formatted print for Debug
func (l *Logger) Debugf(format string, args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Debugf(format, args...)
	}
}

// Debugln prints out logs on debug level with newline
func (l *Logger) Debugln(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Debugln(args...)
	}
}

// Info prints out logs on info level
func (l *Logger) Info(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Info(args...)
	}
}

// Infof is a formatted print for Info
func (l *Logger) Infof(format string, args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Infof(format, args...)
	}
}

// Infoln prints out logs on info level with newline
func (l *Logger) Infoln(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Infoln(args...)
	}
}

// Notice prints out logs on notice level
func (l *Logger) Notice(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Notice(args...)
	}
}

// Noticef is a formatted print for Notice
func (l *Logger) Noticef(format string, args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Noticef(format, args...)
	}
}

// Noticeln prints out logs on notice level with newline
func (l *Logger) Noticeln(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Noticeln(args...)
	}
}

// Warn prints out logs on warn level
func (l *Logger) Warn(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Warn(args...)
	}
}

// Warnf is a formatted print for Warn
func (l *Logger) Warnf(format string, args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Warnf(format, args...)
	}
}

// Warnln prints out logs on warn level with a newline
func (l *Logger) Warnln(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Warnln(args...)
	}
}

// Error prints out logs on error level
func (l *Logger) Error(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Error(args...)
	}
}

// Errorf is a formatted print for Error
func (l *Logger) Errorf(format string, args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Errorf(format, args...)
	}
}

// Errorln prints out logs on error level with a new line
func (l *Logger) Errorln(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Errorln(args...)
	}
}

// Panic prints out logs on panic level
func (l *Logger) Panic(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Panic(args...)
	}
}

// Panicf is a formatted print for Panic
func (l *Logger) Panicf(format string, args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Panicf(format, args...)
	}
}

// Panicln prints out logs on panic level with a newline
func (l *Logger) Panicln(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Panicln(args...)
	}
}

// Fatal prints out logs on fatal level
func (l *Logger) Fatal(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Fatal(args...)
	}
}

// Fatalf is a formatted print for Fatal
func (l *Logger) Fatalf(format string, args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Fatalf(format, args...)
	}
}

// Fatalln prints fatal level with a new line
func (l *Logger) Fatalln(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Fatalln(args...)
	}
}

// Print delegates to sub-loggers
func (l *Logger) Print(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Print(args...)
	}
}

// Printf delegates to sub-loggers
func (l *Logger) Printf(format string, args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Printf(format, args...)
	}
}

// Println delegates to sub-loggers
func (l *Logger) Println(args ...any) {
	mu := l.ensureMu()
	mu.RLock()
	defer mu.RUnlock()
	for _, sl := range l.SubLoggers {
		sl.Println(args...)
	}
}

package log

func (l *logger) SetInfoDepth(depth int) {
	l.FileInfoDepth = depth
}

// Trace prints out logs on trace level
func (l logger) Trace(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Trace(args...)
	}
}

// Formatted print for Trace
func (l logger) Tracef(format string, args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Tracef(format, args...)
	}
}

// Trace prints out logs on trace level with newline
func (l logger) Traceln(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Traceln(args...)
	}
}

// Debug prints out logs on debug level
func (l logger) Debug(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Debug(args...)
	}
}

// Formatted print for Debug
func (l logger) Debugf(format string, args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Debugf(format, args...)
	}
}

// Info prints out logs on info level
func (l logger) Info(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Info(args...)
	}
}

// Formatted print for Info
func (l logger) Infof(format string, args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Infof(format, args...)
	}
}

// Info prints out logs on info level with newline
func (l logger) Infoln(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Infoln(args...)
	}
}

// Notice prints out logs on notice level
func (l logger) Notice(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Notice(args...)
	}
}

// Formatted print for Notice
func (l logger) Noticef(format string, args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Noticef(format, args...)
	}
}

// Notice prints out logs on notice level with newline
func (l logger) Noticeln(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Noticeln(args...)
	}
}

// Warn prints out logs on warn level
func (l logger) Warn(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Warn(args...)
	}
}

// Formatted print for Warn
func (l logger) Warnf(format string, args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Warnf(format, args...)
	}
}

// Warn prints out logs on warn level with a newline
func (l logger) Warnln(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Warnln(args...)
	}
}

// Error prints out logs on error level
func (l logger) Error(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Error(args...)
	}
}

// Formatted print for error
func (l logger) Errorf(format string, args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Errorf(format, args...)
	}
}

// Error prints out logs on error level with a new line
func (l logger) Errorln(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Errorln(args...)
	}
}

// Panic prints out logs on panic level
func (l logger) Panic(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Panic(args...)
	}
}

// Formatted print for panic
func (l logger) Panicf(format string, args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Panicf(format, args...)
	}
}

// Panic prints out logs on panic level with a newline
func (l logger) Panicln(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Panicln(args...)
	}
}

// Fatal prints out logs on fatal level
func (l logger) Fatal(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Fatal(args...)
	}
}

// Formatted print for fatal
func (l logger) Fatalf(format string, args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Fatalf(format, args...)
	}
}

// Fatal prints fatal level with a new line
func (l logger) Fatalln(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Fatalln(args...)
	}
}

// Handles print
func (l logger) Print(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Print(args...)
	}
}

// Handles formatted print
func (l logger) Printf(format string, args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Printf(format, args...)
	}
}

// Handles print with new line
func (l logger) Println(args ...interface{}) {
	for _, sl := range l.SubLoggers {
		sl.Println(args...)
	}
}

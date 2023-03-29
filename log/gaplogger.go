package log

// Trace prints out logs on trace level
func (l gapLogger) Trace(args ...interface{}) {
	l.subLogger.Print(args...)
}

// Formatted print for Trace
func (l gapLogger) Tracef(format string, args ...interface{}) {
	l.subLogger.Printf(format, args...)
}

// Trace prints out logs on trace level with newline
func (l gapLogger) Traceln(args ...interface{}) {
	l.subLogger.Println(args...)
}

// Debug prints out logs on debug level
func (l gapLogger) Debug(args ...interface{}) {
	l.subLogger.Print(args...)
}

// Formatted print for Debug
func (l gapLogger) Debugf(format string, args ...interface{}) {
	l.subLogger.Printf(format, args...)
}

// Formatted print for Debug
func (l gapLogger) Debugln(args ...interface{}) {
	l.subLogger.Println(args...)
}

// Info prints out logs on info level
func (l gapLogger) Info(args ...interface{}) {
	l.subLogger.Print(args...)
}

// Formatted print for Info
func (l gapLogger) Infof(format string, args ...interface{}) {
	l.subLogger.Printf(format, args...)
}

// Info prints out logs on info level with newline
func (l gapLogger) Infoln(args ...interface{}) {
	l.subLogger.Println(args...)
}

// Notice prints out logs on notice level
func (l gapLogger) Notice(args ...interface{}) {
	l.subLogger.Print(args...)
}

// Formatted print for Notice
func (l gapLogger) Noticef(format string, args ...interface{}) {
	l.subLogger.Printf(format, args...)
}

// Notice prints out logs on notice level with newline
func (l gapLogger) Noticeln(args ...interface{}) {
	l.subLogger.Println(args...)
}

// Warn prints out logs on warn level
func (l gapLogger) Warn(args ...interface{}) {
	l.subLogger.Print(args...)
}

// Formatted print for Warn
func (l gapLogger) Warnf(format string, args ...interface{}) {
	l.subLogger.Printf(format, args...)
}

// Warn prints out logs on warn level with a newline
func (l gapLogger) Warnln(args ...interface{}) {
	l.subLogger.Println(args...)
}

// Error prints out logs on error level
func (l gapLogger) Error(args ...interface{}) {
	l.subLogger.Print(args...)
}

// Formatted print for error
func (l gapLogger) Errorf(format string, args ...interface{}) {
	l.subLogger.Printf(format, args...)
}

// Error prints out logs on error level with a new line
func (l gapLogger) Errorln(args ...interface{}) {
	l.subLogger.Println(args...)
}

// Panic prints out logs on panic level
func (l gapLogger) Panic(args ...interface{}) {
	l.subLogger.Panic(args...)
}

// Formatted print for panic
func (l gapLogger) Panicf(format string, args ...interface{}) {
	l.subLogger.Panicf(format, args...)
}

// Panic prints out logs on panic level with a newline
func (l gapLogger) Panicln(args ...interface{}) {
	l.subLogger.Panicln(args...)
}

// Fatal prints out logs on fatal level
func (l gapLogger) Fatal(args ...interface{}) {
	l.subLogger.Fatal(args...)
}

// Formatted print for fatal
func (l gapLogger) Fatalf(format string, args ...interface{}) {
	l.subLogger.Fatalf(format, args...)
}

// Fatal prints fatal level with a new line
func (l gapLogger) Fatalln(args ...interface{}) {
	l.subLogger.Fatalln(args...)
}

// Handles print to info
func (l gapLogger) Print(args ...interface{}) {
	l.subLogger.Print(args...)
}

// Handles formatted print to info
func (l gapLogger) Printf(format string, args ...interface{}) {
	l.subLogger.Printf(format, args...)
}

// Handles print to info with new line
func (l gapLogger) Println(args ...interface{}) {
	l.subLogger.Println(args...)
}

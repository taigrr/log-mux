package log

// Trace prints out logs on trace level
func (l gapLogger) Trace(args ...any) {
	l.subLogger.Print(args...)
}

// Tracef is a formatted print for Trace
func (l gapLogger) Tracef(format string, args ...any) {
	l.subLogger.Printf(format, args...)
}

// Traceln prints out logs on trace level with newline
func (l gapLogger) Traceln(args ...any) {
	l.subLogger.Println(args...)
}

// Debug prints out logs on debug level
func (l gapLogger) Debug(args ...any) {
	l.subLogger.Print(args...)
}

// Debugf is a formatted print for Debug
func (l gapLogger) Debugf(format string, args ...any) {
	l.subLogger.Printf(format, args...)
}

// Debugln prints out logs on debug level with newline
func (l gapLogger) Debugln(args ...any) {
	l.subLogger.Println(args...)
}

// Info prints out logs on info level
func (l gapLogger) Info(args ...any) {
	l.subLogger.Print(args...)
}

// Infof is a formatted print for Info
func (l gapLogger) Infof(format string, args ...any) {
	l.subLogger.Printf(format, args...)
}

// Infoln prints out logs on info level with newline
func (l gapLogger) Infoln(args ...any) {
	l.subLogger.Println(args...)
}

// Notice prints out logs on notice level
func (l gapLogger) Notice(args ...any) {
	l.subLogger.Print(args...)
}

// Noticef is a formatted print for Notice
func (l gapLogger) Noticef(format string, args ...any) {
	l.subLogger.Printf(format, args...)
}

// Noticeln prints out logs on notice level with newline
func (l gapLogger) Noticeln(args ...any) {
	l.subLogger.Println(args...)
}

// Warn prints out logs on warn level
func (l gapLogger) Warn(args ...any) {
	l.subLogger.Print(args...)
}

// Warnf is a formatted print for Warn
func (l gapLogger) Warnf(format string, args ...any) {
	l.subLogger.Printf(format, args...)
}

// Warnln prints out logs on warn level with a newline
func (l gapLogger) Warnln(args ...any) {
	l.subLogger.Println(args...)
}

// Error prints out logs on error level
func (l gapLogger) Error(args ...any) {
	l.subLogger.Print(args...)
}

// Errorf is a formatted print for Error
func (l gapLogger) Errorf(format string, args ...any) {
	l.subLogger.Printf(format, args...)
}

// Errorln prints out logs on error level with a new line
func (l gapLogger) Errorln(args ...any) {
	l.subLogger.Println(args...)
}

// Panic prints out logs on panic level
func (l gapLogger) Panic(args ...any) {
	l.subLogger.Panic(args...)
}

// Panicf is a formatted print for Panic
func (l gapLogger) Panicf(format string, args ...any) {
	l.subLogger.Panicf(format, args...)
}

// Panicln prints out logs on panic level with a newline
func (l gapLogger) Panicln(args ...any) {
	l.subLogger.Panicln(args...)
}

// Fatal prints out logs on fatal level
func (l gapLogger) Fatal(args ...any) {
	l.subLogger.Fatal(args...)
}

// Fatalf is a formatted print for Fatal
func (l gapLogger) Fatalf(format string, args ...any) {
	l.subLogger.Fatalf(format, args...)
}

// Fatalln prints fatal level with a new line
func (l gapLogger) Fatalln(args ...any) {
	l.subLogger.Fatalln(args...)
}

// Print delegates to sub-logger
func (l gapLogger) Print(args ...any) {
	l.subLogger.Print(args...)
}

// Printf delegates to sub-logger
func (l gapLogger) Printf(format string, args ...any) {
	l.subLogger.Printf(format, args...)
}

// Println delegates to sub-logger
func (l gapLogger) Println(args ...any) {
	l.subLogger.Println(args...)
}

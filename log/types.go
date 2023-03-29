package log

func Default() *Logger {
	l := Logger{SubLoggers: []levelLogger{}}
	return &l
}

func EnrichLogger(weak stdLogger) gapLogger {
	return gapLogger{subLogger: weak}
}

type Logger struct {
	SubLoggers []levelLogger
}
type gapLogger struct {
	subLogger stdLogger
}

type stdLogger interface {
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
type levelLogger interface {
	Debug(v ...any)
	Debugf(format string, v ...any)
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

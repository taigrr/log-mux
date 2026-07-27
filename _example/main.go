package main

import (
	"log"
	"os"
	"time"

	mlog "github.com/taigrr/log-mux/v2/log"
)

func main() {
	// Create the multiplexer
	l := mlog.Default()

	// Add the standard library logger (enriched to support all levels)
	l.AddSubLogger(mlog.EnrichLogger(log.Default()))

	// Add a second logger writing to stderr with a custom prefix
	errLog := log.New(os.Stderr, "[ERR] ", log.LstdFlags)
	l.AddSubLogger(mlog.EnrichLogger(errLog))

	// Use as an io.Writer (e.g. for http.Server.ErrorLog)
	writerLog := log.New(l, "[MUX] ", log.LstdFlags)
	writerLog.Println("log-mux also implements io.Writer")

	// Fan out to both loggers simultaneously
	for i := range 5 {
		l.Infof("tick %d", i)
		time.Sleep(time.Second)
	}
}

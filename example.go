package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/taigrr/log-socket/browser"
	ls "github.com/taigrr/log-socket/log"
	"github.com/taigrr/log-socket/ws"

	mlog "github.com/taigrr/log-mux/log"
)

var addr = flag.String("addr", "0.0.0.0:8080", "http service address")

func generateLogs(l mlog.Logger) {
	for {
		l.Info("This is an info log!")
		l.Trace("This is a trace log!")
		l.Debug("This is a debug log!")
		l.Warn("This is a warn log!")
		l.Error("This is an error log!")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	defer ls.Flush()
	flag.Parse()
	http.HandleFunc("/ws", ws.LogSocketHandler)
	http.HandleFunc("/", browser.LogSocketViewHandler)

	l := mlog.Default()
	stdLogger := log.Default()
	lsLogger := ls.Default()
	l.SubLoggers = append(l.SubLoggers, mlog.EnrichLogger(stdLogger), lsLogger)

	go generateLogs(*l)
	l.Fatal(http.ListenAndServe(*addr, nil))
}

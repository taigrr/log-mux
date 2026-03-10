package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/taigrr/log-socket/v2/browser"
	ls "github.com/taigrr/log-socket/v2/log"
	"github.com/taigrr/log-socket/v2/ws"

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
	// NOTE: log-socket/v2 must have the full LevelLogger interface
	// (Debugln, Infoln, etc.) for this to compile. Use the
	// cd/logsocket-parity branch of log-socket.
	lsLogger := ls.Default()
	_ = lsLogger // uncomment below once log-socket v2 has full parity
	l.SubLoggers = append(l.SubLoggers, mlog.EnrichLogger(stdLogger))
	// l.SubLoggers = append(l.SubLoggers, lsLogger)

	go generateLogs(*l)
	l.Fatal(http.ListenAndServe(*addr, nil))
}

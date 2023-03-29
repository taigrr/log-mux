# log-mux

This repo was created so that you can easily mux (combine) the features of
multiple logging packages while still using the best features of each.
Create a logging mux by calling Default() and appending your loggers into the
SubLoggers child property. For example:

```golang
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
```

The above code sample shows how you can log to both log-socket (a logging
library which allows printing json logs over a websocket) and the standard
log libary, which doesn't have calls like Warn() or Info().

Calls to `Panic()` or `Fatal()` (as well as their `ln` and `f` variants call
out to the embedded panic/fatal methods, and thus may result in only one logger
receiving the call. Calls are made in the order the Subloggers are added to the
mux logger.

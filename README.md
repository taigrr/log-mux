# log-mux

A Go logging multiplexer that fans out log calls to multiple backends simultaneously.

## Installation

```bash
go get github.com/taigrr/log-mux
```

## Usage

```go
package main

import (
    "log"
    "os"

    mlog "github.com/taigrr/log-mux/log"
)

func main() {
    l := mlog.Default()

    // Wrap the standard library logger (adds level methods like Info, Warn, etc.)
    l.AddSubLogger(mlog.EnrichLogger(log.Default()))

    // Add a second logger writing to stderr
    errLog := log.New(os.Stderr, "[ERR] ", log.LstdFlags)
    l.AddSubLogger(mlog.EnrichLogger(errLog))

    l.Info("this goes to all sub-loggers")
    l.Warn("so does this")
}
```

### Thread-Safe Sub-Logger Management

```go
l := mlog.Default()

// Add and remove sub-loggers safely from any goroutine
sl := mlog.EnrichLogger(log.Default())
l.AddSubLogger(sl)
l.RemoveSubLogger(sl)
l.Len() // number of registered sub-loggers
```

### io.Writer Interface

Logger implements `io.Writer`, so it can be used anywhere a writer is expected — for example, as the error log for an HTTP server:

```go
l := mlog.Default()
l.AddSubLogger(mlog.EnrichLogger(log.Default()))

server := &http.Server{
    Addr:     ":8080",
    ErrorLog: log.New(l, "[HTTP] ", log.LstdFlags),
}
```

### Namespaced Loggers

```go
// Get or create a namespaced logger
logger, err := mlog.GetNSLogger("http")
if err == mlog.ErrNotExist {
    // First access — logger was just created with no sub-loggers
}

// Configure it
mlog.SetNSLogger("http", mlog.Logger{
    SubLoggers: []mlog.LevelLogger{mlog.EnrichLogger(log.Default())},
})
```

### Custom Loggers

Any type implementing `LevelLogger` can be added as a sub-logger. For loggers that only implement the standard `Print`/`Fatal`/`Panic` methods, use `EnrichLogger` to promote them to `LevelLogger`.

## Log Levels

All standard levels are supported: `Trace`, `Debug`, `Info`, `Notice`, `Warn`, `Error`, `Panic`, `Fatal`, plus `Print` variants. Each level has plain, formatted (`f`), and newline (`ln`) variants.

> **Note:** `Panic` and `Fatal` calls may only reach the first sub-logger, since the standard library's implementations call `panic()`/`os.Exit()` immediately.

## Thread Safety

All logging methods and sub-logger management (`AddSubLogger`, `RemoveSubLogger`, `Len`) are safe for concurrent use from multiple goroutines.

## License

See [LICENSE](LICENSE).

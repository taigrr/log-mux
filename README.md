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
    mlog "github.com/taigrr/log-mux/log"
)

func main() {
    l := mlog.Default()

    // Wrap the standard library logger (adds level methods like Info, Warn, etc.)
    stdLogger := mlog.EnrichLogger(log.Default())
    l.SubLoggers = append(l.SubLoggers, stdLogger)

    // Add any LevelLogger implementation
    // l.SubLoggers = append(l.SubLoggers, myCustomLogger)

    l.Info("this goes to all sub-loggers")
    l.Warn("so does this")
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

## License

See [LICENSE](LICENSE).

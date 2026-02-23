package log

import (
	"bytes"
	"log"
	"strings"
	"sync"
	"testing"
)

// mockLogger captures log output for testing.
type mockLogger struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (m *mockLogger) output() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.buf.String()
}

func (m *mockLogger) reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.buf.Reset()
}

func newMockStdLogger() (*log.Logger, *mockLogger) {
	m := &mockLogger{}
	return log.New(&m.buf, "", 0), m
}

func TestDefaultLogger(t *testing.T) {
	l := Default()
	if l == nil {
		t.Fatal("Default() returned nil")
	}
	if len(l.SubLoggers) != 0 {
		t.Fatalf("expected 0 sub-loggers, got %d", len(l.SubLoggers))
	}
}

func TestEnrichLogger(t *testing.T) {
	stdLog, mock := newMockStdLogger()
	enriched := EnrichLogger(stdLog)

	l := Default()
	l.SubLoggers = append(l.SubLoggers, enriched)

	l.Info("hello")
	if !strings.Contains(mock.output(), "hello") {
		t.Fatalf("expected 'hello' in output, got: %s", mock.output())
	}
}

func TestMultiplexing(t *testing.T) {
	stdLog1, mock1 := newMockStdLogger()
	stdLog2, mock2 := newMockStdLogger()

	l := Default()
	l.SubLoggers = append(l.SubLoggers, EnrichLogger(stdLog1), EnrichLogger(stdLog2))

	l.Warn("warning!")
	if !strings.Contains(mock1.output(), "warning!") {
		t.Fatal("mock1 missing output")
	}
	if !strings.Contains(mock2.output(), "warning!") {
		t.Fatal("mock2 missing output")
	}
}

func TestAllLevels(t *testing.T) {
	methods := []struct {
		name string
		call func(l *Logger)
	}{
		{"Trace", func(l *Logger) { l.Trace("t") }},
		{"Tracef", func(l *Logger) { l.Tracef("%s", "t") }},
		{"Traceln", func(l *Logger) { l.Traceln("t") }},
		{"Debug", func(l *Logger) { l.Debug("t") }},
		{"Debugf", func(l *Logger) { l.Debugf("%s", "t") }},
		{"Debugln", func(l *Logger) { l.Debugln("t") }},
		{"Info", func(l *Logger) { l.Info("t") }},
		{"Infof", func(l *Logger) { l.Infof("%s", "t") }},
		{"Infoln", func(l *Logger) { l.Infoln("t") }},
		{"Notice", func(l *Logger) { l.Notice("t") }},
		{"Noticef", func(l *Logger) { l.Noticef("%s", "t") }},
		{"Noticeln", func(l *Logger) { l.Noticeln("t") }},
		{"Warn", func(l *Logger) { l.Warn("t") }},
		{"Warnf", func(l *Logger) { l.Warnf("%s", "t") }},
		{"Warnln", func(l *Logger) { l.Warnln("t") }},
		{"Error", func(l *Logger) { l.Error("t") }},
		{"Errorf", func(l *Logger) { l.Errorf("%s", "t") }},
		{"Errorln", func(l *Logger) { l.Errorln("t") }},
		{"Print", func(l *Logger) { l.Print("t") }},
		{"Printf", func(l *Logger) { l.Printf("%s", "t") }},
		{"Println", func(l *Logger) { l.Println("t") }},
	}

	for _, m := range methods {
		t.Run(m.name, func(t *testing.T) {
			stdLog, mock := newMockStdLogger()
			l := Default()
			l.SubLoggers = append(l.SubLoggers, EnrichLogger(stdLog))
			m.call(l)
			if !strings.Contains(mock.output(), "t") {
				t.Fatalf("%s: expected output, got: %q", m.name, mock.output())
			}
		})
	}
}

func TestNamespace(t *testing.T) {
	// Reset namespace state
	mu.Lock()
	ns = make(map[string]*Logger)
	mu.Unlock()

	// First call should return ErrNotExist
	l1, err := GetNSLogger("test")
	if err != ErrNotExist {
		t.Fatalf("expected ErrNotExist, got: %v", err)
	}
	if l1 == nil {
		t.Fatal("expected non-nil logger")
	}

	// Second call should return the same logger, no error
	l2, err := GetNSLogger("test")
	if err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
	if l1 != l2 {
		t.Fatal("expected same logger instance")
	}
}

func TestSetNSLogger(t *testing.T) {
	mu.Lock()
	ns = make(map[string]*Logger)
	mu.Unlock()

	stdLog, mock := newMockStdLogger()
	custom := Logger{SubLoggers: []LevelLogger{EnrichLogger(stdLog)}}
	SetNSLogger("custom", custom)

	l, err := GetNSLogger("custom")
	if err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
	l.Info("ns-test")
	if !strings.Contains(mock.output(), "ns-test") {
		t.Fatalf("expected output from namespace logger, got: %q", mock.output())
	}
}

func TestNamespaceConcurrency(t *testing.T) {
	mu.Lock()
	ns = make(map[string]*Logger)
	mu.Unlock()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetNSLogger("concurrent")
		}()
	}
	wg.Wait()
}

package log

import (
	"bytes"
	"io"
	"log"
	"testing"
)

func TestWritePreservesPayloadAcrossBackends(t *testing.T) {
	for _, test := range []struct {
		name    string
		payload []byte
	}{
		{name: "nil"},
		{name: "empty", payload: []byte{}},
		{name: "literal format verbs", payload: []byte("100% complete: %s %!\n")},
		{name: "multiline", payload: []byte("first\nsecond\n")},
		{name: "unicode and binary", payload: []byte("雪\x00\xff")},
	} {
		t.Run(test.name, func(t *testing.T) {
			payload := test.payload
			var outputs [2]bytes.Buffer
			logger := Default()
			for index := range outputs {
				logger.AddSubLogger(EnrichLogger(log.New(&outputs[index], "", 0)))
			}
			written, err := logger.Write(payload)
			if err != nil || written != len(payload) {
				t.Fatalf("Write = (%d, %v), want (%d, nil)", written, err, len(payload))
			}
			// The standard logger adds a newline only when the payload lacks one.
			expected := append([]byte(nil), payload...)
			if len(expected) == 0 || expected[len(expected)-1] != '\n' {
				expected = append(expected, '\n')
			}
			for index := range outputs {
				if !bytes.Equal(outputs[index].Bytes(), expected) {
					t.Errorf("backend %d received %q, want %q", index, outputs[index].Bytes(), expected)
				}
			}
		})
	}
}

func TestWriteWithoutBackends(t *testing.T) {
	for _, logger := range []*Logger{new(Logger), Default()} {
		payload := []byte("discarded output")
		written, err := io.Copy(logger, bytes.NewReader(payload))
		if err != nil || written != int64(len(payload)) {
			t.Fatalf("io.Copy = (%d, %v), want (%d, nil)", written, err, len(payload))
		}
		if logger.Len() != 0 {
			t.Fatal("writing to an empty logger registered a backend")
		}
	}
}

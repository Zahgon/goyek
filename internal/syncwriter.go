package internal

import (
	"io"
	"sync"
)

type syncWriter struct {
	writer io.Writer
	mu     sync.Mutex
}

func (w *syncWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *syncWriter) WriteString(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

var _ io.StringWriter = (*syncWriter)(nil)

// SyncWriter synchronizes writes to the underlying writer.
func SyncWriter(w io.Writer) io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

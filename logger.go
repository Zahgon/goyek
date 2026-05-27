package goyek

import (
	"io"
	"runtime"
	"sync"
)

// Logger is used by A's logging functions.
type Logger interface {
	Log(w io.Writer, args ...interface{})
	Logf(w io.Writer, format string, args ...interface{})
}

// FmtLogger uses fmt when logging. It only appends a new line at the end.
type FmtLogger struct{}

// Log is used by [A] logging functions.
func (l FmtLogger) Log(w io.Writer, args ...interface{}) { _ = "STUB: not implemented"; return }

// Logf is used by [A] logging functions.
func (l FmtLogger) Logf(w io.Writer, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// CodeLineLogger decorates the log with code line information and indentation.
type CodeLineLogger struct {
	mu          sync.Mutex
	helperPCs   map[uintptr]struct{} // functions to be skipped when writing file/line info
	helperNames map[string]struct{}  // helperPCs converted to function names
}

// Log is used by [A] logging functions.
func (l *CodeLineLogger) Log(w io.Writer, args ...interface{}) { _ = "STUB: not implemented"; return }

//nolint:errcheck // not checking errors when writing to output

// Logf is used by [A] logging functions.
func (l *CodeLineLogger) Logf(w io.Writer, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck // not checking errors when writing to output

// Helper marks the calling function as a helper function.
// When printing file and line information, that function will be skipped.
// Helper may be called simultaneously from multiple goroutines.
func (l *CodeLineLogger) Helper() { _ = "STUB: not implemented"; return }

// skip: runtime.Callers + CodeLineLogger.Helper + A.Helper

// map will be recreated next time it is needed

// decorate prefixes the string with the file and line of the call site
// and inserts the final newline and indentation spaces for formatting.
func (l *CodeLineLogger) decorate(s string) string { _ = "STUB: not implemented"; return "" }

// Truncate file name at last file name separator.

// Every line is indented at least 6 spaces.

// Second and subsequent lines are indented an additional 4 spaces.

// frameSkip searches, starting after skip frames, for the first caller frame
// in a function not marked as a helper and returns that frame.
// The search stops if it finds a tRunner function that
// was the entry point into the test and the test is not a subtest.
// This function must be called with l.mu held.
func (l *CodeLineLogger) frameSkip(skip int) runtime.Frame {
	_ = "STUB: not implemented"
	// The maximum number of stack frames to go through when skipping helper functions for
	// the purpose of decorating log messages.
	return *new(runtime.Frame)
}

// skip: runtime.Callers + CodeLineLogger.frameSkip

// We've gone up all the way to the runner calling
// the action (so the user must have
// called a.Helper from inside that action).

// If more helper PCs have been added since we last did the conversion

// Found a frame that wasn't inside a helper function.

func pcToName(pc uintptr) string { _ = "STUB: not implemented"; return "" }

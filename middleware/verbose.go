package middleware

import (
	"github.com/goyek/goyek/v3"
)

// SilentNonFailed is a middleware which makes sure that only output from failed tasks is printed.
//
// The behavior is based on the Go test runner when it is executed without the -v flag.
func SilentNonFailed(next goyek.Runner) goyek.Runner {
	_ = "STUB: not implemented"
	return *new(goyek.Runner)
}

//nolint:errcheck // not checking errors when writing to output

package middleware

import (
	"github.com/goyek/goyek/v3"
)

// BufferParallel is a middleware which buffers the output from parallel tasks
// to not have mixed output from parallel tasks execution.
func BufferParallel(next goyek.Runner) goyek.Runner {
	_ = "STUB: not implemented"
	return *new(goyek.Runner)
}

//nolint:errcheck // not checking errors when writing to output
